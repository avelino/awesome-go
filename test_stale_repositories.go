package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/oauth2"
)

var reGithubRepo = regexp.MustCompile("https://github.com/[a-zA-Z0-9-.]+/[a-zA-Z0-9-.]+$")
var githubGETCOMMITS = "https://api.github.com/repos%s/commits"
var githubPOSTISSUES = "https://api.github.com/repos/avelino/awesome-go/issues"
var numberOfYears time.Duration = 1
var issueFmt = "Investigate %s. Has not had any new commits in a while"
var delay time.Duration = 1

type tokenSource struct {
	AccessToken string
}
type issue struct {
	Title string `json:"title"`
}

func (t *tokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}
func createIssues(issueRequests []*http.Request, oauthClient *http.Client) {
	for _, req := range issueRequests {
		oauthClient.Do(req)
		time.Sleep(delay * time.Second)
	}
}
func testCommitAge(href string, oauthClient *http.Client, issueRequests *[]*http.Request) {
	var isValidRepo bool
	var isAged bool
	var ownerRepo string
	isValidRepo = reGithubRepo.MatchString(href)
	var apiCall string
	var respObj []map[string]interface{}
	now := time.Now()
	since := now.Add(-1 * 365 * 24 * numberOfYears * time.Hour)
	sinceQuery := since.Format(time.RFC3339)
	if isValidRepo {
		ownerRepo = strings.ReplaceAll(href, "https://github.com", "")
		apiCall = fmt.Sprintf(githubGETCOMMITS, ownerRepo)
		req, err := http.NewRequest("GET", apiCall, nil)
		if err != nil {
			log.Printf("Failed at repository %s\n", ownerRepo)
			return
		}
		q := req.URL.Query()
		q.Add("since", sinceQuery)
		req.URL.RawQuery = q.Encode()
		resp, err := oauthClient.Do(req)
		if err != nil {
			log.Printf("Failed at repository %s\n", ownerRepo)
			return
		}
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&respObj)
		isAged = len(respObj) == 0
		if isAged {
			body := &issue{
				Title: fmt.Sprintf(issueFmt, ownerRepo),
			}
			buf := new(bytes.Buffer)
			json.NewEncoder(buf).Encode(body)
			req, err := http.NewRequest("POST", githubPOSTISSUES, buf)
			if err != nil {
				log.Printf("Failed at repository %s\n", ownerRepo)
				return
			}
			*issueRequests = append(*issueRequests, req)
		}
	}
}
func testStaleRepository() {
	query := startQuery()
	var issueRequests []*http.Request
	oauth := os.Getenv("GITHUB_OAUTH_TOKEN")
	tokenSource := &tokenSource{
		AccessToken: oauth,
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	query.Find("body li > a:first-child").Each(func(_ int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			log.Println("expected to have href")
		}
		testCommitAge(href, oauthClient, &issueRequests)
	})
	createIssues(issueRequests, oauthClient)
}

func main() {
	testStaleRepository()
}
