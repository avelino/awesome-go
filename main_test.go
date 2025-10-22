package main

import (
	"bytes"
	"github.com/avelino/awesome-go/pkg/markdown"
	"os"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

var (
	reContainsLink        = regexp.MustCompile(`\* \[.*\]\(.*\)`)
	reOnlyLink            = regexp.MustCompile(`\* \[.*\]\([^()]*\)$`)
	reLinkWithDescription = regexp.MustCompile(`\* \[.*\]\(.*\) - \S.*[\.\!]`)
)

func requireNoErr(t *testing.T, err error, msg string) {
	// FIXME: replace to github.com/stretchr/testify
	t.Helper()

	if msg == "" {
		msg = "unknown error"
	}

	if err != nil {
		t.Fatalf("Received unexpected error [%s]: %+v", msg, err)
	}
}

func goqueryFromReadme(t *testing.T) *goquery.Document {
	t.Helper()

	input, err := os.ReadFile(readmePath)
	requireNoErr(t, err, "readme file should be exists")

	html, err := markdown.ToHTML(input)
	requireNoErr(t, err, "markdown should be rendered to html")

	buf := bytes.NewBuffer(html)
	doc, err := goquery.NewDocumentFromReader(buf)
	requireNoErr(t, err, "html must be valid for goquery")

	return doc
}

func TestAlpha(t *testing.T) {
	doc := goqueryFromReadme(t)
	doc.Find("body > ul").Each(func(i int, s *goquery.Selection) {
		if i != 0 {
			// skip content menu
			// TODO: the sub items (with 3 hash marks `###`) are staying in
			// the main list, not respecting the hierarchy and making it
			// impossible to test the alphabetical order
			testList(t, s)
		}
	})
}

func TestDuplicatedLinks(t *testing.T) {
	doc := goqueryFromReadme(t)
	links := make(map[string]bool, 0)
	doc.Find("body li > a:first-child").Each(func(_ int, s *goquery.Selection) {
		t.Run(s.Text(), func(t *testing.T) {
			href, ok := s.Attr("href")
			if !ok {
				t.Error("expected to have href")
			}
			if links[href] {
				t.Fatalf("duplicated link '%s'", href)
			}
			links[href] = true
		})
	})
}

// Test if an entry has description, it must be separated from link with ` - `
func TestSeparator(t *testing.T) {
	var matched, containsLink, noDescription bool
	input, err := os.ReadFile(readmePath)
	requireNoErr(t, err, "readme should be exists")

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		line = strings.Trim(line, " ")
		containsLink = reContainsLink.MatchString(line)
		if containsLink {
			noDescription = reOnlyLink.MatchString(line)
			if noDescription {
				continue
			}
			matched = reLinkWithDescription.MatchString(line)
			if !matched {
				t.Errorf("expected entry to be in form of `* [link] - description.`, got '%s'", line)
			}
		}
	}
}

func TestRenderIndex(t *testing.T) {
	requireNoErr(t, mkdirAll(outDir), "output dir should exists")

	err := renderIndex(readmePath, outIndexFile)
	requireNoErr(t, err, "html should be rendered")
}

func testList(t *testing.T, list *goquery.Selection) {
	list.Find("ul").Each(func(_ int, items *goquery.Selection) {
		testList(t, items)
		items.RemoveFiltered("ul")
	})
	t.Run(list.Prev().Text(), func(t *testing.T) {
		checkAlphabeticOrder(t, list)
	})
}

func checkAlphabeticOrder(t *testing.T, s *goquery.Selection) {
	items := s.Find("li > a:first-child").Map(func(_ int, li *goquery.Selection) string {
		return strings.ToLower(li.Text())
	})
	sorted := make([]string, len(items))
	copy(sorted, items)
	sort.Strings(sorted)
	for k, item := range items {
		if item != sorted[k] {
			t.Errorf("expected '%s' but actual is '%s'", sorted[k], item)
		}
	}
	if t.Failed() {
		t.Logf("expected order is:\n%s", strings.Join(sorted, "\n"))
	}
}

// TestTagValidation ensures only valid tags are used
func TestTagValidation(t *testing.T) {
	validTypeTags := map[string]bool{"lib": true, "app": true}
	validStatusTags := map[string]bool{"active": true, "stalled": true, "unmaintained": true}

	doc := goqueryFromReadme(t)
	doc.Find("body li > a:first-child").Each(func(_ int, s *goquery.Selection) {
		parent := s.Parent()
		description := parent.Text()

		// Split description at first " - " to get tag zone
		parts := strings.SplitN(description, " - ", 2)
		tagZone := parts[0]

		// Check for bracketed tokens in tag zone only
		tagPattern := regexp.MustCompile(`\[([^\]]+)\]`)
		matches := tagPattern.FindAllStringSubmatch(tagZone, -1)

		for _, match := range matches {
			token := match[1]
			
			// Only validate letters-only tokens (ignore things like file extensions, type signatures)
			lettersOnly := regexp.MustCompile(`^[A-Za-z]+$`)
			if !lettersOnly.MatchString(token) {
				continue // Skip non-letters-only tokens
			}

			// Check if it's a valid tag
			if !validTypeTags[token] && !validStatusTags[token] {
				t.Errorf("Invalid tag [%s] found in tag zone of: %s", token, description)
			}
		}
	})
}

