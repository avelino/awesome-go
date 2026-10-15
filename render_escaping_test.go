package main

import (
	"strings"
	"testing"
)

// TestRenderEscaping verifies that README-derived values rendered into category
// and project pages are context-escaped by html/template: markup and quotes in
// titles/descriptions must not survive as raw HTML, and href attributes must
// not be escapable via quote characters.
func TestRenderEscaping(t *testing.T) {
	const (
		title       = `Bad "Title" & <b>bold</b>`
		description = `Desc with "quotes" & <img src=x onerror=alert(1)>`
	)

	cat := Category{
		Slug:        "bad-title",
		Title:       title,
		Description: description,
		Links: []Link{
			{
				Title:       "proj",
				URL:         "https://github.com/o/p",
				Description: description,
				ProjectSlug: "proj-slug",
			},
		},
	}

	var catBuf strings.Builder
	if err := tpl.Lookup("category-index.tmpl.html").Execute(&catBuf, cat); err != nil {
		t.Fatalf("render category: %v", err)
	}
	catOut := catBuf.String()

	for _, want := range []string{
		`<h1>Bad &#34;Title&#34; &amp; &lt;b&gt;bold&lt;/b&gt;</h1>`,
		`content="Desc with &#34;quotes&#34; &amp; &lt;img src=x onerror=alert(1)&gt; - Curated list`,
	} {
		if !strings.Contains(catOut, want) {
			t.Errorf("category output missing %q", want)
		}
	}
	for _, notWant := range []string{
		`<b>bold</b>`,
		`<img src=x onerror=alert(1)>`,
	} {
		if strings.Contains(catOut, notWant) {
			t.Errorf("category output still contains raw %q", notWant)
		}
	}

	p := &Project{
		Slug:          "proj-slug",
		CategorySlug:  "bad-title",
		CategoryTitle: "Cat & More",
		Title:         title,
		Description:   description,
		URL:           "https://github.com/o/p?a=b&c=d",
	}

	var projBuf strings.Builder
	if err := tpl.Lookup("project.tmpl.html").Execute(&projBuf, p); err != nil {
		t.Fatalf("render project: %v", err)
	}
	projOut := projBuf.String()

	for _, want := range []string{
		`<h1>Bad &#34;Title&#34; &amp; &lt;b&gt;bold&lt;/b&gt;</h1>`,
		`href="https://github.com/o/p?a=b&amp;c=d"`,
		`"name": "Bad \\\u0022Title\\\u0022 \\u0026 \\u003cb\\u003ebold\\u003c\/b\\u003e"`,
	} {
		if !strings.Contains(projOut, want) {
			t.Errorf("project output missing %q", want)
		}
	}
	if strings.Contains(projOut, `<img src=x onerror=alert(1)>`) {
		t.Errorf("project output still contains raw %q", `<img src=x onerror=alert(1)>`)
	}
}
