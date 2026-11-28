package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCaptureMatchCoverageFromTemplate(t *testing.T) {
	body := `- [ ] Coverage service link ([codecov](https://codecov.io/), [coveralls](https://coveralls.io/), etc.): <!-- https://app.codecov.io/gh/org/project -->`

	got := captureMatch(body, reCoverage)
	want := ""
	if got != want {
		t.Fatalf("captureMatch() = %q, want %q", got, want)
	}
}

func TestCaptureMatchCoverageFromFilledTemplate(t *testing.T) {
	body := `- [x] Coverage service link ([codecov](https://codecov.io/), [coveralls](https://coveralls.io/), etc.): https://app.codecov.io/gh/org/project`

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

func TestCaptureMatchCoverageStripsTrailingChars(t *testing.T) {
	tests := []struct {
		name string
		body string
		want string
	}{
		{
			name: "closing paren",
			body: "Coverage: https://app.codecov.io/gh/org/project)",
			want: "https://app.codecov.io/gh/org/project",
		},
		{
			name: "trailing comment arrow without space",
			body: "Coverage: https://app.codecov.io/gh/org/project-->",
			want: "https://app.codecov.io/gh/org/project",
		},
		{
			name: "html comment without spaces",
			body: "- [ ] Coverage service link ([codecov](https://codecov.io/), [coveralls](https://coveralls.io/), etc.): <!--https://app.codecov.io/gh/org/project-->",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := captureMatch(tt.body, reCoverage); got != tt.want {
				t.Fatalf("captureMatch(%q) = %q, want %q", tt.body, got, tt.want)
			}
		})
	}
}

func TestSetOutputMultilineLabels(t *testing.T) {
	outFile := filepath.Join(t.TempDir(), "github_output")
	if err := os.WriteFile(outFile, nil, 0644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	t.Setenv("GITHUB_OUTPUT", outFile)

	labels := []string{"needs-coverage", "quality:ok"}
	setOutput("labels", strings.Join(labels, "\n"))

	data, err := os.ReadFile(outFile)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}
	outputs, err := parseGitHubOutput(string(data))
	if err != nil {
		t.Fatalf("parseGitHubOutput() error = %v", err)
	}
	want := "needs-coverage\nquality:ok"
	if outputs["labels"] != want {
		t.Fatalf("outputs[\"labels\"] = %q, want %q", outputs["labels"], want)
	}
}
