package main

import "testing"

func TestCaptureMatchCoverageFromTemplate(t *testing.T) {
	body := `- [ ] Coverage service link ([codecov](https://codecov.io/), [coveralls](https://coveralls.io/), etc.): <!-- https://app.codecov.io/gh/org/project -->`

	got := captureMatch(body, reCoverage)
	want := "https://app.codecov.io/gh/org/project"
	if got != want {
		t.Fatalf("captureMatch() = %q, want %q", got, want)
	}
}

func TestCaptureMatchCoverageFromPlainLabel(t *testing.T) {
	body := "Coverage: https://coveralls.io/github/org/project"

	got := captureMatch(body, reCoverage)
	want := "https://coveralls.io/github/org/project"
	if got != want {
		t.Fatalf("captureMatch() = %q, want %q", got, want)
	}
}
