// auto-merge-readme squash merges open pull requests that only add packages
// to README.md, once a maintainer approved the head commit, every check is
// green and each added repository is at least five months old.
//
// It never checks out or runs pull request code. Everything it knows comes
// from the GitHub API. Any doubt (unexpected diff, API error, missing data)
// means the pull request is skipped and left for a human.
//
// Environment:
//
//	GITHUB_TOKEN       token with contents and pull-requests write (read is enough for dry runs)
//	GITHUB_REPOSITORY  owner/repo
//	DRY_RUN            "false" to merge; anything else only reports
//	PR_NUMBERS         optional space separated list; defaults to all open PRs
//	HEAD_SHA           optional; only the open PR with this head is evaluated
//	REPLAY             "true" (dry run only) evaluates already merged PRs as of
//	                   their merge time, to check the rules against history
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// "Post quality report" carries the blocking quality and diff checks from
// PR Quality Checks, so it must have run and passed, not just be absent.
var requiredChecks = []string{"Running test", "Post quality report"}

const selfCheck = "Auto-merge README PRs"

type client struct {
	token, repo string
	http        *http.Client
}

func (c *client) do(method, path string, body, out any) error {
	var rd io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		rd = bytes.NewReader(b)
	}
	u := path
	if !strings.HasPrefix(u, "https://") {
		u = "https://api.github.com" + path
	}
	req, err := http.NewRequest(method, u, rd)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("%s %s: %d %s", method, path, resp.StatusCode, strings.TrimSpace(string(data)))
	}
	if out != nil {
		return json.Unmarshal(data, out)
	}
	return nil
}

type pull struct {
	Number       int                  `json:"number"`
	State        string               `json:"state"`
	Draft        bool                 `json:"draft"`
	Mergeable    *bool                `json:"mergeable"`
	ChangedFiles int                  `json:"changed_files"`
	MergedAt     *time.Time           `json:"merged_at"`
	Base         struct{ Ref string } `json:"base"`
	Head         struct{ SHA string } `json:"head"`
}

func (c *client) pulls(head string) ([]int, error) {
	var nums []int
	for page := 1; page <= 10; page++ {
		var ps []pull
		if err := c.do("GET", fmt.Sprintf("/repos/%s/pulls?state=open&base=main&per_page=100&page=%d", c.repo, page), nil, &ps); err != nil {
			return nil, err
		}
		for _, p := range ps {
			if head != "" && p.Head.SHA != head {
				continue
			}
			nums = append(nums, p.Number)
		}
		if len(ps) < 100 {
			return nums, nil
		}
	}
	return nil, fmt.Errorf("too many open pull requests")
}

func (c *client) maintainer(cache map[string]bool) func(string) (bool, error) {
	return func(user string) (bool, error) {
		if v, ok := cache[user]; ok {
			return v, nil
		}
		var p struct {
			Permission string `json:"permission"`
		}
		if err := c.do("GET", fmt.Sprintf("/repos/%s/collaborators/%s/permission", c.repo, url.PathEscape(user)), nil, &p); err != nil {
			return false, err
		}
		ok := p.Permission == "admin" || p.Permission == "maintain" || p.Permission == "write"
		cache[user] = ok
		return ok, nil
	}
}

