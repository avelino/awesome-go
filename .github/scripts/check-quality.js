/*
  PR Quality Checks `CONTRIBUTING.md`
  - Extracts links from PR body (repo, pkg.go.dev, goreportcard, coverage)
  - Validates minimum standards from CONTRIBUTING.md (basic automated subset):
    * Repo accessible and not archived
    * Has go.mod and at least one SemVer release
    * Go Report Card reachable with acceptable grade (A- or better)
    * Coverage link reachable
    * pkg.go.dev page exists
  - Outputs a markdown report as `comment` and sets `fail=true` if critical checks fail
*/

'use strict';

const fs = require('fs');
const https = require('https');

const GITHUB_OUTPUT = process.env.GITHUB_OUTPUT;

/**
 * Read and parse the GitHub event JSON pointed to by GITHUB_EVENT_PATH.
 *
 * Attempts to read the file at the path given by the environment variable
 * `GITHUB_EVENT_PATH` and return its parsed JSON object. If the file is
 * missing, unreadable, or contains invalid JSON, returns an empty object.
 *
 * @returns {Object} The parsed event payload, or an empty object on error.
 */
function readEvent() {
  try {
    return JSON.parse(fs.readFileSync(process.env.GITHUB_EVENT_PATH, 'utf8'));
  } catch {
    return {};
  }
}

/**
 * Extracts the first capture group from `body` using `regex` and returns it trimmed.
 *
 * @param {string} body - Input text to match against.
 * @param {RegExp} regex - Regular expression that contains at least one capture group; the first group's value is returned.
 * @return {string} The trimmed first capture group if matched, otherwise an empty string.
 */
function capture(body, regex) {
  const m = body.match(regex);
  return m && m[1] ? m[1].trim() : '';
}

/**
 * Check whether a URL is reachable using an HTTP HEAD request with a GET fallback.
 *
 * Performs a HEAD request to the provided URL and:
 * - resolves { ok: true, status } for 2xx–3xx responses,
 * - resolves { ok: false, status } for 4xx responses,
 * - for other responses (including redirects or servers that don't support HEAD) retries with GET and resolves { ok: boolean, status } based on the GET status.
 * Network or request errors resolve to { ok: false }.
 *
 * The returned Promise never rejects; it always resolves with an object describing reachability.
 *
 * @param {string} url - The URL to check.
 * @returns {Promise<{ok: boolean, status?: number}>} Reachability result; `status` is the HTTP status code when available.
 */
function httpHeadOrGet(url) {
  return new Promise((resolve) => {
    const req = https.request(url, { method: 'HEAD' }, (res) => {
      if (res.statusCode && res.statusCode >= 200 && res.statusCode < 400) {
        resolve({ ok: true, status: res.statusCode });
      } else if (res.statusCode && res.statusCode >= 400 && res.statusCode < 500) {
        resolve({ ok: false, status: res.statusCode });
      } else {
        // retry with GET on redirect or unsupported HEAD
        const req2 = https.request(url, { method: 'GET' }, (res2) => {
          resolve({ ok: (res2.statusCode || 500) < 400, status: res2.statusCode });
        });
        req2.on('error', () => resolve({ ok: false }));
        req2.end();
        return;
      }
    });
    req.on('error', () => resolve({ ok: false }));
    req.end();
  });
}

/**
 * Parse a GitHub repository URL and return its owner and repository name.
 * @param {string} repoUrl - URL pointing to a GitHub repository (e.g. "https://github.com/owner/repo").
 * @returns {{owner: string, repo: string}|null} An object with `owner` and `repo` when the URL is a valid GitHub repository path; otherwise `null`.
 */
function parseGithubRepo(repoUrl) {
  try {
    const u = new URL(repoUrl);
    if (u.hostname !== 'github.com') return null;
    const [, owner, repo] = u.pathname.split('/');
    if (!owner || !repo) return null;
    return { owner, repo };
  } catch {
    return null;
  }
}

/**
 * Fetch JSON from the given URL via HTTPS GET and return the parsed object.
 *
 * Performs an HTTPS GET request to the provided URL and returns the parsed
 * JSON body. If the response body is not valid JSON or a network error
 * occurs, resolves to null.
 *
 * @param {string} url - The URL to fetch.
 * @param {Object} [headers={}] - Optional HTTP headers to include in the request.
 * @returns {Promise<Object|null>} The parsed JSON object, or null on error or invalid JSON.
 */
async function fetchJson(url, headers = {}) {
  return new Promise((resolve) => {
    https
      .get(url, { headers }, (res) => {
        let data = '';
        res.on('data', (c) => (data += c));
        res.on('end', () => {
          try {
            resolve(JSON.parse(data));
          } catch {
            resolve(null);
          }
        });
      })
      .on('error', () => resolve(null));
  });
}

