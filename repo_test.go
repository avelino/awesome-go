package main

import (
	"io/ioutil"
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

func TestAlpha(t *testing.T) {
	query := startQuery()
	query.Find("body > ul").Each(func(i int, s *goquery.Selection) {
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
	query := startQuery()
	links := make(map[string]bool, 0)
	query.Find("body li > a:first-child").Each(func(_ int, s *goquery.Selection) {
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
	input, err := ioutil.ReadFile("./README.md")
	if err != nil {
		panic(err)
	}
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
func TestGenerateHTML(t *testing.T) {
	err := GenerateHTML()
	if err != nil {
		t.Errorf("html generate error '%s'", err.Error())
	}
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
