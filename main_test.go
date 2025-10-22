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

		// Check for invalid tags
		tagPattern := regexp.MustCompile(`\[([^\]]+)\]`)
		matches := tagPattern.FindAllStringSubmatch(description, -1)

		for _, match := range matches {
			tag := match[1]
			if !validTypeTags[tag] && !validStatusTags[tag] {
				t.Errorf("Invalid tag [%s] found in: %s", tag, description)
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
			expectedDesc: "Some description  with tags.",
		},
		{
			name:         "Application with stalled status",
			input:        "App description [app] [stalled].",
			expectedTags: 2,
			expectedDesc: "App description .",
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
			expectedDesc: "Library  without status.",
		},
		{
			name:         "Only status tag",
			input:        "Project [unmaintained] description.",
			expectedTags: 1,
			expectedDesc: "Project  description.",
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
