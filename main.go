package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/avelino/awesome-go/pkg/slug"
)

type Link struct {
	Title       string
	Url         string
	Description string
}

type Object struct {
	Title       string
	Slug        string
	Description string
	Items       []Link
}

// Source
const readmePath = "README.md"

// Templates
const tmplCategory = "tmpl/cat-tmpl.html"
const tmplSitemap = "tmpl/sitemap-tmpl.xml"

// Output
const outDir = "out/"
const outIndexFile = "index.html"
const outSitemapFile = "sitemap.xml"

func main() {
	outIndexAbs := filepath.Join(outDir, outIndexFile)
	err := GenerateHTML(readmePath, outIndexAbs)
	if err != nil {
		panic(err)
	}

	input, err := os.ReadFile(outIndexAbs)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(input)
	query, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		panic(err)
	}

	objs := make(map[string]*Object)
	query.Find("body #contents").NextFiltered("ul").Find("ul").Each(func(_ int, s *goquery.Selection) {
		s.Find("li a").Each(func(_ int, s *goquery.Selection) {
			selector, exists := s.Attr("href")
			if !exists {
				return
			}
			obj := makeObjByID(selector, query.Find("body"))
			if obj == nil {
				return
			}
			objs[selector] = obj
		})
	})

	if err := makeSiteStruct(objs); err != nil {
		// FIXME: remove all panics
		panic(err)
	}
	changeLinksInIndex(string(input), query, objs)

	makeSitemap(objs)
}

func mkdirAll(path string) error {
	_, err := os.Stat(path)
	// NOTE: directory is exists
	if err == nil {
		return nil
	}

	// NOTE: unknown error
	if !os.IsNotExist(err) {
		return err
	}

	// NOTE: directory is not exists
	if err := os.MkdirAll(path, 0o755); err != nil {
		return err
	}

	return nil
}

func makeSiteStruct(objs map[string]*Object) error {
	for _, obj := range objs {
		outDir := filepath.Join(outDir, obj.Slug)
		if err := mkdirAll(outDir); err != nil {
			return err
		}

		// FIXME: embed templates
		// FIXME: parse templates once at start
		t := template.Must(template.ParseFiles(tmplCategory))
		f, err := os.Create(filepath.Join(outDir, "index.html"))
		if err != nil {
			return err
		}

		if err := t.Execute(f, obj); err != nil {
			return err
		}
	}

	return nil
}

func makeSitemap(objs map[string]*Object) {
	t := template.Must(template.ParseFiles(tmplSitemap))
	f, _ := os.Create(filepath.Join(outDir, outSitemapFile))
	t.Execute(f, objs)
}

func makeObjByID(selector string, s *goquery.Selection) (obj *Object) {
	s.Find(selector).Each(func(_ int, s *goquery.Selection) {
		desc := s.NextFiltered("p")
		ul := s.NextFilteredUntil("ul", "h2")

		links := []Link{}
		ul.Find("li").Each(func(_ int, s *goquery.Selection) {
			url, _ := s.Find("a").Attr("href")
			link := Link{
				Title:       s.Find("a").Text(),
				Description: s.Text(),
				Url:         url,
			}
			links = append(links, link)
		})
		if len(links) == 0 {
			return
		}
		obj = &Object{
			Slug:        slug.Generate(s.Text()),
			Title:       s.Text(),
			Description: desc.Text(),
			Items:       links,
		}
	})
	return
}

func changeLinksInIndex(html string, query *goquery.Document, objs map[string]*Object) {
	query.Find("body #content ul li ul li a").Each(func(_ int, s *goquery.Selection) {
		href, hrefExists := s.Attr("href")
		if !hrefExists {
			return
		}

		// do not replace links if no page has been created for it
		_, objExists := objs[href]
		if !objExists {
			return
		}

		uri := strings.SplitAfter(href, "#")
		if len(uri) >= 2 && uri[1] != "contents" {
			html = strings.ReplaceAll(
				html, fmt.Sprintf(`href="%s"`, href), fmt.Sprintf(`href="%s"`, uri[1]))
		}
	})

	os.WriteFile(filepath.Join(outDir, outIndexFile), []byte(html), 0644)
}
