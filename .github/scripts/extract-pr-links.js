/*
  Extract PR links from the pull request body and expose a single step output `body`.
  This script reads the GitHub event payload from GITHUB_EVENT_PATH, parses the PR body,
  captures the four URLs (repo, pkg.go.dev, goreportcard.com, coverage service),
  and writes a formatted comment to GITHUB_OUTPUT as `body=<comment>`.
*/

'use strict';

const fs = require('fs');

/**
 * Load and parse the GitHub event payload referenced by GITHUB_EVENT_PATH.
 *
 * Reads the file path from the GITHUB_EVENT_PATH environment variable and returns
 * the parsed JSON object. If the env var is missing, the file does not exist,
 * or the file contains invalid JSON, an empty object is returned.
 *
 * @return {Object} The parsed event payload, or an empty object on missing/invalid input.
 */
function readEventPayload() {
  const eventPath = process.env.GITHUB_EVENT_PATH;
  if (!eventPath || !fs.existsSync(eventPath)) {
    return {};
  }
  const content = fs.readFileSync(eventPath, 'utf8');
  try {
    return JSON.parse(content);
  } catch {
    return {};
  }
}

/**
 * Extracts the first capturing group's text from `body` using `regex`.
 *
 * If `regex` matches and contains a capturing group, returns the first capture; otherwise returns an empty string.
 *
 * @param {string} body - Input text to search.
 * @param {RegExp} regex - Regular expression with at least one capturing group.
 * @return {string} The first captured substring, or an empty string if no match or no capture.
 */
function capture(body, regex) {
  const match = body.match(regex);
  return match && match[1] ? match[1] : '';
}

/**
 * Build the multiline comment text summarizing the provided PR links.
 *
 * Any of the properties may be empty or undefined; missing values are replaced with the string `"not provided"`.
 *
 * @param {{repo?: string, pkg?: string, gorep?: string, coverage?: string}} links - URLs extracted from the PR body.
 * @param {string} [links.repo] - Repository/forge link (e.g., GitHub/GitLab/Bitbucket).
 * @param {string} [links.pkg] - pkg.go.dev URL.
 * @param {string} [links.gorep] - goreportcard.com URL.
 * @param {string} [links.coverage] - Coverage URL (e.g., coveralls.io or codecov.io).
 * @returns {string} A single newline-joined string containing the formatted comment.
 */
function buildComment({ repo, pkg, gorep, coverage }) {
  const repoOut = repo || 'not provided';
  const pkgOut = pkg || 'not provided';
  const gorepOut = gorep || 'not provided';
  const coverageOut = coverage || 'not provided';

  return [
    'Thank you for contributing to [awesome-go](https://awesome-go.com/). We will review your contribution as soon as possible.',
    '',
    'Make sure you add the links in the body of the pull request that are requested in the [contribution guide](https://github.com/avelino/awesome-go/blob/main/CONTRIBUTING.md):',
    `- repo link: @${repoOut}`,
    `- pkg.go.dev: @${pkgOut}`,
    `- goreportcard.com: @${gorepOut}`,
    `- coverage: @${coverageOut}`,
    '',
    '> Your project is under review. It may take a few days to be approved.'
  ].join('\n');
}

/**
 * Orchestrates extraction of PR links from the GitHub event payload and emits a formatted comment.
 *
 * Reads the GitHub Actions event payload (via readEventPayload), extracts up to four URLs from the pull request body
 * (a repository forge link, pkg.go.dev, goreportcard.com, and a coverage service URL) using predefined regexes,
 * builds a multi-line comment with buildComment, and exposes it as the step output named `body`.
 *
 * Side effects:
 * - If the GITHUB_OUTPUT environment variable is set, appends a `body<<EOF ... EOF` block to that file.
 * - If GITHUB_OUTPUT is not set, prints the comment to stdout.
 */
function main() {
  const event = readEventPayload();
  const prBody = (event.pull_request && event.pull_request.body) || '';

  const repo = capture(prBody, /forge\s+link[^:]*:\s*(https?:\/\/(?:github\.com|gitlab\.com|bitbucket\.org)\/\S+)/i);
  const pkg = capture(prBody, /pkg\.go\.dev:\s*(https?:\/\/pkg\.go\.dev\/\S+)/i);
  const gorep = capture(prBody, /goreportcard\.com:\s*(https?:\/\/goreportcard\.com\/\S+)/i);
  const coverage = capture(prBody, /coverage[^:]*:\s*(https?:\/\/(?:coveralls\.io|codecov\.io)\/\S+)/i);

  const comment = buildComment({ repo, pkg, gorep, coverage });

  const outPath = process.env.GITHUB_OUTPUT;
  if (!outPath) {
    console.log(comment);
    return;
  }
  // Write to GITHUB_OUTPUT
  fs.appendFileSync(outPath, `body<<EOF\n${comment}\nEOF\n`);
}

main();


