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
	"text/template"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/oauth2"
)

const issueTemplate = `
{{range .}}
- [ ] {{.}}
{{end}}
`

var reGithubRepo = regexp.MustCompile("https://github.com/[a-zA-Z0-9-.]+/[a-zA-Z0-9-.]+$")
var githubGETCOMMITS = "https://api.github.com/repos%s/commits"
var githubPOSTISSUES = "https://api.github.com/repos/avelino/awesome-go/issues"
var awesomeGoGETISSUES = "http://api.github.com/repos/avelino/awesome-go/issues" //only returns open issues
var numberOfYears time.Duration = 1

const issueTitle = "Investigate repositories with more than 1 year without update"
const deadLinkMessage = " this repository might no longer exist!"

var delay time.Duration = 1

//LIMIT specifies the max number of repositories that are added in a single run of the script
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
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}
func getRepositoriesFromBody(body string) []string {
	links := strings.Split(body, "- ")
	for idx, link := range links {
		str := strings.ReplaceAll(link, "\r", "")
		str = strings.ReplaceAll(str, deadLinkMessage, "")
		str = strings.ReplaceAll(str, "[ ]", "")
		str = strings.ReplaceAll(str, "[x]", "")
		str = strings.ReplaceAll(str, " ", "")
		str = strings.ReplaceAll(str, "\n", "")
		links[idx] = str
	}
	return links
}
func generateIssueBody(repositories []string) (string, error) {
	var writer bytes.Buffer
	t := template.New("issue")
	temp, err := t.Parse(issueTemplate)
	if err != nil {
		log.Print("Failed to generate template")
		return "", err
	}
	err = temp.Execute(&writer, repositories)
	if err != nil {
		log.Print("Failed to generate template")
		return "", err
	}
	issueBody := writer.String()
	return issueBody, nil
}
func createIssue(staleRepos []string, oauthClient *http.Client) {
	if len(staleRepos) == 0 {
		log.Print("NO STALE REPOSITORIES")
		return
	}
	body, err := generateIssueBody(staleRepos)
	if err != nil {
		log.Print("Failed at CreateIssue")
		return
	}
	newIssue := &issue{
		Title: issueTitle,
		Body:  body,
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(newIssue)
	req, err := http.NewRequest("POST", githubPOSTISSUES, buf)
	if err != nil {
		log.Print("Failed at CreateIssue")
		return
	}
	oauthClient.Do(req)
}
func getAllFlaggedRepositories(oauthClient *http.Client, flaggedRepositories *map[string]bool) error {
	req, err := http.NewRequest("GET", awesomeGoGETISSUES, nil)
	if err != nil {
		log.Print("Failed to get all issues")
		return err
	}
	res, err := oauthClient.Do(req)
	if err != nil {
		log.Print("Failed to get all issues")
		return err
	}
	target := []issue{}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&target)
	for _, i := range target {
		if i.Title == issueTitle {
			repos := getRepositoriesFromBody(i.Body)
			fmt.Println(repos)
			for _, repo := range repos {
				(*flaggedRepositories)[repo] = true
			}
		}
	}
	return nil
}
func containsOpenIssue(link string, openIssues map[string]bool) bool {
	_, ok := openIssues[link]
	if ok {
		return true
	}
	return false
}
func testCommitAge(href string, oauthClient *http.Client, staleRepos *[]string, openIssues map[string]bool) {
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
		issueExists := containsOpenIssue(href, openIssues)
		if issueExists {
			log.Printf("issue already exists for %s\n", href)
			return
		}
		ownerRepo = strings.ReplaceAll(href, "https://github.com", "")
		apiCall = fmt.Sprintf(githubGETCOMMITS, ownerRepo)
		req, err := http.NewRequest("GET", apiCall, nil)
		if err != nil {
			log.Printf("Failed at repository %s\n", href)
			return
		}
		q := req.URL.Query()
		q.Add("since", sinceQuery)
		req.URL.RawQuery = q.Encode()
		resp, err := oauthClient.Do(req)
		if err != nil {
			log.Printf("Failed at repository %s\n", href)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode > 400 {
			*staleRepos = append(*staleRepos, href+" this repository might no longer exist!")
			log.Printf("%s might not exist!", href)
		} else {
			json.NewDecoder(resp.Body).Decode(&respObj)
			isAged = len(respObj) == 0
			if isAged {
				log.Printf("%s has not had a commit in a while", href)
				*staleRepos = append(*staleRepos, href)
				ctr++
			}
		}
	}
}
func testStaleRepository() {
	query := startQuery()
	var staleRepos []string
	addressedRepositories := make(map[string]bool)
	oauth := os.Getenv("GITHUB_OAUTH_TOKEN")
	tokenSource := &tokenSource{
		AccessToken: oauth,
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	err := getAllFlaggedRepositories(oauthClient, &addressedRepositories)

	if err != nil {
		log.Println("Failed to get existing issues. Exiting...")
		return
	}
	query.Find("body li > a:first-child").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		href, ok := s.Attr("href")
		if !ok {
			log.Println("expected to have href")
		} else {
			if ctr >= LIMIT && LIMIT != -1 {
				log.Print("Max number of issues created")
				return false
			}
			testCommitAge(href, oauthClient, &staleRepos, addressedRepositories)
		}
		return true
	})
	createIssue(staleRepos, oauthClient)
}

func main() {
	f, err := os.OpenFile("test_stale_repositories_log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Println("FAILED TO INIT LOG FILE")
		return
	}
	log.SetOutput(f)
	testStaleRepository()
}
