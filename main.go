// Package main contains code for generate static site.
package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/avelino/awesome-go/pkg/markdown"
	cp "github.com/otiai10/copy"
	template2 "html/template"
	"net/url"
	"os"
	"path/filepath"
	"text/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/avelino/awesome-go/pkg/slug"
)

// Link contains info about awesome url
type Link struct {
	Title       string
	URL         string
	Description string
}

// Category describe link category
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
	"tmpl/robots.txt",
}

// TODO: embed
// Templates
var tplIndex = template.Must(template.ParseFiles("tmpl/index.tmpl.html"))
var tplCategoryIndex = template.Must(template.ParseFiles("tmpl/category-index.tmpl.html"))
var tplSitemap = template.Must(template.ParseFiles("tmpl/sitemap.tmpl.xml"))

// Output files
const outDir = "out/" // NOTE: trailing slash is required

var outIndexFile = filepath.Join(outDir, "index.html")
var outSitemapFile = filepath.Join(outDir, "sitemap.xml")

func main() {
	if err := buildStaticSite(); err != nil {
		panic(err)
	}
}

func buildStaticSite() error {
	if err := dropCreateDir(outDir); err != nil {
		return fmt.Errorf("drop-create out dir: %w", err)
	}

	if err := renderIndex(readmePath, outIndexFile); err != nil {
		return fmt.Errorf("convert markdown to html: %w", err)
	}

	input, err := os.ReadFile(outIndexFile)
	if err != nil {
		return fmt.Errorf("read converted html: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(input))
	if err != nil {
		return fmt.Errorf("create goquery instance: %w", err)
	}

	categories, err := extractCategories(doc)
	if err != nil {
		return fmt.Errorf("extract categories: %w", err)
	}

	if err := renderCategories(categories); err != nil {
		return fmt.Errorf("render categories: %w", err)
	}

	if err := rewriteLinksInIndex(doc, categories); err != nil {
		return fmt.Errorf("rewrite links in index: %w", err)
	}

	if err := renderSitemap(categories); err != nil {
		return fmt.Errorf("render sitemap: %w", err)
	}

	for _, srcFilename := range staticFiles {
		dstFilename := filepath.Join(outDir, filepath.Base(srcFilename))
		fmt.Printf("Copy static file: %s -> %s\n", srcFilename, dstFilename)
		if err := cp.Copy(srcFilename, dstFilename); err != nil {
			return fmt.Errorf("copy static file `%s` to `%s`: %w", srcFilename, dstFilename, err)
		}
	}

	return nil
}

// dropCreateDir drop and create output directory
func dropCreateDir(dir string) error {
	if err := os.RemoveAll(dir); err != nil {
		return fmt.Errorf("remove dir: %w", err)
	}

	if err := mkdirAll(dir); err != nil {
		return fmt.Errorf("create dir: %w", err)
	}

	return nil
}

func mkdirAll(path string) error {
	_, err := os.Stat(path)
	// directory is exists
	if err == nil {
		return nil
	}

	// unexpected error
	if !os.IsNotExist(err) {
		return fmt.Errorf("unexpected result of dir stat: %w", err)
	}

	// directory is not exists
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("midirAll: %w", err)
	}

	return nil
}

func renderCategories(categories map[string]Category) error {
	for _, category := range categories {
		categoryDir := filepath.Join(outDir, category.Slug)
		if err := mkdirAll(categoryDir); err != nil {
			return fmt.Errorf("create category dir `%s`: %w", categoryDir, err)
		}

		// FIXME: embed templates
		categoryIndexFilename := filepath.Join(categoryDir, "index.html")
		fmt.Printf("Write category Index file: %s\n", categoryIndexFilename)

		buf := bytes.NewBuffer(nil)
		if err := tplCategoryIndex.Execute(buf, category); err != nil {
			return fmt.Errorf("render category `%s`: %w", categoryDir, err)
		}

		// Sanitize HTML. This is not necessary, but allows to have content
		// of all html files in same style.
		{
			doc, err := goquery.NewDocumentFromReader(buf)
			if err != nil {
				return fmt.Errorf("create goquery instance for `%s`: %w", categoryDir, err)
			}

			html, err := doc.Html()
			if err != nil {
				return fmt.Errorf("render goquery html for `%s`: %w", categoryDir, err)
			}

			if err := os.WriteFile(categoryIndexFilename, []byte(html), 0644); err != nil {
				return fmt.Errorf("write category file `%s`: %w", categoryDir, err)
			}
		}
	}

	return nil
}

func renderSitemap(categories map[string]Category) error {
	f, err := os.Create(outSitemapFile)
	if err != nil {
		return fmt.Errorf("create sitemap file `%s`: %w", outSitemapFile, err)
	}

	fmt.Printf("Render Sitemap to: %s\n", outSitemapFile)

	if err := tplSitemap.Execute(f, categories); err != nil {
		return fmt.Errorf("render sitemap: %w", err)
	}

	return nil
}

func extractCategories(doc *goquery.Document) (map[string]Category, error) {
	categories := make(map[string]Category)
	var rootErr error

	doc.
		Find("body #contents").
		NextFiltered("ul").
		Find("ul").
		EachWithBreak(func(_ int, selUl *goquery.Selection) bool {
			if rootErr != nil {
				return false
			}

			selUl.
				Find("li a").
				EachWithBreak(func(_ int, s *goquery.Selection) bool {
					selector, exists := s.Attr("href")
					if !exists {
						return true
					}

					category, err := extractCategory(doc, selector)
					if err != nil {
						rootErr = fmt.Errorf("extract category: %w", err)
						return false
					}

					categories[selector] = *category

					return true
				})

			return true
		})

	if rootErr != nil {
		return nil, fmt.Errorf("extract categories: %w", rootErr)
	}

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
				// FIXME(kazhuravlev): Title contains only title but
				// 	description contains Title + description
				Description: selLi.Text(),
				URL:         url,
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
		return nil, fmt.Errorf("build a category: %w", err)
	}

	return &category, nil
}

