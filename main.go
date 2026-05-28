// Package main contains code for generate static site.
package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	template2 "html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/avelino/awesome-go/pkg/markdown"
	cp "github.com/otiai10/copy"

	"github.com/PuerkitoBio/goquery"
	"github.com/avelino/awesome-go/pkg/slug"
)

// Link contains info about awesome url
type Link struct {
	Title       string
	URL         string
	Description string
	ProjectSlug string // non-empty if project page exists → internal link
}

// Category describe link category
type Category struct {
	Title       string
	Slug        string
	Description string
	Links       []Link
}

// RepoMeta holds metadata fetched from GitHub/GitLab API
type RepoMeta struct {
	Stars      int      `json:"stars"`
	Forks      int      `json:"forks"`
	License    string   `json:"license"`
	Language   string   `json:"language"`
	Topics     []string `json:"topics"`
	LastPush   string   `json:"last_push"`
	OpenIssues int      `json:"open_issues"`
	Archived   bool     `json:"archived"`
	FetchedAt  string   `json:"fetched_at"`
}

// Project represents an individual project page
type Project struct {
	Title         string
	URL           string
	Description   string
	Slug          string
	Host          string // "github" or "gitlab"
	Owner         string
	Repo          string
	Meta          *RepoMeta
	CategoryTitle string
	CategorySlug  string
	Related       []Link
}

// SitemapData holds data for sitemap generation
type SitemapData struct {
	Categories map[string]Category
	Projects   []*Project
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

var tpl = template.Must(
	template.New("").Funcs(template.FuncMap{
		"now": func() string { return time.Now().Format("2006-01-02") },
		"jsonEscape": func(s string) string {
			b, _ := json.Marshal(s)
			if len(b) < 2 {
				return ""
			}
			return string(b[1 : len(b)-1])
		},
	}).ParseFS(tplFs, "tmpl/*.tmpl.html", "tmpl/*.tmpl.xml"),
)

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

	projects := buildProjects(categories)
	if err := fetchProjectMeta(projects); err != nil {
		return fmt.Errorf("fetch project metadata: %w", err)
	}

	if err := renderCategories(categories); err != nil {
		return fmt.Errorf("render categories: %w", err)
	}

	if err := renderProjects(projects); err != nil {
		return fmt.Errorf("render projects: %w", err)
	}

	if err := rewriteLinksInIndex(doc, categories); err != nil {
		return fmt.Errorf("rewrite links in index: %w", err)
	}

	if err := renderSitemap(categories, projects); err != nil {
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
		if err := tpl.Lookup("category-index.tmpl.html").Execute(buf, category); err != nil {
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

func renderSitemap(categories map[string]Category, projects []*Project) error {
	f, err := os.Create(outSitemapFile)
	if err != nil {
		return fmt.Errorf("create sitemap file `%s`: %w", outSitemapFile, err)
	}

	fmt.Printf("Render Sitemap to: %s\n", outSitemapFile)

	data := SitemapData{
		Categories: categories,
		Projects:   projects,
	}

	if err := tpl.Lookup("sitemap.tmpl.xml").Execute(f, data); err != nil {
		return fmt.Errorf("render sitemap: %w", err)
	}

	return nil
}

func extractCategories(doc *goquery.Document) (map[string]Category, error) {
	categories := make(map[string]Category)
	var rootErr error

	contentsHeading := doc.Find("body #contents")
	// Support both direct <ul> sibling and <ul> wrapped in <details>
	contentsList := contentsHeading.NextFiltered("ul")
	if contentsList.Length() == 0 {
		contentsList = contentsHeading.NextFiltered("details").Find("ul").First()
	}

	contentsList.
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
						// Skip entries without links (e.g. #contents, #awesome-go)
						return true
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
		// sub-categories until the next category. To prevent this we should
		// find only first ul
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
	if err := tpl.Lookup("index.tmpl.html").Execute(f, data); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("close index file: %w", err)
	}

	return nil
}

// parseRepoURL extracts host, owner, and repo from a GitHub or GitLab URL
func parseRepoURL(rawURL string) (host, owner, repo string, ok bool) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", "", "", false
	}

	parts := strings.Split(strings.Trim(u.Path, "/"), "/")

	switch u.Hostname() {
	case "github.com":
		if len(parts) < 2 {
			return "", "", "", false
		}
		return "github", parts[0], parts[1], true
	case "gitlab.com":
		if len(parts) < 2 {
			return "", "", "", false
		}
		return "gitlab", parts[0], parts[len(parts)-1], true
	}

	return "", "", "", false
}