// evaluate returns the head SHA to merge, or an error explaining why not.
func (c *client) evaluate(n int, now time.Time, perms map[string]bool, replay bool) (string, error) {
	var p pull
	// GitHub computes mergeability in the background, so the first read
	// often has it unset.
	for try := 0; try < 4; try++ {
		if try > 0 {
			time.Sleep(3 * time.Second)
		}
		if err := c.do("GET", fmt.Sprintf("/repos/%s/pulls/%d", c.repo, n), nil, &p); err != nil {
			return "", err
		}
		if p.Mergeable != nil || replay {
			break
		}
	}
	if replay {
		if p.MergedAt == nil {
			return "", fmt.Errorf("replay needs a merged PR")
		}
		now = *p.MergedAt
	} else {
		if p.State != "open" || p.Draft || p.Base.Ref != "main" {
			return "", fmt.Errorf("not an open, ready PR against main")
		}
		if p.Mergeable == nil || !*p.Mergeable {
			return "", fmt.Errorf("not mergeable")
		}
	}
	head := p.Head.SHA

	var files []struct {
		Filename string `json:"filename"`
		Status   string `json:"status"`
		Patch    string `json:"patch"`
	}
	if err := c.do("GET", fmt.Sprintf("/repos/%s/pulls/%d/files?per_page=100", c.repo, n), nil, &files); err != nil {
		return "", err
	}
	names := make([]string, len(files))
	for i, f := range files {
		names[i] = f.Filename
	}
	if p.ChangedFiles != 1 || !onlyReadme(names) || files[0].Status != "modified" {
		return "", fmt.Errorf("changes files other than README.md: %v", names)
	}
	repos, err := addedRepos(files[0].Patch)
	if err != nil {
		return "", err
	}
	for _, r := range repos {
		var meta struct {
			FullName  string    `json:"full_name"`
			CreatedAt time.Time `json:"created_at"`
			Archived  bool      `json:"archived"`
		}
		if err := c.do("GET", "/repos/"+r.Owner+"/"+r.Name, nil, &meta); err != nil {
			return "", fmt.Errorf("repository %s: %w", r, err)
		}
		if !strings.EqualFold(meta.FullName, r.String()) {
			return "", fmt.Errorf("repository %s redirects to %s", r, meta.FullName)
		}
		if meta.Archived {
			return "", fmt.Errorf("repository %s is archived", r)
		}
		if !oldEnough(meta.CreatedAt, now) {
			return "", fmt.Errorf("repository %s was created %s, less than %d months ago", r, meta.CreatedAt.Format("2006-01-02"), minRepoAgeMonths)
		}
	}

	var raw []struct {
		User     struct{ Login string } `json:"user"`
		State    string                 `json:"state"`
		CommitID string                 `json:"commit_id"`
	}
	if err := c.do("GET", fmt.Sprintf("/repos/%s/pulls/%d/reviews?per_page=100", c.repo, n), nil, &raw); err != nil {
		return "", err
	}
	if len(raw) == 100 {
		return "", fmt.Errorf("too many reviews to read safely")
	}
	reviews := make([]review, len(raw))
	for i, r := range raw {
		reviews[i] = review{r.User.Login, r.State, r.CommitID}
	}
	approvers, err := approvalStatus(reviews, head, c.maintainer(perms))
	if err != nil {
		return "", err
	}
	if len(approvers) == 0 {
		return "", fmt.Errorf("no maintainer approval on head %s", head[:7])
	}

	var cr struct {
		Total int `json:"total_count"`
		Runs  []struct {
			Name       string `json:"name"`
			Status     string `json:"status"`
			Conclusion string `json:"conclusion"`
		} `json:"check_runs"`
	}
	if err := c.do("GET", fmt.Sprintf("/repos/%s/commits/%s/check-runs?per_page=100", c.repo, head), nil, &cr); err != nil {
		return "", err
	}
	if cr.Total > len(cr.Runs) {
		return "", fmt.Errorf("too many check runs to read safely")
	}
	runs := make([]checkRun, len(cr.Runs))
	for i, r := range cr.Runs {
		runs[i] = checkRun{r.Name, r.Status, r.Conclusion}
	}
	if err := checksGreen(runs, requiredChecks, selfCheck); err != nil {
		return "", err
	}
	var st struct {
		State string `json:"state"`
		Total int    `json:"total_count"`
	}
	if err := c.do("GET", fmt.Sprintf("/repos/%s/commits/%s/status", c.repo, head), nil, &st); err != nil {
		return "", err
	}
	if st.Total > 0 && st.State != "success" {
		return "", fmt.Errorf("commit status is %s", st.State)
	}
	fmt.Printf("#%d: approved by %s, adds %v\n", n, strings.Join(approvers, ", "), repos)
	return head, nil
}

func main() {
	c := &client{token: os.Getenv("GITHUB_TOKEN"), repo: os.Getenv("GITHUB_REPOSITORY"), http: &http.Client{Timeout: 30 * time.Second}}
	if c.repo == "" {
		fmt.Fprintln(os.Stderr, "GITHUB_REPOSITORY is not set")
		os.Exit(1)
	}
	dry := os.Getenv("DRY_RUN") != "false"
	replay := os.Getenv("REPLAY") == "true"
	if replay && !dry {
		fmt.Fprintln(os.Stderr, "REPLAY only works as a dry run")
		os.Exit(1)
	}

	var nums []int
	if s := strings.Fields(os.Getenv("PR_NUMBERS")); len(s) > 0 {
		for _, f := range s {
			n, err := strconv.Atoi(strings.TrimPrefix(f, "#"))
			if err != nil {
				fmt.Fprintf(os.Stderr, "bad PR number %q\n", f)
				os.Exit(1)
			}
			nums = append(nums, n)
		}
	} else {
		var err error
		if nums, err = c.pulls(os.Getenv("HEAD_SHA")); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	perms := map[string]bool{}
	merged, failed := 0, false
	for _, n := range nums {
		head, err := c.evaluate(n, time.Now().UTC(), perms, replay)
		if err != nil {
			fmt.Printf("#%d: skip, %v\n", n, err)
			continue
		}
		if dry {
			fmt.Printf("#%d: would squash merge %s (dry run)\n", n, head[:7])
			continue
		}
		// Passing sha makes GitHub refuse the merge if a new commit landed
		// after the checks above.
		if err := c.do("PUT", fmt.Sprintf("/repos/%s/pulls/%d/merge", c.repo, n), map[string]string{"sha": head, "merge_method": "squash"}, nil); err != nil {
			fmt.Printf("#%d: merge failed, %v\n", n, err)
			failed = true
			continue
		}
		fmt.Printf("#%d: squash merged %s\n", n, head[:7])
		merged++
	}
	if out := os.Getenv("GITHUB_OUTPUT"); out != "" {
		if f, err := os.OpenFile(out, os.O_APPEND|os.O_WRONLY, 0o644); err == nil {
			fmt.Fprintf(f, "merged=%d\n", merged)
			f.Close()
		}
	}
	if failed {
		os.Exit(1)
	}
}
