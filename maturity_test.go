package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	githubApiAuthorizationToken = os.Getenv("GITHUB_API_TOKEN")
	minimumMaturityDate         = time.Now().AddDate(0, -5, 0)
)

func TestMaturity(t *testing.T) {
	doc := goqueryFromReadme(t)
	doc.Find("body li > a:first-child").Each(func(_ int, s *goquery.Selection) {
		t.Run(s.Text(), func(t *testing.T) {
			href, ok := s.Attr("href")
			if !ok {
				t.Error("expected to have href")
			}

			matches := reGithubRepo.FindStringSubmatch(href)
			if matches == nil {
				return
			}

			if len(matches) != 3 {
				t.Fatalf("failed to extract repo and user from: %s, got [%v]", href, strings.Join(matches, ", "))
			}

			if err := checkRepositoryMaturity(matches[1], matches[2]); err != nil {
				t.Fatal(err)
			}
		})
	})
}

func checkRepositoryMaturity(user, repo string) error {
	until := minimumMaturityDate.Format(time.RFC3339)
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits?per_page=1&until=%s", user, repo, until)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request for `%s`, %v", url, err)
	}

	request.Header.Set("Accept", "application/vnd.github+json")
	request.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	request.Header.Set("User-Agent", "avelino")
	request.Header.Set("Authorization", "Bearer "+githubApiAuthorizationToken)

	http.DefaultClient.Timeout = time.Minute
	httpRes, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("failed to fetch commits for [%s/%s], %v", user, repo, err)
	}
	defer httpRes.Body.Close()

	var commits []any
	err = json.NewDecoder(httpRes.Body).Decode(&commits)
	if err != nil {
		return fmt.Errorf("failed to decode response for [%s/%s], %v", user, repo, err)
	}

	if len(commits) == 0 {
		minimumDate := minimumMaturityDate.Format(time.DateOnly)
		return fmt.Errorf("the project [%s/%s] doesn't have any commits before %s, this is a maturity violation", user, repo, minimumDate)
	}

	return nil
}
