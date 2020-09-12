package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
	"testing"
	"text/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	gfm "github.com/shurcooL/github_flavored_markdown"
)

type content struct {
	Body string
}

func TestAlpha(t *testing.T) {
	query := startQuery()

	query.Find("body > ul").Each(func(_ int, s *goquery.Selection) {
		testList(t, s)
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

var (
	reContainsLink        = regexp.MustCompile(`\* \[.*\]\(.*\)`)
	reOnlyLink            = regexp.MustCompile(`\* \[.*\]\([^()]*\)$`)
	reLinkWithDescription = regexp.MustCompile(`\* \[.*\]\(.*\) - \S.*[\.\!]`)
)

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
	err := generateHTML()
	if err != nil {
		t.Errorf("html generate error '%s'", err.Error())
	}
}

func testList(t *testing.T, list *goquery.Selection) {
	list.Find("ul").Each(func(_ int, items *goquery.Selection) {
		testList(t, items)
		items.RemoveFiltered("ul")
	})

	category := list.Prev().Text()

	t.Run(category, func(t *testing.T) {
		checkAlphabeticOrder(t, list)
	})
}

func readme() []byte {
	input, err := ioutil.ReadFile("./README.md")
	if err != nil {
		panic(err)
	}
	html := append([]byte("<body>"), blackfriday.MarkdownCommon(input)...)
	html = append(html, []byte("</body>")...)
	return html
}

func startQuery() *goquery.Document {
	buf := bytes.NewBuffer(readme())
	query, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		panic(err)
	}

	return query
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

func generateHTML() (err error) {
	// options
	readmePath := "./README.md"
	tplPath := "tmpl/tmpl.html"
	idxPath := "tmpl/index.html"
	input, _ := ioutil.ReadFile(readmePath)
	body := string(gfm.Markdown(input))
	c := &content{Body: body}
	t := template.Must(template.ParseFiles(tplPath))
	f, err := os.Create(idxPath)
	t.Execute(f, c)
	return
}