// buildProjects creates Project structs from category links and sets internal link slugs
func buildProjects(categories map[string]Category) []*Project {
	var projects []*Project

	for key, cat := range categories {
		for i, link := range cat.Links {
			host, owner, repo, ok := parseRepoURL(link.URL)
			if !ok {
				continue
			}

			projectSlug := slug.Generate(repo + "-" + owner + "-" + host)
			cat.Links[i].ProjectSlug = projectSlug

			p := &Project{
				Title:         link.Title,
				URL:           link.URL,
				Description:   link.Description,
				Slug:          projectSlug,
				Host:          host,
				Owner:         owner,
				Repo:          repo,
				CategoryTitle: cat.Title,
				CategorySlug:  cat.Slug,
			}
			projects = append(projects, p)
		}
		categories[key] = cat
	}

	// Populate Related: group by category, pick up to 10 siblings
	byCat := make(map[string][]*Project)
	for _, p := range projects {
		byCat[p.CategorySlug] = append(byCat[p.CategorySlug], p)
	}
	for _, p := range projects {
		var related []Link
		for _, s := range byCat[p.CategorySlug] {
			if s.Slug == p.Slug {
				continue
			}
			related = append(related, Link{
				Title:       s.Title,
				URL:         s.URL,
				Description: s.Description,
				ProjectSlug: s.Slug,
			})
			if len(related) >= 10 {
				break
			}
		}
		p.Related = related
	}

	fmt.Printf("Built %d project pages\n", len(projects))
	return projects
}

// fetchProjectMeta fetches metadata from GitHub/GitLab APIs with file-per-repo caching
func fetchProjectMeta(projects []*Project) error {
	if os.Getenv("AWESOME_SKIP_FETCH") != "" {
		log.Println("AWESOME_SKIP_FETCH set, skipping metadata fetch")
		return nil
	}

	token := os.Getenv("GITHUB_TOKEN")
	client := &http.Client{Timeout: 10 * time.Second}
	fetched := 0

	for _, p := range projects {
		cached, err := readCachedMeta(p)
		if err == nil && cached != nil {
			p.Meta = cached
			continue
		}

		var meta *RepoMeta
		switch p.Host {
		case "github":
			meta = fetchGitHubMeta(client, p.Owner, p.Repo, token)
		case "gitlab":
			meta = fetchGitLabMeta(client, p.URL)
		}

		if meta != nil {
			meta.FetchedAt = time.Now().Format("2006-01-02")
			p.Meta = meta
			if err := writeCachedMeta(p, meta); err != nil {
				log.Printf("warning: cache write failed for %s/%s: %v", p.Owner, p.Repo, err)
			}
			fetched++
		}

		time.Sleep(50 * time.Millisecond)
	}

	fmt.Printf("Fetched metadata for %d projects (%d from cache)\n", fetched, len(projects)-fetched)
	return nil
}

func cacheFilePath(p *Project) string {
	return filepath.Join(".cache", "repos", p.Owner, p.Repo+".json")
}

func readCachedMeta(p *Project) (*RepoMeta, error) {
	data, err := os.ReadFile(cacheFilePath(p))
	if err != nil {
		return nil, err
	}

	var meta RepoMeta
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, err
	}

	fetchedAt, err := time.Parse("2006-01-02", meta.FetchedAt)
	if err != nil {
		return nil, err
	}
	if time.Since(fetchedAt) > 7*24*time.Hour {
		return nil, fmt.Errorf("cache stale")
	}

	return &meta, nil
}

