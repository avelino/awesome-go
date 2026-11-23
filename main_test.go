package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/avelino/awesome-go/pkg/markdown"
	"io"
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

	buf := strings.NewReader(string(html))
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

const (
	hostileTitle       = `Evil <script>alert(1)</script> "quoted" & co`
	hostileDescription = `Evil <script>alert(2)</script> "quoted" & 'single' <b>bold</b> - does things.`
	hostileURL         = `javascript:alert(document.cookie)`
)

func hostileCategory() Category {
	return Category{
		Title:       hostileTitle,
		Slug:        "evil-category",
		Description: hostileDescription,
		Links: []Link{
			{Title: hostileTitle, URL: hostileURL, Description: hostileDescription},
			{Title: "ok", URL: "https://example.com/x", Description: "ok - fine.", ProjectSlug: "ok-x-github"},
		},
	}
}

func hostileProject() *Project {
	return &Project{
		Title:         hostileTitle,
		URL:           hostileURL,
		Description:   hostileDescription,
		Slug:          "evil-project",
		Host:          "github",
		Owner:         "evil",
		Repo:          "project",
		CategoryTitle: hostileTitle,
		CategorySlug:  "evil-category",
		Related: []Link{
			{Title: hostileTitle, URL: hostileURL, Description: hostileDescription},
		},
	}
}

func renderedProject(t *testing.T, p *Project) string {
	t.Helper()

	var buf bytes.Buffer
	requireNoErr(t, renderProject(&buf, p), "project should be rendered")

	return buf.String()
}

func renderedCategory(t *testing.T, c Category) string {
	t.Helper()

	html, err := renderCategory(c)
	requireNoErr(t, err, "category should be rendered")

	return html
}

func docFromString(t *testing.T, html string) *goquery.Document {
	t.Helper()

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	requireNoErr(t, err, "html must be valid for goquery")

	return doc
}

// assertNoLiveMarkup fails if hostile README values became executable markup.
func assertNoLiveMarkup(t *testing.T, html string) {
	t.Helper()

	doc := docFromString(t, html)

	if strings.Contains(html, "<script>alert") {
		t.Errorf("rendered page contains a raw injected <script> tag")
	}

	// Only the JSON-LD block, the gtag loader and the gtag inline snippet are expected.
	if n := doc.Find("script").Length(); n != 3 {
		t.Errorf("expected exactly 3 script elements, got %d", n)
	}
	doc.Find("script").Each(func(_ int, s *goquery.Selection) {
		if typ, _ := s.Attr("type"); typ == "application/ld+json" {
			// Inert data, but it must not be able to close the script element.
			if strings.Contains(s.Text(), "<") {
				t.Errorf("raw < inside JSON-LD block: %q", s.Text())
			}
			return
		}
		if strings.Contains(s.Text(), "alert(") {
			t.Errorf("injected code reached a script element: %q", s.Text())
		}
	})

	doc.Find("[href]").Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if strings.HasPrefix(strings.ToLower(strings.TrimSpace(href)), "javascript:") {
			t.Errorf("live javascript: URL in href %q", href)
		}
	})

	if doc.Find("b").Length() != 0 {
		t.Errorf("injected <b> element was not escaped")
	}
}

func TestProjectPageEscapesReadmeValues(t *testing.T) {
	assertNoLiveMarkup(t, renderedProject(t, hostileProject()))
}

func TestCategoryPageEscapesReadmeValues(t *testing.T) {
	assertNoLiveMarkup(t, renderedCategory(t, hostileCategory()))
}

func ldJSON(t *testing.T, html string) map[string]any {
	t.Helper()

	sel := docFromString(t, html).Find(`script[type="application/ld+json"]`)
	if sel.Length() != 1 {
		t.Fatalf("expected one JSON-LD block, got %d", sel.Length())
	}

	var data map[string]any
	if err := json.Unmarshal([]byte(sel.Text()), &data); err != nil {
		t.Fatalf("JSON-LD is not valid JSON: %v\n%s", err, sel.Text())
	}

	return data
}

func breadcrumbNames(t *testing.T, data map[string]any) []string {
	t.Helper()

	crumb, _ := data["breadcrumb"].(map[string]any)
	items, _ := crumb["itemListElement"].([]any)

	var names []string
	for _, it := range items {
		m, _ := it.(map[string]any)
		name, _ := m["name"].(string)
		names = append(names, name)
	}

	return names
}

func assertEqual(t *testing.T, what, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("%s: got %q, want %q", what, got, want)
	}
}

func TestProjectJSONLDRoundTrips(t *testing.T) {
	p := hostileProject()
	p.URL = `https://example.com/a?x=1&y="2"`
	data := ldJSON(t, renderedProject(t, p))

	assertEqual(t, "name", data["name"].(string), hostileTitle)
	assertEqual(t, "description", data["description"].(string), hostileDescription)
	assertEqual(t, "codeRepository", data["codeRepository"].(string), p.URL)

	names := breadcrumbNames(t, data)
	if len(names) != 3 {
		t.Fatalf("expected 3 breadcrumb items, got %v", names)
	}
	assertEqual(t, "breadcrumb category", names[1], hostileTitle)
	assertEqual(t, "breadcrumb project", names[2], hostileTitle)
}

