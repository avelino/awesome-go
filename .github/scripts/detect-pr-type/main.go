// Command detect-pr-type tells the PR Quality Checks workflow whether a pull
// request adds a package to README.md, so the repository checks (forge link,
// pkg.go.dev, go.mod, ...) only run when there is a new link to check.
// Removal-only and formatting-only README changes skip them.
//
// It reads the PR file list through the API and never checks out or runs PR
// code, so it is safe under pull_request_target.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type prFile struct {
	Filename string `json:"filename"`
	Patch    string `json:"patch"`
}

var httpClient = &http.Client{Timeout: 30 * time.Second}

func main() {
	repo := os.Getenv("GITHUB_REPOSITORY")
	number := prNumber()
	if repo == "" || number == 0 {
		fail("GITHUB_REPOSITORY or the pull request number is missing")
	}

	files, err := listFiles(repo, number)
	if err != nil {
		fail(err.Error())
	}

	isPackage, reason := decide(files)
	fmt.Println(reason)
	setOutput("is_package_pr", fmt.Sprintf("%t", isPackage))
}

// decide maps the changed files to the is_package_pr output. When in doubt it
// answers true, so the quality checks run.
func decide(files []prFile) (bool, string) {
	for _, f := range files {
		if f.Filename != "README.md" {
			continue
		}
		if strings.TrimSpace(f.Patch) == "" {
			return true, "README.md changed but GitHub sent no patch (diff too large), running the quality checks"
		}
		c := classify(f.Patch)
		switch c.Kind {
		case kindPackage:
			return true, fmt.Sprintf("README.md adds %d new link(s): %s, this is a package PR", len(c.Added), strings.Join(c.Added, ", "))
		case kindRemoval:
			return false, fmt.Sprintf("README.md only removes link(s): %s, skipping quality checks", strings.Join(c.Removed, ", "))
		default:
			return false, "README.md changes formatting only (no link added or removed), skipping quality checks"
		}
	}
	return false, "README.md not modified, skipping quality checks"
}

func prNumber() int {
	path := os.Getenv("GITHUB_EVENT_PATH")
	if path == "" {
		return 0
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return 0
	}
	var ev struct {
		PullRequest struct {
			Number int `json:"number"`
		} `json:"pull_request"`
	}
	_ = json.Unmarshal(data, &ev)
	return ev.PullRequest.Number
}

func listFiles(repo string, number int) ([]prFile, error) {
	var all []prFile
	// The API lists at most 3000 files, 100 per page.
	for page := 1; page <= 30; page++ {
		url := fmt.Sprintf("https://api.github.com/repos/%s/pulls/%d/files?per_page=100&page=%d", repo, number, page)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("User-Agent", "awesome-go-detect-pr-type")
		req.Header.Set("Accept", "application/vnd.github+json")
		if token := os.Getenv("GITHUB_TOKEN"); token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		}
		resp, err := httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("listing PR files: HTTP %d", resp.StatusCode)
		}
		var files []prFile
		if err := json.Unmarshal(body, &files); err != nil {
			return nil, err
		}
		all = append(all, files...)
		if len(files) < 100 {
			break
		}
	}
	return all, nil
}

// setOutput writes a single-line value; is_package_pr is only ever true or false.
func setOutput(name, value string) {
	path := os.Getenv("GITHUB_OUTPUT")
	if path == "" {
		fmt.Printf("%s=%s\n", name, value)
		return
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil {
		fail(err.Error())
	}
	defer f.Close()
	fmt.Fprintf(f, "%s=%s\n", name, value)
}

func fail(msg string) {
	fmt.Fprintln(os.Stderr, "detect-pr-type:", msg)
	os.Exit(1)
}
