package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/oauth2"
)

var reGithubRepo = regexp.MustCompile("https://github.com/[a-zA-Z0-9-.]+/[a-zA-Z0-9-.]+$")
var githubGETCOMMITS = "https://api.github.com/repos%s/commits"
var githubPOSTISSUES = "https://api.github.com/repos/avelino/awesome-go/issues"
var NUMBEROFYEARS time.Duration = 1
var ISSUEFMT = "Investigate %s. Has not had any new commits in a while"
var DELAYINSECONDS time.Duration = 1

type TokenSource struct {
	AccessToken string
}
type Issue struct {
	Title string `json:"title"`
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func createIssues(issueRequests []*http.Request, oauthClient *http.Client, t *testing.T) {
	for _, req := range issueRequests {
		oauthClient.Do(req)
		time.Sleep(DELAYINSECONDS * time.Second)
	}
}
func testCommitAge(href string, oauthClient *http.Client, issueRequests *[]*http.Request, t *testing.T) {
	var isValidRepo bool
	var isAged bool
	var ownerRepo string
	isValidRepo = reGithubRepo.MatchString(href)
	var apiCall string
	var respObj []map[string]interface{}
	now := time.Now()
	since := now.Add(-1 * 365 * 24 * NUMBEROFYEARS * time.Hour)
	sinceQuery := since.Format(time.RFC3339)
	if isValidRepo {
		ownerRepo = strings.ReplaceAll(href, "https://github.com", "")
		apiCall = fmt.Sprintf(githubGETCOMMITS, ownerRepo)
		req, err := http.NewRequest("GET", apiCall, nil)
		if err != nil {
			t.Errorf("Failed at repository %s", ownerRepo)
			return
		}
		q := req.URL.Query()
		q.Add("since", sinceQuery)
		req.URL.RawQuery = q.Encode()
		resp, err := oauthClient.Do(req)
		if err != nil {
			t.Errorf("Failed at repository %s", ownerRepo)
			return
		}
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&respObj)
		isAged = len(respObj) == 0
		if isAged {
			body := &Issue{
				Title: fmt.Sprintf(ISSUEFMT, ownerRepo),
			}
			buf := new(bytes.Buffer)
			json.NewEncoder(buf).Encode(body)
			req, err := http.NewRequest("POST", githubPOSTISSUES, buf)
			if err != nil {
				t.Errorf("Failed at repository %s", ownerRepo)
				return
			}
			*issueRequests = append(*issueRequests, req)
		}
	}
}
func TestStaleRepository(t *testing.T) {
	query := startQuery()
	var issueRequests []*http.Request
	oauth := os.Getenv("GITHUB_OAUTH_TOKEN")
	tokenSource := &TokenSource{
		AccessToken: oauth,
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	query.Find("body li > a:first-child").Each(func(_ int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			t.Error("expected to have href")
		}
		testCommitAge(href, oauthClient, &issueRequests, t)
	})
	createIssues(issueRequests, oauthClient, t)
}
