package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"testing"
	"text/template"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/oauth2"
)

const issueTemplateContent = `
{{range .}}
- [ ] {{.}}
{{end}}
`

var issueTemplate = template.Must(template.New("issue").Parse(issueTemplateContent))

var reGithubRepo = regexp.MustCompile(`^https://github\.com/([-a-zA-Z0-9_.]+)/([-a-zA-Z0-9_.]+)/?$`)
var githubGETREPO = "https://api.github.com/repos/%s/%s"
var githubGETCOMMITS = "https://api.github.com/repos/%s/%s/commits"
var githubPOSTISSUES = "https://api.github.com/repos/avelino/awesome-go/issues"
var awesomeGoGETISSUES = "https://api.github.com/repos/avelino/awesome-go/issues"

var numberOfYears = 1
var timeNow = time.Now()
var issueTitle = fmt.Sprintf("Investigate repositories with more than 1 year without update - %s", timeNow.Format(time.DateOnly))

const deadLinkMessage = " this repository might no longer exist! (status code 404/410 returned)"
const movedPermanently = " status code 301 received"
const status302 = " status code 302 received"
const archived = " repository has been archived"
const noRecentCommits = " repository has not received any commits in over 1 year"

var LIMIT = 10
var ctr = 0

type tokenSource struct {
	AccessToken string
}

type issue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (t *tokenSource) Token() (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: t.AccessToken,
	}, nil
}

func getRepositoriesFromBody(body string) []string {
	links := strings.Split(body, "- ")
	for i, link := range links {
		link = strings.ReplaceAll(link, "\r", "")
		link = strings.ReplaceAll(link, "[ ]", "")
		link = strings.ReplaceAll(link, "[x]", "")
		link = strings.ReplaceAll(link, " ", "")
		link = strings.ReplaceAll(link, "\n", "")
		link = strings.ReplaceAll(link, deadLinkMessage, "")
		link = strings.ReplaceAll(link, movedPermanently, "")
		link = strings.ReplaceAll(link, status302, "")
		link = strings.ReplaceAll(link, archived, "")
		link = strings.ReplaceAll(link, noRecentCommits, "")
		links[i] = link
	}
	return links
}

func generateIssueBody(t *testing.T, repositories []string) (string, error) {
	t.Helper()
	buf := bytes.NewBuffer(nil)
	err := issueTemplate.Execute(buf, repositories)
	requireNoErr(t, err, "Failed to generate template")
	return buf.String(), nil
}

func createIssue(t *testing.T, staleRepos []string, client *http.Client) {
	t.Helper()
	if len(staleRepos) == 0 {
		t.Log("NO STALE REPOSITORIES")
		return
	}
	body, err := generateIssueBody(t, staleRepos)
	requireNoErr(t, err, "failed to generate issue body")

	newIssue := &issue{
		Title: issueTitle,
		Body:  body,
	}
	buf := bytes.NewBuffer(nil)
	requireNoErr(t, json.NewEncoder(buf).Encode(newIssue), "failed to encode json req")

	req, err := http.NewRequest(http.MethodPost, githubPOSTISSUES, buf)
	requireNoErr(t, err, "failed to create request")

	resp, roundTripErr := client.Do(req)
	requireNoErr(t, roundTripErr, "failed to send request")
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		t.Fatalf("create issue failed: status %d, body: %s", resp.StatusCode, string(respBody))
	}
}

func getAllFlaggedRepositories(t *testing.T, client *http.Client) map[string]bool {
	t.Helper()
	req, err := http.NewRequest(http.MethodGet, awesomeGoGETISSUES, nil)
	requireNoErr(t, err, "failed to create request")

	res, err := client.Do(req)
	requireNoErr(t, err, "failed to send request")
	defer res.Body.Close()

	if res.StatusCode == http.StatusForbidden || res.StatusCode == http.StatusUnauthorized {
		t.Fatalf("getAllFlaggedRepositories: rate limit or auth failure (status %d)", res.StatusCode)
	}

	var issues []issue
	requireNoErr(t, json.NewDecoder(res.Body).Decode(&issues), "failed to unmarshal response")

	addressedRepositories := make(map[string]bool)
	for _, issue := range issues {
		if issue.Title != issueTitle {
			continue
		}
		repos := getRepositoriesFromBody(issue.Body)
		for _, repo := range repos {
			addressedRepositories[repo] = true
		}
	}
	return addressedRepositories
}