/**
 * Validate a GitHub repository URL for Go project quality requirements.
 *
 * Performs API checks to ensure the repository exists, is not archived,
 * contains a top-level `go.mod` file, and has at least one SemVer-formatted
 * release (tag matching `^v\d+\.\d+\.\d+`).
 *
 * @param {string} repoUrl - A repository URL (expected to be a GitHub URL).
 * @returns {Promise<{ok: boolean, reason?: string}>} Resolves with an object where
 *   `ok` is true only if the repo is reachable, not archived, has `go.mod`, and
 *   has a SemVer release. When `ok` is false, `reason` is one of:
 *   - "invalid repo url"
 *   - "repo api not reachable"
 *   - "repo is archived"
 *   - "missing go.mod"
 *   - "missing semver release"
 */
async function checkGithubRepo(repoUrl) {
  const parsed = parseGithubRepo(repoUrl);
  if (!parsed) return { ok: false, reason: 'invalid repo url' };
  const { owner, repo } = parsed;
  const base = 'https://api.github.com';
  const headers = { 
    'User-Agent': 'awesome-go-quality-check',
    'Accept': 'application/vnd.github+json',
  };
  const token = process.env.GITHUB_TOKEN;
  if (token) headers.Authorization = `Bearer ${token}`;
  const repoData = await fetchJson(`${base}/repos/${owner}/${repo}`, headers);
  if (!repoData) return { ok: false, reason: 'repo api not reachable' };
  if (repoData.archived) return { ok: false, reason: 'repo is archived' };
  const hasGoMod = await fetchJson(`${base}/repos/${owner}/${repo}/contents/go.mod`, headers);
  const releases = await fetchJson(`${base}/repos/${owner}/${repo}/releases`, headers);
  const hasRelease = Array.isArray(releases) && releases.some((r) => /^v\d+\.\d+\.\d+/.test(r.tag_name || ''));
  const hasGoModOk = Boolean(hasGoMod && hasGoMod.name === 'go.mod');
  return {
    ok: Boolean(hasGoModOk && hasRelease),
    reason: !hasGoModOk ? 'missing go.mod' : !hasRelease ? 'missing semver release' : undefined,
  };
}

/**
 * Check a Go Report Card page for reachability and grade (pass if A- or better).
 *
 * Performs a HEAD/GET reachability check; if reachable, fetches the page and
 * attempts to parse a `Grade: X` token (case-insensitive). Returns:
 * - { ok: false, reason: 'unreachable' } if the initial reachability check fails.
 * - { ok: false, reason: 'fetch error' } if the page fetch errors.
 * - { ok: true, grade: 'unknown' } if the page is reachable but no grade is found.
 * - { ok: true, grade: 'A'|'A+'|'A-'|... } for parsed grades; `ok` is true only
 *   when the parsed grade is A, A+ or A- (passes), otherwise `ok` is false.
 *
 * @param {string} url - URL of the Go Report Card page to check.
 * @return {Promise<Object>} Result object containing `ok` and either `grade` or `reason`.
 */
async function checkGoReportCard(url) {
  // Accept A- or better
  const res = await httpHeadOrGet(url);
  if (!res.ok) return { ok: false, reason: 'unreachable' };
  // Fetch page to parse grade (best-effort)
  return new Promise((resolve) => {
    https
      .get(url, (res2) => {
        let html = '';
        res2.on('data', (c) => (html += c));
        res2.on('end', () => {
          const m = html.match(/Grade:\s*([A-F][+-]?)/i);
          if (!m) return resolve({ ok: true, grade: 'unknown' });
          const grade = m[1].toUpperCase();
          const pass = /^A[-+]?$/.test(grade);
          resolve({ ok: pass, grade });
        });
      })
      .on('error', () => resolve({ ok: false, reason: 'fetch error' }));
  });
}

/**
 * Verify that a pkg.go.dev page is reachable.
 *
 * Performs an HTTP HEAD (with GET fallback) to the provided URL and returns an object indicating reachability.
 *
 * @param {string} url - The pkg.go.dev URL to check.
 * @return {{ok: boolean}} An object with `ok: true` when the page is reachable, otherwise `{ok: false}`.
 */
async function checkPkgGoDev(url) {
  const res = await httpHeadOrGet(url);
  return { ok: res.ok };
}

/**
 * Check whether a coverage service URL is reachable.
 *
 * @param {string} url - The coverage page URL to validate (e.g., Coveralls or Codecov link).
 * @returns {{ok: boolean}} An object with `ok: true` if the URL responded successfully, otherwise `ok: false`.
 */
