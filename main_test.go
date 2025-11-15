// Package main contains tests for the static site generator.
package main

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

// TestRenderIndex ensures that the index.html is generated correctly.
func TestRenderIndex(t *testing.T) {
	// Ensure output directory exists
	if err := mkdirAll(outDir); err != nil {
		t.Fatalf("mkdirAll failed: %v", err)
	}

	// Render index
	if err := renderIndex(readmePath, outIndexFile); err != nil {
		t.Fatalf("renderIndex failed: %v", err)
	}

	info, err := os.Stat(outIndexFile)
	if err != nil {
		t.Fatalf("index.html does not exist: %v", err)
	}
	if info.Size() == 0 {
		t.Fatal("index.html is empty")
	}

	// Parse HTML to ensure body exists
	content, _ := os.ReadFile(outIndexFile)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
	if err != nil {
		t.Fatalf("failed to parse index.html: %v", err)
	}
	if doc.Find("body").Length() == 0 {
		t.Fatal("index.html has no <body> tag")
	}
}

// TestLeaderboardGeneration ensures contributors JSON is correctly generated.
func TestLeaderboardGeneration(t *testing.T) {
	// Use mock leaderboard for safe testing
	os.Setenv("TEST_LEADERBOARD", "1")
	defer os.Unsetenv("TEST_LEADERBOARD")

	// Build static site
	if err := buildStaticSite(); err != nil {
		t.Fatalf("buildStaticSite failed: %v", err)
	}

	// Check leaderboard.json exists
	info, err := os.Stat(leaderboardJSON)
	if err != nil {
		t.Fatalf("leaderboard.json missing: %v", err)
	}
	if info.Size() == 0 {
		t.Fatal("leaderboard.json is empty")
	}

	// Read and validate JSON
	data, err := os.ReadFile(leaderboardJSON)
	if err != nil {
		t.Fatalf("read leaderboard.json failed: %v", err)
	}

	var contributors []Contributor
	if err := json.Unmarshal(data, &contributors); err != nil {
		t.Fatalf("invalid leaderboard JSON: %v", err)
	}
	if len(contributors) == 0 {
		t.Fatal("no contributors found in leaderboard.json")
	}

	// Validate first mock contributor
	first := contributors[0]
	if first.Login == "" || first.Contributions == 0 {
		t.Errorf("first contributor missing required fields: %+v", first)
	}
}