func checkRepoAvailability(t *testing.T, toRun bool, href string, client *http.Client) (warnings []string, aborted bool) {
	t.Helper()
	if !toRun {
		return nil, false
	}

	matches := reGithubRepo.FindStringSubmatch(href)
	if len(matches) < 3 {
		t.Logf("Invalid github repo url format: %s", href)
		return nil, false
	}
	owner, repo := matches[1], matches[2]
	apiCall := fmt.Sprintf(githubGETREPO, owner, repo)

	req, err := http.NewRequest(http.MethodGet, apiCall, nil)
	if err != nil {
		t.Logf("Failed to create request for repository %s\n", href)
		return nil, false
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Logf("Failed at repository %s\n", href)
		return nil, false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden || resp.StatusCode == http.StatusUnauthorized {
		t.Logf("GitHub API rate limit exceeded or unauthorized (Status: %d). Signal abort.", resp.StatusCode)
		return nil, true
	}

	var repoResp struct {
		Archived bool `json:"archived"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&repoResp); err != nil {
		return nil, false
	}

	if resp.StatusCode == http.StatusMovedPermanently {
		warnings = append(warnings, href+movedPermanently)
		t.Logf("%s returned %d", href, resp.StatusCode)
	} else if resp.StatusCode == http.StatusFound {
		warnings = append(warnings, href+status302)
		t.Logf("%s returned %d", href, resp.StatusCode)
	} else if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusGone {
		warnings = append(warnings, href+deadLinkMessage)
		t.Logf("%s might not exist! (Status: %d)", href, resp.StatusCode)
	}

	if repoResp.Archived && len(warnings) == 0 {
		warnings = append(warnings, href+archived)
		t.Logf("%s is archived!", href)
	}

	return warnings, false
}

func checkRepoCommitActivity(t *testing.T, toRun bool, href string, client *http.Client) (warnings []string, aborted bool) {
	t.Helper()
	if !toRun {
		return nil, false
	}

	matches := reGithubRepo.FindStringSubmatch(href)
	if len(matches) < 3 {
		return nil, false
	}
	owner, repo := matches[1], matches[2]
	apiCall := fmt.Sprintf(githubGETCOMMITS, owner, repo)

	req, err := http.NewRequest(http.MethodGet, apiCall, nil)
	if err != nil {
		t.Logf("Failed to create request for repository %s\n", href)
		return nil, false
	}

	since := timeNow.Add(-1 * time.Duration(numberOfYears) * 365 * 24 * time.Hour)
	sinceQuery := since.Format(time.RFC3339)

	q := req.URL.Query()
	q.Add("since", sinceQuery)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		t.Logf("Failed at repository %s\n", href)
		return nil, false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden || resp.StatusCode == http.StatusUnauthorized {
		t.Logf("GitHub API rate limit exceeded or unauthorized (Status: %d). Signal abort.", resp.StatusCode)
		return nil, true
	}

	var respObj []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respObj); err != nil {
		return nil, false
	}

	if len(respObj) == 0 {
		t.Logf("%s has not had a commit in a while", href)
		warnings = append(warnings, href+noRecentCommits)
	}

	return warnings, false
}

func TestStaleRepository(t *testing.T) {
	doc := goqueryFromReadme(t)

	oauth := os.Getenv("OAUTH_TOKEN")
	client := &http.Client{
		Transport: &http.Transport{},
	}

	if oauth == "" {
		t.Log("No oauth token found. Using unauthenticated client ...")
	} else {
		tokenSource := &tokenSource{
			AccessToken: oauth,
		}
		client = oauth2.NewClient(context.Background(), tokenSource)
	}

	addressedRepositories := getAllFlaggedRepositories(t, client)

	var staleRepos []string
	doc.
		Find("body li > a:first-child").
		EachWithBreak(func(_ int, s *goquery.Selection) bool {
			href, ok := s.Attr("href")
			if !ok {
				t.Log("expected to have href")
				return true
			}

			if strings.HasPrefix(href, "#") {
				return true
			}

			u, parseErr := url.Parse(href)
			if parseErr == nil {
				u.RawQuery = ""
				u.Fragment = ""
				href = u.String()
			}

			if ctr >= LIMIT && LIMIT != -1 {
				t.Log("Max number of issues created")
				return false
			}

			if _, issueExists := addressedRepositories[href]; issueExists {
				t.Logf("issue already exists for %s\n", href)
				return true
			}

			if !reGithubRepo.MatchString(href) {
				t.Logf("%s non-github repo not currently handled", href)
				return true
			}

			var aborted bool
			var availWarnings []string
			availWarnings, aborted = checkRepoAvailability(t, true, href, client)
			if aborted {
				t.Logf("Scan aborted due to API limits. Proceeding to create issue with %d records.", len(staleRepos))
				return false
			}
			staleRepos = append(staleRepos, availWarnings...)

			isRepoAdded := len(availWarnings) > 0
			var commitWarnings []string
			commitWarnings, aborted = checkRepoCommitActivity(t, !isRepoAdded, href, client)
			if aborted {
				t.Logf("Scan aborted due to API limits. Proceeding to create issue with %d records.", len(staleRepos))
				return false
			}
			staleRepos = append(staleRepos, commitWarnings...)

			if len(availWarnings)+len(commitWarnings) > 0 {
				ctr++
			}

			return true
		})

	createIssue(t, staleRepos, client)
}