// TestTagParsing tests the tag parsing functionality
func TestTagParsing(t *testing.T) {
	testCases := []struct {
		name         string
		input        string
		expectedTags int
		expectedDesc string
	}{
		{
			name:         "Library with active status",
			input:        "Some description [lib] [active] with tags.",
			expectedTags: 2,
			expectedDesc: "Some description with tags.",
		},
		{
			name:         "Application with stalled status",
			input:        "App description [app] [stalled].",
			expectedTags: 2,
			expectedDesc: "App description.",
		},
		{
			name:         "No tags",
			input:        "No tags here.",
			expectedTags: 0,
			expectedDesc: "No tags here.",
		},
		{
			name:         "Only type tag",
			input:        "Library [lib] without status.",
			expectedTags: 1,
			expectedDesc: "Library without status.",
		},
		{
			name:         "Only status tag",
			input:        "Project [unmaintained] description.",
			expectedTags: 1,
			expectedDesc: "Project description.",
		},
		{
			name:         "Multiple spaces and spaces before punctuation",
			input:        "Project  [lib]   [active]  description  .",
			expectedTags: 2,
			expectedDesc: "Project description.",
		},
		{
			name:         "Tags before hyphen",
			input:        "Project [lib] [active] - detailed description here.",
			expectedTags: 2,
			expectedDesc: "Project - detailed description here.",
		},
		{
			name:         "Extra whitespace normalization",
			input:        "  Multiple   spaces  [app]  [stalled]   everywhere  ,  test  .  ",
			expectedTags: 2,
			expectedDesc: "Multiple spaces everywhere, test.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tags, desc := parseTagsFromDescription(tc.input)
			if len(tags) != tc.expectedTags {
				t.Errorf("Expected %d tags, got %d for input: %s", tc.expectedTags, len(tags), tc.input)
			}
			if desc != tc.expectedDesc {
				t.Errorf("Expected description '%s', got '%s'", tc.expectedDesc, desc)
			}
		})
	}
}

// TestTagParsingWithValues tests the tag parsing functionality and validates actual tag values
func TestTagParsingWithValues(t *testing.T) {
	testCases := []struct {
		name         string
		input        string
		expectedTags []struct {
			tagType string
			status  string
		}
		expectedDesc string
	}{
		{
			name:  "Library with active status",
			input: "Some description [lib] [active] with tags.",
			expectedTags: []struct {
				tagType string
				status  string
			}{
				{"lib", ""},
				{"", "active"},
			},
			expectedDesc: "Some description with tags.",
		},
		{
			name:  "Application with stalled status",
			input: "App description [app] [stalled].",
			expectedTags: []struct {
				tagType string
				status  string
			}{
				{"app", ""},
				{"", "stalled"},
			},
			expectedDesc: "App description.",
		},
		{
			name:         "No tags",
			input:        "No tags here.",
			expectedTags: []struct{ tagType, status string }{},
			expectedDesc: "No tags here.",
		},
		{
			name:  "Only type tag",
			input: "Library [lib] without status.",
			expectedTags: []struct {
				tagType string
				status  string
			}{
				{"lib", ""},
			},
			expectedDesc: "Library without status.",
		},
		{
			name:  "Only status tag", 
			input: "Project [unmaintained] description.",
			expectedTags: []struct {
				tagType string
				status  string
			}{
				{"", "unmaintained"},
			},
			expectedDesc: "Project description.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tags, desc := parseTagsFromDescription(tc.input)
			
			// Check tag count
			if len(tags) != len(tc.expectedTags) {
				t.Errorf("Expected %d tags, got %d for input: %s", len(tc.expectedTags), len(tags), tc.input)
				return
			}
			
			// Check actual tag values
			for i, expectedTag := range tc.expectedTags {
				if i >= len(tags) {
					t.Errorf("Missing tag at index %d", i)
					continue
				}
				
				actualTag := tags[i]
				if actualTag.Type != expectedTag.tagType || actualTag.Status != expectedTag.status {
					t.Errorf("Tag mismatch at index %d. Expected {Type: %q, Status: %q}, got {Type: %q, Status: %q}",
						i, expectedTag.tagType, expectedTag.status, actualTag.Type, actualTag.Status)
				}
			}
			
			// Check description
			if desc != tc.expectedDesc {
				t.Errorf("Expected description '%s', got '%s'", tc.expectedDesc, desc)
			}
		})
	}
}

