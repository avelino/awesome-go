package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/kr/pretty"
)

const gqlQuery = `
query {
  repository(owner:"%s", name:"%s") {
    name
    openIssues: issues(states: OPEN) {
      totalCount
    }
    lastUpdatedIssue: issues(first: 1, orderBy: {field: UPDATED_AT, direction: DESC}) {
      edges {
        node {
          updatedAt
        }
      }
    }
    defaultBranchRef {
      target {
        ... on Commit {
          committedDate
        }
      }
    }
    forks {
      totalCount
    }
    stargazers {
      totalCount
    }
  }
}
`

var authToken = os.ExpandEnv("$GITHUB_PAT")

func repoTimeSinceLastUpdated(repo GithubRepoResponse) time.Duration {
	return time.Since(repo.Data.Repository.DefaultBranchRef.Target.CommittedDate)
}

type awesomeRepo struct {
	RepoName         string
	RepoOrganization string
	ReadmeLine       int
}

type GithubRepoResponse struct {
	Data struct {
		Repository struct {
			Name       string `json:"name"`
			OpenIssues struct {
				TotalCount int `json:"totalCount"`
			} `json:"openIssues"`
			LastUpdatedIssue struct {
				Edges []struct {
					Node struct {
						UpdatedAt time.Time `json:"updatedAt"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"lastUpdatedIssue"`
			DefaultBranchRef struct {
				Target struct {
					CommittedDate time.Time `json:"committedDate"`
				} `json:"target"`
			} `json:"defaultBranchRef"`
			Forks struct {
				TotalCount int `json:"totalCount"`
			} `json:"forks"`
			Stargazers struct {
				TotalCount int `json:"totalCount"`
			} `json:"stargazers"`
		} `json:"repository"`
	} `json:"data"`
}

type repoStats struct {
	TimeLastUpdated time.Duration
	Stars           int
	Forks           int
}

func checkHealth(repo awesomeRepo) (*repoStats, error) {
	requestBody := struct {
		Query string `json:"query"`
	}{
		fmt.Sprintf(gqlQuery, repo.RepoOrganization, repo.RepoName),
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", "https://api.github.com/graphql", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+authToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resbody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response GithubRepoResponse

	if err := json.Unmarshal(resbody, &response); err != nil {
		return nil, err
	}
	timeSinceUpdated := repoTimeSinceLastUpdated(response)
	health := repoStats{
		TimeLastUpdated: timeSinceUpdated,
	}
	return &health, nil
}

func TestCleanUpLibs(t *testing.T) {
	bytes, err := os.ReadFile("README.md")
	if err != nil {
		t.Error(err)
	}

	r := regexp.MustCompile(`\((https.*)\)`)

	// only support github right now
	for _, result := range r.FindAllSubmatch(bytes, -1) {
		org, p, err := extractGithubRepo(result[1])
		if err != nil {
			continue
		}
		fmt.Println("org", org, "package", p, "err", err)

		a := awesomeRepo{
			RepoOrganization: org,
			RepoName:         p,
		}

		health, err := checkHealth(a)
		if err != nil {
			t.Error(err)
		}
		pretty.Println(health)
	}
}

var getRepo = regexp.MustCompile(`https://github.com/(.*)/(.*)[\)]?`)

func extractGithubRepo(b []byte) (string, string, error) {
	results := getRepo.FindSubmatch(b)
	if len(results) == 0 {
		return "", "", fmt.Errorf("no match found for %s", string(b))
	}

	return string(results[1]), string(results[2]), nil
}
