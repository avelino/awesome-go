// Package main contains code for generate static site.
package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	template2 "html/template"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/avelino/awesome-go/pkg/markdown"
	"github.com/avelino/awesome-go/pkg/slug"
	cp "github.com/otiai10/copy"
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

// Contributor for leaderboard
type Contributor struct {
	Login         string `json:"login"`
	AvatarURL     string `json:"avatar_url"`
	HTMLURL       string `json:"html_url"`
	Contributions int    `json:"contributions"`
}

// Source files
const readmePath = "README.md"

// This files should be copied 'as is' to outDir directory
var staticFiles = []string{
	"tmpl/assets",
	"tmpl/robots.txt",
}

// Templates
//go:embed tmpl/*.tmpl.html tmpl/*.tmpl.xml
var tplFs embed.FS

var tpl = template.Must(template.ParseFS(tplFs, "tmpl/*.tmpl.html", "tmpl/*.tmpl.xml"))

// Output files
const outDir = "out/" // NOTE: trailing slash is required

var (
	outIndexFile       = filepath.Join(outDir, "index.html")
	outSitemapFile     = filepath.Join(outDir, "sitemap.xml")
	outContributorsDir = filepath.Join(outDir, "contributors")
)

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

	// ✅ Generate leaderboard.json from git history
	if err := generateLeaderboardJSON(); err != nil {
		fmt.Printf("⚠️ Could not generate leaderboard.json: %v\n", err)
	}

	// ✅ Render contributors leaderboard page
	if err := renderContributorsPage(); err != nil {
		fmt.Printf("⚠️ Skipping contributors page: %v\n", err)
	} else {
		fmt.Println("✅ Contributors page generated successfully.")
	}

	// Copy static files
	for _, srcFilename := range staticFiles {
		dstFilename := filepath.Join(outDir, filepath.Base(srcFilename))
		fmt.Printf("Copy static file: %s -> %s\n", srcFilename, dstFilename)
		if err := cp.Copy(srcFilename, dstFilename); err != nil {
			return fmt.Errorf("copy static file `%s` to `%s`: %w", srcFilename, dstFilename, err)
		}
	}

	return nil
}

// ---------------------- Generate Leaderboard JSON ----------------------

func generateLeaderboardJSON() error {
	out, err := exec.Command("git", "shortlog", "-sne").Output()
	if err != nil {
		return err
	}

	var contributors []Contributor
	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		count, _ := strconv.Atoi(parts[0])
		login := strings.Join(parts[1:len(parts)-1], " ")
		contributors = append(contributors, Contributor{
			Login:         login,
			HTMLURL:       "", // Optionally populate GitHub profile
			AvatarURL:     "", // Optionally populate avatar
			Contributions: count,
		})
	}

	data, _ := json.MarshalIndent(contributors, "", "  ")
	return os.WriteFile("leaderboard.json", data, 0644)
}

// ---------------------- Core Build Steps ----------------------

func dropCreateDir(dir string) error {
	if err := os.RemoveAll(dir); err != nil {
		return fmt.Errorf("remove dir: %w", err)
	}
	return mkdirAll(dir)
}

func mkdirAll(path string) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	}
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("mkdirAll: %w", err)
	}
	return nil
}

func renderCategories(categories map[string]Category) error {
	for _, category := range categories {
		categoryDir := filepath.Join(outDir, category.Slug)
		if err := mkdirAll(categoryDir); err != nil {
			return fmt.Errorf("create category dir `%s`: %w", categoryDir, err)
		}

		categoryIndexFilename := filepath.Join(categoryDir, "index.html")
		fmt.Printf("Write category Index file: %s\n", categoryIndexFilename)

		buf := bytes.NewBuffer(nil)
		if err := tpl.Lookup("category-index.tmpl.html").Execute(buf, category); err != nil {
			return fmt.Errorf("render category `%s`: %w", categoryDir, err)
		}

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
	return nil
}

func renderSitemap(categories map[string]Category) error {
	f, err := os.Create(outSitemapFile)
	if err != nil {
		return fmt.Errorf("create sitemap file `%s`: %w", outSitemapFile, err)
	}
	fmt.Printf("Render Sitemap to: %s\n", outSitemapFile)
	return tpl.Lookup("sitemap.tmpl.xml").Execute(f, categories)
}

// ---------------------- Contributors Page ----------------------

func renderContributorsPage() error {
	data, err := os.ReadFile("leaderboard.json")
	if err != nil {
		return fmt.Errorf("read leaderboard.json: %w", err)
	}

	var contributors []Contributor
	if err := json.Unmarshal(data, &contributors); err != nil {
		return fmt.Errorf("parse leaderboard.json: %w", err)
	}

	if len(contributors) == 0 {
		return fmt.Errorf("no contributors found")
	}

	if err := mkdirAll(outContributorsDir); err != nil {
		return fmt.Errorf("create contributors dir: %w", err)
	}

	outFile := filepath.Join(outContributorsDir, "index.html")
	f, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("create contributors file: %w", err)
	}
	defer f.Close()

	fmt.Println("Render Contributors Leaderboard →", outFile)

	return tpl.Lookup("contributors.tmpl.html").Execute(f, map[string]interface{}{
		"Contributors": contributors,
	})
}

// ---------------------- Category Extraction ----------------------

func extractCategories(doc *goquery.Document) (map[string]Category, error) {
	categories := make(map[string]Category)
	var rootErr error

	doc.Find("body #contents").
		NextFiltered("ul").
		Find("ul").
		EachWithBreak(func(_ int, selUl *goquery.Selection) bool {
			if rootErr != nil {
				return false
			}
			selUl.Find("li a").EachWithBreak(func(_ int, s *goquery.Selection) bool {
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
		return nil, rootErr
	}
	return categories, nil
}

func extractCategory(doc *goquery.Document, selector string) (*Category, error) {
	var category Category
	var err error

	doc.Find(selector).EachWithBreak(func(_ int, selCatHeader *goquery.Selection) bool {
		selDescr := selCatHeader.NextFiltered("p")
		ul := selCatHeader.NextFilteredUntil("ul", "h2")

		var links []Link
		ul.Find("li").Each(func(_ int, selLi *goquery.Selection) {
			selLink := selLi.Find("a")
			url, _ := selLink.Attr("href")
			link := Link{
				Title:       selLink.Text(),
				Description: selLi.Text(),
				URL:         url,
			}
			links = append(links, link)
		})

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
		return nil, err
	}
	return &category, nil
}

func rewriteLinksInIndex(doc *goquery.Document, categories map[string]Category) error {
	doc.Find("body #content ul li ul li a").Each(func(_ int, s *goquery.Selection) {
		href, hrefExists := s.Attr("href")
		if !hrefExists {
			return
		}
		_, catExists := categories[href]
		if !catExists {
			return
		}
		linkURL, err := url.Parse(href)
		if err == nil && linkURL.Fragment != "" && linkURL.Fragment != "contents" {
			s.SetAttr("href", linkURL.Fragment)
		}
	})

	fmt.Printf("Rewrite links in Index file: %s\n", outIndexFile)
	resultHTML, err := doc.Html()
	if err != nil {
		return err
	}
	return os.WriteFile(outIndexFile, []byte(resultHTML), 0644)
}

// ---------------------- Markdown Conversion ----------------------

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
	if err := tpl.Lookup("index.tmpl.html").Execute(f, data); err != nil {
		return err
	}
	return f.Close()
}