// TestTagColorAssignment tests proper color class assignment
func TestTagColorAssignment(t *testing.T) {
	testCases := []struct {
		category string
		value    string
		expected string
	}{
		{"type", "lib", "tag-lib"},
		{"type", "app", "tag-app"},
		{"status", "active", "tag-active"},
		{"status", "stalled", "tag-stalled"},
		{"status", "unmaintained", "tag-unmaintained"},
		{"invalid", "test", "tag-default"},
	}

	for _, tc := range testCases {
		result := getTagColor(tc.category, tc.value)
		if result != tc.expected {
			t.Errorf("Expected color class '%s' for category '%s' value '%s', got '%s'",
				tc.expected, tc.category, tc.value, result)
		}
	}
}

// TestTagValidationEdgeCases tests that tag validation correctly handles edge cases
func TestTagValidationEdgeCases(t *testing.T) {
	testCases := []struct {
		name        string
		description string
		shouldPass  bool
		reason      string
	}{
		{
			name:        "Valid tags in tag zone",
			description: "project-name [lib] [active] - Description with [.go] files and [v1.2.3] version.",
			shouldPass:  true,
			reason:      "Valid tags should pass, non-letter brackets in description should be ignored",
		},
		{
			name:        "Invalid tag in tag zone",
			description: "project-name [invalid] - Description text.",
			shouldPass:  false,
			reason:      "Invalid letter-only tag should fail validation",
		},
		{
			name:        "File extensions in tag zone should be ignored",
			description: "project-name [.go] [.md] - File handling library.",
			shouldPass:  true,
			reason:      "File extensions (non-letters-only) should be ignored",
		},
		{
			name:        "Version numbers in tag zone should be ignored",
			description: "project-name [v1.2.3] - Version information.",
			shouldPass:  true,
			reason:      "Version numbers (non-letters-only) should be ignored",
		},
		{
			name:        "Type signatures in description should be ignored",
                        shouldPass:  false,
                        reason:      "Without ' - ' separator, [error] is letters-only but invalid tag",
			reason:      "Type signatures in description (after ' - ') should be ignored",
		},
		{
			name:        "Mixed valid and ignored tokens",
			description: "project-name [lib] [v1.0] [active] [.json] - Handles [string] and [int] types.",
			shouldPass:  true,
			reason:      "Should validate only letters-only tokens in tag zone, ignore rest",
		},
		{
			name:        "No tag separator",
			description: "project-name [lib] [active] Description without separator [.go] [error]",
                        shouldPass:  false,
                        reason:      "Without ' - ' separator, [error] is letters-only but invalid tag",
		},
	}

	validTypeTags := map[string]bool{"lib": true, "app": true}
	validStatusTags := map[string]bool{"active": true, "stalled": true, "unmaintained": true}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hasInvalidTag := false

			// Split description at first " - " to get tag zone
			parts := strings.SplitN(tc.description, " - ", 2)
			tagZone := parts[0]

			// Check for bracketed tokens in tag zone only
			tagPattern := regexp.MustCompile(`\[([^\]]+)\]`)
			matches := tagPattern.FindAllStringSubmatch(tagZone, -1)

			for _, match := range matches {
				token := match[1]
				
				// Only validate letters-only tokens
				lettersOnly := regexp.MustCompile(`^[A-Za-z]+$`)
				if !lettersOnly.MatchString(token) {
					continue // Skip non-letters-only tokens
				}

				// Check if it's a valid tag
				if !validTypeTags[token] && !validStatusTags[token] {
					hasInvalidTag = true
					break
				}
			}

			if tc.shouldPass && hasInvalidTag {
				t.Errorf("Expected to pass but found invalid tag in: %s (reason: %s)", tc.description, tc.reason)
			} else if !tc.shouldPass && !hasInvalidTag {
				t.Errorf("Expected to fail but no invalid tag found in: %s (reason: %s)", tc.description, tc.reason)
			}
		})
	}
}

// TestCleanupDescription tests the description cleanup functionality
func TestCleanupDescription(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Multiple spaces collapsed",
			input:    "Hello    world   test",
			expected: "Hello world test",
		},
		{
			name:     "Spaces before punctuation removed",
			input:    "Hello world  ,  test  .  End  !",
			expected: "Hello world, test. End!",
		},
		{
			name:     "Leading and trailing spaces trimmed",
			input:    "   Hello world   ",
			expected: "Hello world",
		},
		{
			name:     "Leading hyphen with spaces removed",
			input:    " - Hello world description",
			expected: "Hello world description",
		},
		{
			name:     "Leading hyphen without spaces removed",
			input:    "-Hello world description",
			expected: "Hello world description",
		},
		{
			name:     "Multiple leading hyphens and spaces",
			input:    "  -  - Hello world description",
			expected: "- Hello world description",
		},
		{
			name:     "Mixed whitespace and hyphen issues",
			input:    "  -  Multiple   spaces  ,   before   punctuation  .  ",
			expected: "Multiple spaces, before punctuation.",
		},
		{
			name:     "Various punctuation marks",
			input:    "Test  ;  semicolon  :  colon  ?  question  !  exclamation",
			expected: "Test; semicolon: colon? question! exclamation",
		},
		{
			name:     "Hyphen in middle preserved",
			input:    "Hello - world description",
			expected: "Hello - world description",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := cleanupDescription(tc.input)
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result)
			}
		})
	}
}
