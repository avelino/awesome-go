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
var awesomeGoGETISSUES = "http://api.github.com/repos/avelino/awesome-go/issues" //only returns open issues
var numberOfYears time.Duration = 1
var issueFmt = "Investigate %s. This repository has not had any new commits in a while."
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
func getAllOpenIssues(oauthClient *http.Client) (map[string]bool, error) {
	openIssues := make(map[string]bool)
	req, err := http.NewRequest("GET", awesomeGoGETISSUES, nil)
	if err != nil {
		log.Print("Failed to get all issues")
		return nil, err
	}
	res, err := oauthClient.Do(req)
	if err != nil {
		log.Print("Failed to get all issues")
		return nil, err
	}
	target := []issue{}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&target)
	for _, i := range target {
		openIssues[i.Title] = true
	}
	return openIssues, nil
}
func containsOpenIssue(ownerRepo string, openIssues map[string]bool) bool {
	issueTitle := fmt.Sprintf(issueFmt, ownerRepo)
	_, ok := openIssues[issueTitle]
	if ok {
		return true
	}
	return false
}
func testCommitAge(href string, oauthClient *http.Client, issueRequests *[]*http.Request, openIssues map[string]bool) {
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
		issueExists := containsOpenIssue(ownerRepo, openIssues)
		if issueExists {
			log.Printf("issue already exists for %s\n", ownerRepo)
			return
		}
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
	openIssues, err := getAllOpenIssues(oauthClient)
	if err != nil {
		log.Println("Failed to get existing issues. Exiting...")
		return
	}
	query.Find("body li > a:first-child").Each(func(_ int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			log.Println("expected to have href")
		} else {
			testCommitAge(href, oauthClient, &issueRequests, openIssues)
		}
	})
	createIssues(issueRequests, oauthClient)
}

func main() {
	testStaleRepository()
}
