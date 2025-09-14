/*
  Extract PR links from the pull request body and expose a single step output `body`.
  This script reads the GitHub event payload from GITHUB_EVENT_PATH, parses the PR body,
  captures the four URLs (repo, pkg.go.dev, goreportcard.com, coverage service),
  and writes a formatted comment to GITHUB_OUTPUT as `body=<comment>`.
*/

'use strict';

const fs = require('fs');

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

function capture(body, regex) {
  const match = body.match(regex);
  return match && match[1] ? match[1] : '';
}

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


