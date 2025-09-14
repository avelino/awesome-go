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

const fs = require('fs');
const https = require('https');

const GITHUB_OUTPUT = process.env.GITHUB_OUTPUT;

function readEvent() {
  try {
    return JSON.parse(fs.readFileSync(process.env.GITHUB_EVENT_PATH, 'utf8'));
  } catch {
    return {};
  }
}

function capture(body, regex) {
  const m = body.match(regex);
  return m && m[1] ? m[1].trim() : '';
}

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

async function checkGithubRepo(repoUrl) {
  const parsed = parseGithubRepo(repoUrl);
  if (!parsed) return { ok: false, reason: 'invalid repo url' };
  const { owner, repo } = parsed;
  const base = 'https://api.github.com';
  const headers = { 'User-Agent': 'awesome-go-quality-check' };
  const repoData = await fetchJson(`${base}/repos/${owner}/${repo}`, headers);
  if (!repoData) return { ok: false, reason: 'repo api not reachable' };
  if (repoData.archived) return { ok: false, reason: 'repo is archived' };
  const hasGoMod = await fetchJson(`${base}/repos/${owner}/${repo}/contents/go.mod`, headers);
  const releases = await fetchJson(`${base}/repos/${owner}/${repo}/releases`, headers);
  const hasRelease = Array.isArray(releases) && releases.some((r) => /^v\d+\.\d+\.\d+/.test(r.tag_name || ''));
  return {
    ok: Boolean(hasGoMod && hasGoMod.name === 'go.mod' && hasRelease),
    reason: !hasGoMod ? 'missing go.mod' : !hasRelease ? 'missing semver release' : undefined,
  };
}

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

async function checkPkgGoDev(url) {
  const res = await httpHeadOrGet(url);
  return { ok: res.ok };
}

async function checkCoverage(url) {
  const res = await httpHeadOrGet(url);
  return { ok: res.ok };
}

function setOutput(name, value) {
  if (!GITHUB_OUTPUT) return;
  fs.appendFileSync(GITHUB_OUTPUT, `${name}<<EOF\n${value}\nEOF\n`);
}

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
  const knownLabels = ['quality:ok', 'quality:fail', 'needs-info', 'needs-coverage'];
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


