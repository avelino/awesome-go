package main

import (
	"bytes"
	"errors"
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

type Category struct {
	Title       string
	Slug        string
	Description string
	Links       []Link
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
	if err := renderAll(); err != nil {
		panic(err)
	}
}

// FIXME: choose a better name
func renderAll() error {
	if err := dropCreateDir(outDir); err != nil {
		return fmt.Errorf("unable to drop-create out dir: %w", err)
	}

	if err := renderIndex(readmePath, outIndexFile); err != nil {
		return fmt.Errorf("unable to convert markdown to html: %w", err)
	}

	input, err := os.ReadFile(outIndexFile)
	if err != nil {
		return fmt.Errorf("unable to read converted html: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(input))
	if err != nil {
		return fmt.Errorf("unable to create goquery instance: %w", err)
	}

	categories, err := extractCategories(doc)
	if err != nil {
		return fmt.Errorf("unable to extract categories: %w", err)
	}

	if err := renderCategories(categories); err != nil {
		return fmt.Errorf("unable to render categories: %w", err)
	}

	if err := rewriteLinksInIndex(doc, categories); err != nil {
		return fmt.Errorf("unable to rewrite links in index: %w", err)
	}

	if err := renderSitemap(categories); err != nil {
		return fmt.Errorf("unable to render sitemap: %w", err)
	}

	for _, srcFilename := range staticFiles {
		dstFilename := filepath.Join(outDir, filepath.Base(srcFilename))
		fmt.Printf("Copy static file: %s -> %s\n", srcFilename, dstFilename)
		if err := cp.Copy(srcFilename, dstFilename); err != nil {
			return fmt.Errorf("unable to copy static file `%s` to `%s`: %w", srcFilename, dstFilename, err)
		}
	}

	return nil
}

// dropCreateDir drop and create output directory
func dropCreateDir(dir string) error {
	if err := os.RemoveAll(dir); err != nil {
		return fmt.Errorf("unable to remove dir: %w", err)
	}

	if err := mkdirAll(dir); err != nil {
		return fmt.Errorf("unable to create dir: %w", err)
	}

	return nil
}

func mkdirAll(path string) error {
	_, err := os.Stat(path)
	// NOTE: directory is exists
	if err == nil {
		return nil
	}

	// NOTE: unknown error
	if !os.IsNotExist(err) {
		return fmt.Errorf("unexpected result of dir stat: %w", err)
	}

	// NOTE: directory is not exists
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("unable to midirAll: %w", err)
	}

	return nil
}

func renderCategories(categories map[string]Category) error {
	for _, category := range categories {
		categoryDir := filepath.Join(outDir, category.Slug)
		if err := mkdirAll(categoryDir); err != nil {
			return fmt.Errorf("unable to create category dir `%s`: %w", categoryDir, err)
		}

		// FIXME: embed templates
		// FIXME: parse templates once at start
		categoryIndexFilename := filepath.Join(categoryDir, "index.html")
		fmt.Printf("Write category Index file: %s\n", categoryIndexFilename)

		buf := bytes.NewBuffer(nil)
		if err := tplCategoryIndex.Execute(buf, category); err != nil {
			return fmt.Errorf("unable to render category `%s`: %w", categoryDir, err)
		}

		// Sanitize HTML. This is not necessary, but allows to have content
		// of all html files in same style.
		{
			query, err := goquery.NewDocumentFromReader(buf)
			if err != nil {
				// FIXME: remove `unable to` from all fmt.Errorf
				return fmt.Errorf("unable to create goquery instance for `%s`: %w", categoryDir, err)
			}

			html, err := query.Html()
			if err != nil {
				return fmt.Errorf("unable to render goquery html for `%s`: %w", categoryDir, err)
			}

			if err := os.WriteFile(categoryIndexFilename, []byte(html), 0644); err != nil {
				return fmt.Errorf("unable to write category file `%s`: %w", categoryDir, err)
			}
		}
	}

	return nil
}

func renderSitemap(categories map[string]Category) error {
	// FIXME: handle error
	f, err := os.Create(outSitemapFile)
	if err != nil {
		return fmt.Errorf("unable to create sitemap file `%s`: %w", outSitemapFile, err)
	}

	fmt.Printf("Render Sitemap to: %s\n", outSitemapFile)

	if err := tplSitemap.Execute(f, categories); err != nil {
		return fmt.Errorf("unable to render sitemap: %w", err)
	}

	return nil
}

func extractCategories(doc *goquery.Document) (map[string]Category, error) {
	categories := make(map[string]Category)
	doc.
		Find("body #contents").
		NextFiltered("ul").
		Find("ul").
		Each(func(_ int, selUl *goquery.Selection) {
			selUl.
				Find("li a").
				Each(func(_ int, s *goquery.Selection) {
					selector, exists := s.Attr("href")
					if !exists {
						return
					}

					category, err := extractCategory(doc, selector)
					if err != nil {
						return
					}

					categories[selector] = *category
				})
		})

	// FIXME: handle error
	return categories, nil
}

func extractCategory(doc *goquery.Document, selector string) (*Category, error) {
	var category Category
	var err error

	doc.Find(selector).EachWithBreak(func(_ int, selCatHeader *goquery.Selection) bool {
		selDescr := selCatHeader.NextFiltered("p")
		// FIXME: bug. this would select links from all neighboring
		//   sub-categories until the next category. To prevent this we should
		//   find only first ul
		ul := selCatHeader.NextFilteredUntil("ul", "h2")

		var links []Link
		ul.Find("li").Each(func(_ int, selLi *goquery.Selection) {
			selLink := selLi.Find("a")
			url, _ := selLink.Attr("href")
			link := Link{
				Title: selLink.Text(),
				// FIXME: Title contains only title but description contains Title + description
				Description: selLi.Text(),
				Url:         url,
			}
			links = append(links, link)
		})
		// FIXME: In this case we would have an empty category in main index.html with link to 404 page.
		if len(links) == 0 {
			err = errors.New("category does not contain links")
			return false
		}

		category = Category{
			Slug:        slug.Generate(selCatHeader.Text()),
			Title:       selCatHeader.Text(),
			Description: selDescr.Text(),
			Links:       links,
		}

		return true
	})

	if err != nil {
		return nil, fmt.Errorf("unable to build a category: %w", err)
	}

	return &category, nil
}

func rewriteLinksInIndex(doc *goquery.Document, categories map[string]Category) error {
	doc.
		Find("body #content ul li ul li a").
		Each(func(_ int, s *goquery.Selection) {
			href, hrefExists := s.Attr("href")
			if !hrefExists {
				// FIXME: looks like is an error. Tag `a` in our case always
				//   	  should have `href` attr.
				return
			}

			// do not replace links if no page has been created for it
			_, catExists := categories[href]
			if !catExists {
				return
			}

			// FIXME: parse url
			uri := strings.SplitAfter(href, "#")
			if len(uri) >= 2 && uri[1] != "contents" {
				s.SetAttr("href", uri[1])
			}
		})

	fmt.Printf("Rewrite links in Index file: %s\n", outIndexFile)
	resultHtml, err := doc.Html()
	if err != nil {
		return fmt.Errorf("unable to render html: %w", err)
	}

	if err := os.WriteFile(outIndexFile, []byte(resultHtml), 0644); err != nil {
		return fmt.Errorf("unable to rewrite index file: %w", err)
	}

	return nil
}
