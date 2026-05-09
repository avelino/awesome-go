# awesome-go · LLM Contribution Guide

This document summarizes the project context and the conventions that language models must follow when helping on this repository.

## Project Snapshot

- Purpose: maintain the curated `README.md` list of Go resources and generate the static site via `go run .`.
- Primary language: Go 1.23 (see `go.mod`). Supporting JavaScript exists only for the GitHub Action in `.github/scripts`.
- Key entry points:
  - `main.go`: reads `README.md`, builds category pages, and writes artifacts to `out/` using templates in `tmpl/`.
  - `pkg/markdown` and `pkg/slug`: helper packages used by the generator.
  - `.github/workflows/`: CI pipelines for PR quality validation, stale checks, site deployment, and Go tests.

## When Modifying the Awesome List

- Read and respect `CONTRIBUTING.md` (alphabetical order, one item per PR, descriptions end with a period, etc.).
- Keep categories with at least three entries and verify surrounding entries still meet quality standards.
- Avoid promotional copy; descriptions must stay concise and neutral.
- Do not drop existing content unless removal is requested and justified.

## Coding Guidelines

- Go:
  - Use standard formatting (`gofmt`) and idiomatic Go style.
  - Favor small, testable functions; keep exported APIs documented with Go-style comments.
  - Maintain ≥80% coverage for non-data packages and ≥90% for data packages when adding new testable code.
- JavaScript (GitHub Action script):
  - Keep Node 20 compatibility.
  - Uphold strict mode and existing patterns (async helpers, early returns, descriptive errors).
- Generated site output is not committed; do not add files under `out/`.

## Testing & Validation

- Preferred commands before submitting Go changes:
  - `go test ./...`
  - `go test -run ^TestStaleRepository$` (mirrors the scheduled workflow focus).
- For list-only modifications, run linting/formatting on touched files if tools are configured locally.
- The `PR Quality Checks` workflow will run `.github/scripts/check-quality.js`; ensure referenced links in PR bodies are reachable.

## CI Overview

- `tests.yaml`: runs `go test main_test.go main.go` on pushes/PRs.
- `pr-quality-check.yaml`: validates PR metadata (forge link, pkg.go.dev, Go Report Card, coverage).
- `run-check.yaml`: scheduled stale repository audit via `go test -run ^TestStaleRepository$`.
- `site-deploy.yaml`: builds and deploys the static site to Netlify on `main` pushes.

## Documentation & Housekeeping

- Update this `AGENTS.md` whenever repository conventions change.
- Keep documentation in English; follow American English spelling for code and comments.
- Remove unused files/modules when confirmed obsolete.
- Align rendered documentation (`README.md`, `COVERAGE.md`, etc.) with any behavior changes made to `main.go` or helper packages.
