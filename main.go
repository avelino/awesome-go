// Package main contains code for generating a static site.
package main

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	template2 "html/template"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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

// Static files to copy
var staticFiles = []string{
	"tmpl/assets",
	"tmpl/robots.txt",
}

// Templates
//go:embed tmpl/*.tmpl.html tmpl/*.tmpl.xml
var tplFs embed.FS

var tpl = template.Must(template.New("").Funcs(template2.FuncMap{
	"min": func(a, b int) int {
		if a < b {
			return a
		}
		return b
	},
	"add": func(nums ...int) int {
		sum := 0
		for _, n := range nums {
			sum += n
		}
		return sum
	},
}).ParseFS(tplFs, "tmpl/*.tmpl.html", "tmpl/*.tmpl.xml"))


// Output files
const outDir = "out/"

var (
	outIndexFile       = filepath.Join(outDir, "index.html")
	outSitemapFile     = filepath.Join(outDir, "sitemap.xml")
	outContributorsDir = filepath.Join(outDir, "contributors")
	leaderboardJSON    = filepath.Join(outDir, "leaderboard.json")
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

	if err := generateLeaderboardJSON(); err != nil {
		fmt.Printf("⚠️ Could not generate leaderboard.json: %v\n", err)
	}

	if err := renderContributorsPage(); err != nil {
		fmt.Printf("⚠️ Skipping contributors page: %v\n", err)
	} else {
		fmt.Println("✅ Contributors page generated successfully.")
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

// ---------------------- Generate Leaderboard JSON via GitHub API ----------------------

func generateLeaderboardJSON() error {
	repoOwner := "avelino"
	repoName := "awesome-go"

	var contributors []Contributor
	page := 1

	for {
		apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contributors?per_page=100&anon=1&page=%d", repoOwner, repoName, page)
		req, _ := http.NewRequestWithContext(context.Background(), "GET", apiURL, nil)

		if token := os.Getenv("GITHUB_TOKEN"); token != "" {
			req.Header.Set("Authorization", "token "+token)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("fetch contributors: %w", err)
		}
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("github api error: status %d", resp.StatusCode)
		}

		var pageContributors []Contributor
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(&pageContributors); err != nil {
			return fmt.Errorf("decode github contributors: %w", err)
		}

		if len(pageContributors) == 0 {
			break
		}

		for i, c := range pageContributors {
			if c.Login == "" {
				pageContributors[i].Login = "Anonymous Contributor"
				pageContributors[i].HTMLURL = "#"
			}
			if c.AvatarURL == "" {
				// Use GitHub-style default identicon based on login
				pageContributors[i].AvatarURL = fmt.Sprintf(
					"https://avatars.githubusercontent.com/%s?v=4",
					url.PathEscape(pageContributors[i].Login),
				)
			}
		}		

		contributors = append(contributors, pageContributors...)
		page++
	}

	if len(contributors) == 0 {
		return errors.New("no contributors found from GitHub API")
	}

	data, _ := json.MarshalIndent(contributors, "", "  ")
	if err := os.WriteFile(leaderboardJSON, data, 0644); err != nil {
		return fmt.Errorf("write leaderboard.json: %w", err)
	}

	fmt.Printf("✅ Generated leaderboard.json with %d contributors → %s\n", len(contributors), leaderboardJSON)
	return nil
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
	// Read leaderboard JSON
	data, err := os.ReadFile(leaderboardJSON)
	if err != nil {
		return fmt.Errorf("read leaderboard.json: %w", err)
	}

	var contributors []Contributor
	if err := json.Unmarshal(data, &contributors); err != nil {
		return fmt.Errorf("parse leaderboard.json: %w", err)
	}

	if contributors == nil {
		contributors = []Contributor{}
	}

	// Split contributors into sections
	top3 := contributors
	if len(top3) > 3 {
		top3 = contributors[:3]
	}

	top10 := []Contributor{}
	if len(contributors) > 3 {
		if len(contributors) > 13 {
			top10 = contributors[3:13]
		} else {
			top10 = contributors[3:]
		}
	}

	top100 := []Contributor{}
	if len(contributors) > 13 {
		if len(contributors) > 113 {
			top100 = contributors[13:113]
		} else {
			top100 = contributors[13:]
		}
	}

	others := []Contributor{}
	if len(contributors) > 113 {
		others = contributors[113:]
	}

	// Ensure output directory exists
	if err := mkdirAll(outContributorsDir); err != nil {
		return fmt.Errorf("create contributors dir: %w", err)
	}

	// Create HTML file
	outFile := filepath.Join(outContributorsDir, "index.html")
	f, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("create contributors file: %w", err)
	}
	defer f.Close()

	fmt.Println("Render Contributors Leaderboard →", outFile)

	// Execute template
	return tpl.Lookup("contributors.tmpl.html").Execute(f, map[string]interface{}{
		"Top3":    top3,
		"Top10":   top10,
		"Top100":  top100,
		"Others":  others,
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

	// Pass an empty Contributors slice to avoid template panics
	data := map[string]interface{}{
		"Body":         template2.HTML(body),
		"Contributors": []Contributor{},
	}

	if err := tpl.Lookup("index.tmpl.html").Execute(f, data); err != nil {
		return err
	}
	return f.Close()
}