func TestProjectJSONLDWithMetaIsValid(t *testing.T) {
	p := hostileProject()
	p.Meta = &RepoMeta{Stars: 42}
	data := ldJSON(t, renderedProject(t, p))

	rating, _ := data["aggregateRating"].(map[string]any)
	if rating["ratingCount"] != float64(42) {
		t.Errorf("ratingCount: got %v, want 42", rating["ratingCount"])
	}
}

func TestCategoryJSONLDRoundTrips(t *testing.T) {
	data := ldJSON(t, renderedCategory(t, hostileCategory()))

	assertEqual(t, "name", data["name"].(string), hostileTitle+" - Awesome Go")
	assertEqual(t, "description", data["description"].(string), hostileDescription)

	names := breadcrumbNames(t, data)
	if len(names) != 2 {
		t.Fatalf("expected 2 breadcrumb items, got %v", names)
	}
	assertEqual(t, "breadcrumb category", names[1], hostileTitle)
}

func TestSitemapIsWellFormedXML(t *testing.T) {
	categories := map[string]Category{
		"c": {Title: "C", Slug: `cat&<evil>"'`},
	}
	projects := []*Project{
		{Slug: `proj&<evil>"'`, CategorySlug: `cat&<evil>"'`},
	}

	var buf bytes.Buffer
	requireNoErr(t, writeSitemap(&buf, categories, projects), "sitemap should be rendered")
	out := buf.String()

	const decl = `<?xml version="1.0" encoding="UTF-8"?>`
	if !strings.HasPrefix(out, decl) {
		t.Fatalf("sitemap must start with %s, got %.80q", decl, out)
	}

	if strings.Contains(out, "<evil>") {
		t.Errorf("unescaped markup in sitemap:\n%s", out)
	}

	var locs []string
	dec := xml.NewDecoder(strings.NewReader(out))
	for {
		tok, err := dec.Token()
		if errors.Is(err, io.EOF) {
			break
		}
		requireNoErr(t, err, "sitemap should be well-formed XML")

		if se, ok := tok.(xml.StartElement); ok && se.Name.Local == "loc" {
			var loc string
			requireNoErr(t, dec.DecodeElement(&loc, &se), "loc should decode")
			locs = append(locs, loc)
		}
	}

	want := []string{
		"https://awesome-go.com/",
		`https://awesome-go.com/cat&<evil>"'/`,
		`https://awesome-go.com/cat&<evil>"'/proj&<evil>"'/`,
	}
	want[1] = `https://awesome-go.com/cat&<evil>"'/`
	want[2] = `https://awesome-go.com/cat&<evil>"'/proj&<evil>"'/`
	if strings.Join(locs, "\n") != strings.Join(want, "\n") {
		t.Errorf("loc values after decoding:\n got %q\nwant %q", locs, want)
	}
}

func TestParseRepoURLKeepsGitLabNamespace(t *testing.T) {
	tests := []struct {
		url                           string
		wantHost, wantOwner, wantRepo string
	}{
		{"https://github.com/foo/bar", "github", "foo", "bar"},
		{"https://gitlab.com/foo/bar", "gitlab", "foo", "bar"},
		{"https://gitlab.com/group/sub/bar", "gitlab", "group/sub", "bar"},
		{"https://gitlab.com/group/sub/bar/-/tree/main", "gitlab", "group/sub", "bar"},
	}
	for _, tt := range tests {
		host, owner, repo, ok := parseRepoURL(tt.url)
		if !ok {
			t.Errorf("parseRepoURL(%q) not ok", tt.url)
			continue
		}
		if host != tt.wantHost || owner != tt.wantOwner || repo != tt.wantRepo {
			t.Errorf("parseRepoURL(%q) = (%q, %q, %q), want (%q, %q, %q)",
				tt.url, host, owner, repo, tt.wantHost, tt.wantOwner, tt.wantRepo)
		}
	}
}

func TestCacheFilePathDoesNotCollide(t *testing.T) {
	pairs := [][2]string{
		{"https://github.com/foo/bar", "https://gitlab.com/foo/bar"},
		{"https://gitlab.com/group/a/bar", "https://gitlab.com/group/b/bar"},
	}
	for _, pair := range pairs {
		var paths [2]string
		for i, u := range pair {
			host, owner, repo, ok := parseRepoURL(u)
			if !ok {
				t.Fatalf("parseRepoURL(%q) not ok", u)
			}
			paths[i] = cacheFilePath(&Project{Host: host, Owner: owner, Repo: repo})
		}
		if paths[0] == paths[1] {
			t.Errorf("cache paths collide for %q and %q: %s", pair[0], pair[1], paths[0])
		}
	}
}