func writeCachedMeta(p *Project, meta *RepoMeta) error {
	path := cacheFilePath(p)
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func fetchGitHubMeta(client *http.Client, owner, repo, token string) *RepoMeta {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	var ghRepo struct {
		StargazersCount int `json:"stargazers_count"`
		ForksCount      int `json:"forks_count"`
		License         *struct {
			SpdxID string `json:"spdx_id"`
		} `json:"license"`
		Language        string   `json:"language"`
		Topics          []string `json:"topics"`
		PushedAt        string   `json:"pushed_at"`
		OpenIssuesCount int      `json:"open_issues_count"`
		Archived        bool     `json:"archived"`
	}

	if err := json.Unmarshal(body, &ghRepo); err != nil {
		return nil
	}

	license := ""
	if ghRepo.License != nil {
		license = ghRepo.License.SpdxID
	}

	lastPush := ""
	if ghRepo.PushedAt != "" {
		if t, err := time.Parse(time.RFC3339, ghRepo.PushedAt); err == nil {
			lastPush = t.Format("2006-01-02")
		}
	}

	return &RepoMeta{
		Stars:      ghRepo.StargazersCount,
		Forks:      ghRepo.ForksCount,
		License:    license,
		Language:   ghRepo.Language,
		Topics:     ghRepo.Topics,
		LastPush:   lastPush,
		OpenIssues: ghRepo.OpenIssuesCount,
		Archived:   ghRepo.Archived,
	}
}

func fetchGitLabMeta(client *http.Client, projectURL string) *RepoMeta {
	u, err := url.Parse(projectURL)
	if err != nil {
		return nil
	}

	path := strings.Trim(u.Path, "/")
	// Remove tree/branch paths (e.g., /-/tree/main/...)
	if idx := strings.Index(path, "/-/"); idx != -1 {
		path = path[:idx]
	}

	apiURL := fmt.Sprintf("https://gitlab.com/api/v4/projects/%s", url.PathEscape(path))
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	var glProject struct {
		StarCount       int `json:"star_count"`
		ForksCount      int `json:"forks_count"`
		License         *struct {
			Key string `json:"key"`
		} `json:"license"`
		Topics          []string `json:"topics"`
		LastActivityAt  string   `json:"last_activity_at"`
		OpenIssuesCount int      `json:"open_issues_count"`
		Archived        bool     `json:"archived"`
	}

	if err := json.Unmarshal(body, &glProject); err != nil {
		return nil
	}

	license := ""
	if glProject.License != nil {
		license = glProject.License.Key
	}

	lastPush := ""
	if glProject.LastActivityAt != "" {
		if t, err := time.Parse(time.RFC3339, glProject.LastActivityAt); err == nil {
			lastPush = t.Format("2006-01-02")
		}
	}

	return &RepoMeta{
		Stars:      glProject.StarCount,
		Forks:      glProject.ForksCount,
		License:    license,
		Language:   "Go",
		Topics:     glProject.Topics,
		LastPush:   lastPush,
		OpenIssues: glProject.OpenIssuesCount,
		Archived:   glProject.Archived,
	}
}

func renderProjects(projects []*Project) error {
	for _, p := range projects {
		projectDir := filepath.Join(outDir, p.CategorySlug, p.Slug)
		if err := mkdirAll(projectDir); err != nil {
			return fmt.Errorf("create project dir %s: %w", projectDir, err)
		}

		projectFile := filepath.Join(projectDir, "index.html")

		buf := bytes.NewBuffer(nil)
		if err := tpl.Lookup("project.tmpl.html").Execute(buf, p); err != nil {
			return fmt.Errorf("render project %s: %w", p.Slug, err)
		}

		doc, err := goquery.NewDocumentFromReader(buf)
		if err != nil {
			return fmt.Errorf("goquery for project %s: %w", p.Slug, err)
		}

		html, err := doc.Html()
		if err != nil {
			return fmt.Errorf("html for project %s: %w", p.Slug, err)
		}

		if err := os.WriteFile(projectFile, []byte(html), 0644); err != nil {
			return fmt.Errorf("write project file %s: %w", p.Slug, err)
		}
	}

	fmt.Printf("Rendered %d project pages\n", len(projects))
	return nil
}