func rewriteLinksInIndex(doc *goquery.Document, categories map[string]Category) error {
	var iterErr error
	doc.
		Find("body #content ul li ul li a").
		EachWithBreak(func(_ int, s *goquery.Selection) bool {
			href, hrefExists := s.Attr("href")
			if !hrefExists {
				// FIXME: looks like is an error. Tag `a` in our case always
				//   	  should have `href` attr.
				return true
			}

			// do not replace links if no page has been created for it
			_, catExists := categories[href]
			if !catExists {
				return true
			}

			linkURL, err := url.Parse(href)
			if err != nil {
				iterErr = err
				return false
			}

			if linkURL.Fragment != "" && linkURL.Fragment != "contents" {
				s.SetAttr("href", linkURL.Fragment)
			}

			return true
		})

	if iterErr != nil {
		return iterErr
	}

	fmt.Printf("Rewrite links in Index file: %s\n", outIndexFile)
	resultHTML, err := doc.Html()
	if err != nil {
		return fmt.Errorf("render html: %w", err)
	}

	if err := os.WriteFile(outIndexFile, []byte(resultHTML), 0644); err != nil {
		return fmt.Errorf("rewrite index file: %w", err)
	}

	return nil
}

// renderIndex generate site html (index.html) from markdown file
func renderIndex(srcFilename, outFilename string) error {
	input, err := os.ReadFile(srcFilename)
	if err != nil {
		return err
	}

	body, err := markdown.ToHTML(input)
	if err != nil {
		return err
	}

	f, err := os.Create(outFilename)
	if err != nil {
		return err
	}

	fmt.Printf("Write Index file: %s\n", outIndexFile)
	data := map[string]interface{}{
		"Body": template2.HTML(body),
	}
	if err := tplIndex.Execute(f, data); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("close index file: %w", err)
	}

	return nil
}
