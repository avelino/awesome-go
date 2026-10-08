package main

import (
	"os"
	"strings"
	"testing"
)

func TestEntrySafety(t *testing.T) {
	tests := []struct {
		name string
		line string
		fail bool
	}{
		{
			name: "normal entry passes",
			line: "- [project](https://github.com/org/project) - Short, clear description.",
			fail: false,
		},
		{
			name: "http url passes",
			line: "- [project](http://github.com/org/project) - Short, clear description.",
			fail: false,
		},
		{
			name: "uppercase https url passes",
			line: "- [project](HTTPS://github.com/org/project) - Short, clear description.",
			fail: false,
		},
		{
			name: "angle bracket autolink passes",
			line: "- [sftp](https://github.com/pkg/sftp) - Package sftp implements the protocol as described in <https://example.com/spec.txt>.",
			fail: false,
		},
		{
			name: "comparison passes",
			line: "- [project](https://github.com/org/project) - Returns true when a < b.",
			fail: false,
		},
		{
			name: "metadata colon passes",
			line: "- [project](https://github.com/org/project) - Reads metadata: owner and license.",
			fail: false,
		},
		{
			name: "only equals passes",
			line: "- [project](https://github.com/org/project) - Keeps only=one mode.",
			fail: false,
		},
		{
			name: "img onerror fails",
			line: "- [project](https://github.com/org/project) - <img src=x onerror=alert(1)>.",
			fail: true,
		},
		{
			name: "star bullet img onerror fails",
			line: "* [project](https://github.com/org/project) - <img src=x onerror=alert(1)>.",
			fail: true,
		},
		{
			name: "star bullet normal entry passes",
			line: "* [project](https://github.com/org/project) - Short, clear description.",
			fail: false,
		},
		{
			name: "img onerror in name fails",
			line: "- [<img src=x onerror=alert(1)>](https://github.com/org/project) - Short description.",
			fail: true,
		},
		{
			name: "script tag fails",
			line: "- [project](https://github.com/org/project) - Runs <script>alert(1)</script>.",
			fail: true,
		},
		{
			name: "closing tag fails",
			line: "- [project](https://github.com/org/project) - Closes </div> early.",
			fail: true,
		},
		{
			name: "self closing tag fails",
			line: "- [project](https://github.com/org/project) - Breaks <br/> lines.",
			fail: true,
		},
		{
			name: "slash attribute tag fails",
			line: "- [project](https://github.com/org/project) - Uses <img/src=x onerror=alert(1)>.",
			fail: true,
		},
		{
			name: "anchor tag fails",
			line: "- [project](https://github.com/org/project) - See <a href=\"https://example.com\">docs</a>.",
			fail: true,
		},
		{
			name: "javascript url in description fails",
			line: "- [project](https://github.com/org/project) - Open javascript:alert(1) here.",
			fail: true,
		},
		{
			name: "javascript scheme with space fails",
			line: "- [project](https://github.com/org/project) - Open javascript :alert(1) here.",
			fail: true,
		},
		{
			name: "uppercase javascript scheme fails",
			line: "- [project](https://github.com/org/project) - Open JAVASCRIPT:alert(1) here.",
			fail: true,
		},
		{
			name: "data url in description fails",
			line: "- [project](https://github.com/org/project) - Loads data:text/html,hi.",
			fail: true,
		},
		{
			name: "vbscript url in description fails",
			line: "- [project](https://github.com/org/project) - Runs vbscript:msgbox(1).",
			fail: true,
		},
		{
			name: "onclick attribute fails",
			line: "- [project](https://github.com/org/project) - Fires onclick=alert(1) now.",
			fail: true,
		},
		{
			name: "onerror with space before equals fails",
			line: "- [project](https://github.com/org/project) - Fires onerror =alert(1) now.",
			fail: true,
		},
		{
			name: "mixed case onload in name fails",
			line: "- [widget OnLoad=alert(1)](https://github.com/org/project) - Short description.",
			fail: true,
		},
		{
			name: "parenthesized https url passes",
			line: "- [project](https://en.wikipedia.org/wiki/Go_(programming_language)) - Short description.",
			fail: false,
		},
		{
			name: "javascript entry url fails",
			line: "- [project](javascript:alert(1)) - Short description.",
			fail: true,
		},
		{
			name: "data entry url fails",
			line: "- [project](data:text/html,hi) - Short description.",
			fail: true,
		},
		{
			name: "vbscript entry url fails",
			line: "- [project](vbscript:msgbox(1)) - Short description.",
			fail: true,
		},
		{
			name: "ftp entry url fails",
			line: "- [project](ftp://example.com/project) - Short description.",
			fail: true,
		},
		{
			name: "scheme-less entry url fails",
			line: "- [project](github.com/org/project) - Short description.",
			fail: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, ok := parseEntry(tt.line)
			if !ok {
				t.Fatalf("parseEntry rejected fixture %q", tt.line)
			}
			switch tt.name {
			case "parenthesized https url passes":
				const want = "https://en.wikipedia.org/wiki/Go_(programming_language)"
				if e.url != want {
					t.Fatalf("url = %q, want %q", e.url, want)
				}
			case "javascript entry url fails":
				if e.url != "javascript:alert(1)" {
					t.Fatalf("url = %q, want javascript:alert(1)", e.url)
				}
			case "img onerror fails":
				if !strings.Contains(e.description, "<img src=x onerror=alert(1)>") {
					t.Fatalf("description = %q", e.description)
				}
			}
			lines, failed := reviewEntrySafety(e)
			if failed != tt.fail {
				t.Fatalf("reviewEntrySafety fail=%v, want %v, lines=%v", failed, tt.fail, lines)
			}
			if tt.fail && !strings.Contains(strings.Join(lines, "\n"), "Entry safety") {
				t.Fatalf("failure lines = %v", lines)
			}
			if !tt.fail && failed {
				t.Fatalf("normal entry failed: %v", lines)
			}
		})
	}
}

func TestCurrentReadmeEntriesPassSafety(t *testing.T) {
	data, err := os.ReadFile("../../../README.md")
	if err != nil {
		t.Fatal(err)
	}
	checked := 0
	for _, line := range strings.Split(string(data), "\n") {
		e, ok := parseEntry(line)
		if !ok || !isHTTPURL(e.url) {
			continue
		}
		checked++
		lines, failed := reviewEntrySafety(e)
		if failed {
			t.Errorf("current entry failed safety: %s\n%v", line, lines)
		}
	}
	if checked < 1000 {
		t.Fatalf("checked %d http(s) entries, expected the curated list", checked)
	}
}
