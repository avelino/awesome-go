// Package main contains tests for the static site generator.
package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/avelino/awesome-go/pkg/markdown"
)

var (
	readmeFile      = readmePath
	outputDir       = outDir
	indexOutputFile = outIndexFile
)

// Helper to fail test if error occurs
func requireNoErr(t *testing.T, err error, msg string) {
	t.Helper()
	if err != nil {
		if msg == "" {
			msg = "unexpected error"
		}
		t.Fatalf("%s: %v", msg, err)
	}
}

// Converts README.md into a goquery HTML document
func goqueryFromReadme(t *testing.T) *goquery.Document {
	t.Helper()
	content, err := os.ReadFile(readmeFile)
	requireNoErr(t, err, "read README.md")

	html, err := markdown.ToHTML(content)
	requireNoErr(t, err, "convert markdown to HTML")

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	requireNoErr(t, err, "parse HTML with goquery")

	return doc
}

// ------------------------ Tests ------------------------

// TestAlpha ensures categories and links are alphabetically ordered
func TestAlpha(t *testing.T) {
	doc := goqueryFromReadme(t)
	doc.Find("body > ul").Each(func(i int, s *goquery.Selection) {
		if i != 0 { // skip first root list if needed
			testListAlphabetical(t, s)
		}
	})
}

// TestDuplicatedLinks ensures no duplicate links in README
func TestDuplicatedLinks(t *testing.T) {
	doc := goqueryFromReadme(t)
	seen := make(map[string]bool)
	doc.Find("body li > a:first-child").Each(func(_ int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			t.Error("expected href attribute")
		}
		if seen[href] {
			t.Fatalf("duplicated link: %s", href)
		}
		seen[href] = true
	})
}

// TestSeparator ensures all entries with description use " - " separator
func TestSeparator(t *testing.T) {
	content, err := os.ReadFile(readmeFile)
	requireNoErr(t, err, "read README.md")

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "* [") && strings.Contains(line, "](") {
			if strings.Contains(line, " - ") {
				continue
			}
			if strings.HasSuffix(line, ")") {
				// only link, no description
				continue
			}
			t.Errorf("line missing description separator ' - ': %s", line)
		}
	}
}

// TestRenderIndex ensures README renders to index.html
func TestRenderIndex(t *testing.T) {
	err := mkdirAll(outputDir)
	requireNoErr(t, err, "create output directory")

	err = renderIndex(readmeFile, indexOutputFile)
	requireNoErr(t, err, "render index.html")

	info, err := os.Stat(indexOutputFile)
	requireNoErr(t, err, "index.html should exist")
	if info.Size() == 0 {
		t.Error("index.html is empty")
	}
}

// ------------------------ Helper functions ------------------------

// Recursively test list alphabetical order
func testListAlphabetical(t *testing.T, list *goquery.Selection) {
	// Test nested lists first
	list.Find("ul").Each(func(_ int, ul *goquery.Selection) {
		testListAlphabetical(t, ul)
		ul.RemoveFiltered("ul") // remove inner ul to check current list
	})

	// Then test current list
	items := list.Find("li > a:first-child").Map(func(_ int, sel *goquery.Selection) string {
		return strings.ToLower(sel.Text())
	})
	sorted := make([]string, len(items))
	copy(sorted, items)
	sort.Strings(sorted)
	for i, val := range items {
		if val != sorted[i] {
			t.Errorf("expected order '%s', got '%s'", sorted[i], val)
		}
	}
	if t.Failed() {
		t.Logf("expected order:\n%s", strings.Join(sorted, "\n"))
	}
}

// ------------------------ Leaderboard Test ------------------------

func TestLeaderboardGeneration(t *testing.T) {
	err := buildStaticSite()
	requireNoErr(t, err, "build static site")

	leaderboardPath := "leaderboard.json"
	info, err := os.Stat(leaderboardPath)
	requireNoErr(t, err, "leaderboard.json should exist")
	if info.Size() == 0 {
		t.Fatal("leaderboard.json is empty")
	}

	data, err := os.ReadFile(leaderboardPath)
	requireNoErr(t, err, "read leaderboard.json")

	var contributors []Contributor
	if err := json.Unmarshal(data, &contributors); err != nil {
		t.Fatalf("invalid leaderboard JSON: %v", err)
	}
	if len(contributors) == 0 {
		t.Fatal("no contributors found in leaderboard.json")
	}

	// Check first contributor has all required fields
	first := contributors[0]
	if first.Login == "" || first.AvatarURL == "" || first.HTMLURL == "" || first.Contributions == 0 {
		t.Errorf("first contributor missing required fields: %+v", first)
	}
}
