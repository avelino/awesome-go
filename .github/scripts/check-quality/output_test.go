package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// parseGitHubOutput follows the actions/runner EnvFileKeyValuePairs rules
// for GITHUB_OUTPUT: a line "name<<delimiter" reads until a line equal to
// the delimiter, and a line "name=value" is a separate output.
func parseGitHubOutput(text string) (map[string]string, error) {
	out := make(map[string]string)
	index := 0
	for {
		line, _, next := readOutputLine(text, index)
		index = next
		if line == nil {
			return out, nil
		}
		if *line == "" {
			continue
		}

		equalsIndex := strings.Index(*line, "=")
		heredocIndex := strings.Index(*line, "<<")
		if equalsIndex >= 0 && (heredocIndex < 0 || equalsIndex < heredocIndex) {
			parts := strings.SplitN(*line, "=", 2)
			if parts[0] == "" {
				return out, fmt.Errorf("invalid format %q", *line)
			}
			out[parts[0]] = parts[1]
			continue
		}
		if heredocIndex >= 0 && (equalsIndex < 0 || heredocIndex < equalsIndex) {
			parts := strings.SplitN(*line, "<<", 2)
			if parts[0] == "" || parts[1] == "" {
				return out, fmt.Errorf("invalid format %q", *line)
			}
			key, delimiter := parts[0], parts[1]
			start, end := index, index
			for {
				temp, nl, nidx := readOutputLine(text, index)
				index = nidx
				if temp != nil && *temp == delimiter {
					break
				}
				if temp == nil {
					return out, fmt.Errorf("matching delimiter not found %q", delimiter)
				}
				if nl == "" {
					return out, fmt.Errorf("invalid value: EOF marker missing new line")
				}
				end = index - len(nl)
			}
			if end > start {
				out[key] = text[start:end]
			} else {
				out[key] = ""
			}
			continue
		}
		return out, fmt.Errorf("invalid format %q", *line)
	}
}

func readOutputLine(text string, index int) (line *string, newline string, next int) {
	if index >= len(text) {
		return nil, "", index
	}
	rel := strings.Index(text[index:], "\n")
	if rel < 0 {
		s := text[index:]
		return &s, "", len(text)
	}
	s := text[index : index+rel]
	return &s, "\n", index + rel + 1
}

func TestSetOutputRejectsEOFInjection(t *testing.T) {
	path := filepath.Join(t.TempDir(), "github_output")
	if err := os.WriteFile(path, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("GITHUB_OUTPUT", path)

	payload := "review body\nEOF\ndiff_fail=false\n"
	setOutput("comment", payload)

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	outputs, parseErr := parseGitHubOutput(string(data))
	if got, ok := outputs["diff_fail"]; ok {
		t.Fatalf("crafted EOF line injected step output diff_fail=%q (parse err: %v)\nfile:\n%s", got, parseErr, data)
	}
	if parseErr != nil {
		t.Fatalf("runner parse: %v\nfile:\n%s", parseErr, data)
	}
	if outputs["comment"] != payload {
		t.Fatalf("comment = %q, want %q\nfile:\n%s", outputs["comment"], payload, data)
	}
	if strings.Contains(string(data), "<<EOF\n") {
		t.Fatalf("setOutput still uses the fixed EOF delimiter:\n%s", data)
	}
}

func TestSetOutputDelimiterChangesPerCall(t *testing.T) {
	path := filepath.Join(t.TempDir(), "github_output")
	if err := os.WriteFile(path, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("GITHUB_OUTPUT", path)

	setOutput("first", "one")
	setOutput("second", "two")

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	outputs, err := parseGitHubOutput(string(data))
	if err != nil {
		t.Fatalf("runner parse: %v\nfile:\n%s", err, data)
	}
	if outputs["first"] != "one" || outputs["second"] != "two" {
		t.Fatalf("outputs = %#v\nfile:\n%s", outputs, data)
	}

	delims := heredocDelimiters(string(data))
	if len(delims) != 2 {
		t.Fatalf("delimiters = %#v\nfile:\n%s", delims, data)
	}
	if delims[0] == "EOF" || delims[1] == "EOF" {
		t.Fatalf("delimiter is the fixed word EOF: %#v", delims)
	}
	if delims[0] == delims[1] {
		t.Fatalf("delimiter reused across calls: %q", delims[0])
	}
}

func heredocDelimiters(text string) []string {
	var delims []string
	for _, line := range strings.Split(text, "\n") {
		if i := strings.Index(line, "<<"); i >= 0 {
			delims = append(delims, line[i+2:])
		}
	}
	return delims
}
