package main

import (
	"bytes"
	"fmt"
	cp "github.com/otiai10/copy"
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

// Source files
const readmePath = "README.md"

// This files should be copied 'as is' to outDir directory
var staticFiles = []string{
	"tmpl/assets",
	"tmpl/_redirects",
	"tmpl/robots.txt",
}

// TODO: embed
// Templates
var tplIndex = template.Must(template.ParseFiles("tmpl/tmpl.html"))
var tplCategoryIndex = template.Must(template.ParseFiles("tmpl/cat-tmpl.html"))
var tplSitemap = template.Must(template.ParseFiles("tmpl/sitemap-tmpl.xml"))

// Output files
const outDir = "out/" // NOTE: trailing slash is required

var outIndexFile = filepath.Join(outDir, "index.html")
var outSitemapFile = filepath.Join(outDir, "sitemap.xml")

func main() {
	// Cleanup and re-create output directory
	{
		if err := os.RemoveAll(outDir); err != nil {
			panic(err)
		}

		if err := mkdirAll(outDir); err != nil {
			panic(err)
		}
	}

	err := GenerateHTML(readmePath, outIndexFile)
	if err != nil {
		panic(err)
	}

	input, err := os.ReadFile(outIndexFile)
	if err != nil {
		panic(err)
	}

	query, err := goquery.NewDocumentFromReader(bytes.NewReader(input))
	if err != nil {
		panic(err)
	}

	objs := make(map[string]Object)
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
			objs[selector] = *obj
		})
	})

	if err := makeSiteStruct(objs); err != nil {
		// FIXME: remove all panics
		panic(err)
	}
	changeLinksInIndex(string(input), query, objs)

	makeSitemap(objs)

	for _, srcFilename := range staticFiles {
		dstFilename := filepath.Join(outDir, filepath.Base(srcFilename))
		fmt.Printf("Copy static file: %s -> %s\n", srcFilename, dstFilename)
		if err := cp.Copy(srcFilename, dstFilename); err != nil {
			panic(err)
		}
	}
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

func makeSiteStruct(objs map[string]Object) error {
	for _, obj := range objs {
		categoryDir := filepath.Join(outDir, obj.Slug)
		if err := mkdirAll(categoryDir); err != nil {
			return err
		}

		// FIXME: embed templates
		// FIXME: parse templates once at start
		categoryIndexFilename := filepath.Join(categoryDir, "index.html")
		f, err := os.Create(categoryIndexFilename)
		if err != nil {
			return err
		}

		fmt.Printf("Write category Index file: %s\n", categoryIndexFilename)

		if err := tplCategoryIndex.Execute(f, obj); err != nil {
			return err
		}
	}

	return nil
}

func makeSitemap(objs map[string]Object) {
	// FIXME: handle error
	f, _ := os.Create(outSitemapFile)
	fmt.Printf("Render Sitemap to: %s\n", outSitemapFile)

	_ = tplSitemap.Execute(f, objs)
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

func changeLinksInIndex(html string, query *goquery.Document, objs map[string]Object) {
	query.Find("body #content ul li ul li a").Each(func(_ int, s *goquery.Selection) {
		href, hrefExists := s.Attr("href")
		if !hrefExists {
			// FIXME: looks like is an error. Tag `a` in our case always
			//   	  should have `href` attr.
			return
		}

		// do not replace links if no page has been created for it
		_, objExists := objs[href]
		if !objExists {
			return
		}

		// FIXME: parse url
		uri := strings.SplitAfter(href, "#")
		if len(uri) >= 2 && uri[1] != "contents" {
			// FIXME: use s.SetAttr
			html = strings.ReplaceAll(
				html,
				fmt.Sprintf(`href="%s"`, href),
				fmt.Sprintf(`href="%s"`, uri[1]),
			)
		}
	})

	fmt.Printf("Rewrite links in Index file: %s\n", outIndexFile)
	_ = os.WriteFile(outIndexFile, []byte(html), 0644)
}