async function checkCoverage(url) {
  const res = await httpHeadOrGet(url);
  return { ok: res.ok };
}

/**
 * Writes a named workflow output to the GitHub Actions GITHUB_OUTPUT file using a heredoc block.
 *
 * If the GITHUB_OUTPUT environment variable is unset, this is a no-op.
 *
 * @param {string} name - The output name (key) to write.
 * @param {string} value - The output value; written between the heredoc delimiters.
 */
function setOutput(name, value) {
  if (!GITHUB_OUTPUT) return;
  fs.appendFileSync(GITHUB_OUTPUT, `${name}<<EOF\n${value}\nEOF\n`);
}

/**
 * Run automated PR quality checks and emit GitHub Actions outputs.
 *
 * Reads the current PR body, extracts links for the repository (forge), pkg.go.dev,
 * goreportcard, and coverage services, performs best-effort validations for each,
 * and emits three GitHub Actions outputs: `comment` (markdown summary), `fail`
 * ("true"/"false"), and `labels` (JSON array). The comment summarizes missing links
 * and the result of each check; labels reflect missing information, coverage status,
 * and overall quality pass/fail.
 */
async function main() {
  const event = readEvent();
  const body = (event.pull_request && event.pull_request.body) || '';
  const repo = capture(body, /forge\s+link[^:]*:\s*(https?:\/\/(?:github\.com|gitlab\.com|bitbucket\.org)\/\S+)/i);
  const pkg = capture(body, /pkg\.go\.dev:\s*(https?:\/\/pkg\.go\.dev\/\S+)/i);
  const gorep = capture(body, /goreportcard\.com:\s*(https?:\/\/goreportcard\.com\/\S+)/i);
  const coverage = capture(body, /coverage[^:]*:\s*(https?:\/\/(?:coveralls\.io|codecov\.io)\/\S+)/i);

  const results = [];
  let criticalFail = false;
  let repoOk = false, pkgOk = false, gorepOk = false; // track core checks

  if (!repo) {
    results.push('- Repo link: missing');
    criticalFail = true;
  } else {
    const r = await checkGithubRepo(repo);
    if (!r.ok) { results.push(`- Repo: FAIL (${r.reason || 'unknown'})`); criticalFail = true; }
    else { results.push('- Repo: OK'); repoOk = true; }
  }

  if (!pkg) {
    results.push('- pkg.go.dev: missing');
    criticalFail = true;
  } else {
    const r = await checkPkgGoDev(pkg);
    if (!r.ok) { results.push('- pkg.go.dev: FAIL (unreachable)'); criticalFail = true; }
    else { results.push('- pkg.go.dev: OK'); pkgOk = true; }
  }

  if (!gorep) {
    results.push('- goreportcard: missing');
    criticalFail = true;
  } else {
    const r = await checkGoReportCard(gorep);
    if (!r.ok) { results.push(`- goreportcard: FAIL (${r.reason || r.grade || 'unreachable/low grade'})`); criticalFail = true; }
    else { results.push(`- goreportcard: OK${r.grade ? ` (grade ${r.grade})` : ''}`); gorepOk = true; }
  }

  let coverageOk = false;
  if (!coverage) {
    results.push('- coverage: missing');
  } else {
    const r = await checkCoverage(coverage);
    if (!r.ok) { results.push('- coverage: FAIL (unreachable)'); }
    else { results.push('- coverage: OK'); coverageOk = true; }
  }

  const header = 'Automated Quality Checks (from CONTRIBUTING minimum standards)';
  const note = 'These checks are a best-effort automation and do not replace human review.';
  const summary = results.join('\n');
  const comment = [header, '', summary, '', note].join('\n');

  // Labels to reflect status
  const labels = [];
  const missingInfo = (!repo || !pkg || !gorep);
  if (missingInfo) labels.push('needs-info');
  if (!coverageOk) labels.push('needs-coverage');
  if (criticalFail) labels.push('quality:fail');
  if (!criticalFail && repoOk && pkgOk && gorepOk) labels.push('quality:ok');

  setOutput('comment', comment);
  setOutput('fail', criticalFail ? 'true' : 'false');
  setOutput('labels', JSON.stringify(labels));
}

main().catch((e) => {
  const msg = `Quality checks failed to run: ${e?.message || e}`;
  if (GITHUB_OUTPUT) {
    fs.appendFileSync(GITHUB_OUTPUT, `comment<<EOF\n${msg}\nEOF\n`);
    fs.appendFileSync(GITHUB_OUTPUT, `fail<<EOF\ntrue\nEOF\n`);
  }
  process.exitCode = 0; // Do not crash the job here; let the fail step handle it
});


