package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

// FIXME: use official github client
var reGithubRepo = regexp.MustCompile("https://github.com/[a-zA-Z0-9-._]+/[a-zA-Z0-9-._]+$")
var githubGETREPO = "https://api.github.com/repos%s"
var githubGETCOMMITS = "https://api.github.com/repos%s/commits"
var githubPOSTISSUES = "https://api.github.com/repos/avelino/awesome-go/issues"

// FIXME: use https
var awesomeGoGETISSUES = "http://api.github.com/repos/avelino/awesome-go/issues" //only returns open issues
// FIXME: variable has type Duration, but contains a number. we should use
//
//	time.Hour * ... or change type of variable
var numberOfYears time.Duration = 1
var timeNow = time.Now()
var issueTitle = fmt.Sprintf("Investigate repositories with more than 1 year without update - %s", timeNow.Format("2006-01-02"))

const deadLinkMessage = " this repository might no longer exist! (status code >= 400 returned)"
const movedPermanently = " status code 301 received"
const status302 = " status code 302 received"
const archived = " repository has been archived"

// LIMIT specifies the max number of repositories that are added in a single run of the script
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
		log.Print("NO STALE REPOSITORIES")
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

	_, roundTripErr := client.Do(req)
	requireNoErr(t, roundTripErr, "failed to send request")
}

func getAllFlaggedRepositories(t *testing.T, client *http.Client) map[string]bool {
	t.Helper()

	req, err := http.NewRequest(http.MethodGet, awesomeGoGETISSUES, nil)
	requireNoErr(t, err, "failed to create request")

	res, err := client.Do(req)
	requireNoErr(t, err, "failed to send request")

	defer res.Body.Close()

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

func checkRepoAvailability(toRun bool, href string, client *http.Client) ([]string, bool) {
	if !toRun {
		return nil, false
	}

	ownerRepo := strings.ReplaceAll(href, "https://github.com", "")
	apiCall := fmt.Sprintf(githubGETREPO, ownerRepo)
	req, err := http.NewRequest(http.MethodGet, apiCall, nil)
	if err != nil {
		log.Printf("Failed at repository %s\n", href)
		return nil, false
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed at repository %s\n", href)
		return nil, false
	}

	defer resp.Body.Close()

	var repoResp struct {
		Archived bool `json:"archived"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&repoResp); err != nil {
		return nil, false
	}

	var isRepoAdded bool

	var warnings []string
	if resp.StatusCode == http.StatusMovedPermanently {
		warnings = append(warnings, href+movedPermanently)
		log.Printf("%s returned %d", href, resp.StatusCode)
		isRepoAdded = true
	}

	if resp.StatusCode == http.StatusFound && !isRepoAdded {
		warnings = append(warnings, href+status302)
		log.Printf("%s returned %d", href, resp.StatusCode)
		isRepoAdded = true
	}

	if resp.StatusCode >= http.StatusBadRequest && !isRepoAdded {
		warnings = append(warnings, href+deadLinkMessage)
		log.Printf("%s might not exist!", href)
		isRepoAdded = true
	}

	if repoResp.Archived && !isRepoAdded {
		warnings = append(warnings, href+archived)
		log.Printf("%s is archived!", href)
		isRepoAdded = true
	}

	// FIXME: expression `(len(warnings) > 0) == isRepoAdded` is always true.
	return warnings, isRepoAdded
}

func checkRepoCommitActivity(toRun bool, href string, client *http.Client) ([]string, bool) {
	if !toRun {
		return nil, false
	}

	ownerRepo := strings.ReplaceAll(href, "https://github.com", "")
	apiCall := fmt.Sprintf(githubGETCOMMITS, ownerRepo)
	req, err := http.NewRequest(http.MethodGet, apiCall, nil)
	if err != nil {
		log.Printf("Failed at repository %s\n", href)
		return nil, false
	}

	since := timeNow.Add(-1 * 365 * 24 * numberOfYears * time.Hour)
	sinceQuery := since.Format(time.RFC3339)

	q := req.URL.Query()
	q.Add("since", sinceQuery)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed at repository %s\n", href)
		return nil, false
	}

	defer resp.Body.Close()

	var respObj []map[string]interface{}
	// FIXME: handle error in all that cases
	if err := json.NewDecoder(resp.Body).Decode(&respObj); err != nil {
		return nil, false
	}

	var warnings []string
	var isRepoAdded bool
	isAged := len(respObj) == 0
	if isAged {
		log.Printf("%s has not had a commit in a while", href)
		warnings = append(warnings, href)
		isRepoAdded = true
	}

	// FIXME: expression `(len(warnings) > 0) == isRepoAdded` is always true.
	return warnings, isRepoAdded
}

func TestStaleRepository(t *testing.T) {
	doc := goqueryFromReadme(t)

	oauth := os.Getenv("OAUTH_TOKEN")
	client := &http.Client{
		Transport: &http.Transport{},
	}

	if oauth == "" {
		log.Print("No oauth token found. Using unauthenticated client ...")
	} else {
		tokenSource := &tokenSource{
			AccessToken: oauth,
		}
		client = oauth2.NewClient(context.Background(), tokenSource)
	}

	// FIXME: return addressedRepositories, no need to pass
	addressedRepositories := getAllFlaggedRepositories(t, client)

	var staleRepos []string
	doc.
		Find("body li > a:first-child").
		EachWithBreak(func(_ int, s *goquery.Selection) bool {
			href, ok := s.Attr("href")
			if !ok {
				log.Println("expected to have href")
				return true
			}

			if ctr >= LIMIT && LIMIT != -1 {
				log.Print("Max number of issues created")
				return false
			}

			if _, issueExists := addressedRepositories[href]; issueExists {
				log.Printf("issue already exists for %s\n", href)
				return true
			}

			if !reGithubRepo.MatchString(href) {
				log.Printf("%s non-github repo not currently handled", href)
			}

			// FIXME: this is `or` expres24sion. Probably we need `and`?
			warnings, isRepoAdded := checkRepoAvailability(true, href, client)
			staleRepos = append(staleRepos, warnings...)

			warnings, isRepoAdded = checkRepoCommitActivity(!isRepoAdded, href, client)
			staleRepos = append(staleRepos, warnings...)

			if isRepoAdded {
				ctr++
			}

			return true
		})

	createIssue(t, staleRepos, client)
}
