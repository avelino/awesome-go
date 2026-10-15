package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// MinRepoAge is how old a listed repository must be before its PR can merge
// without a human pressing the button.
const minRepoAgeMonths = 5

var (
	reEntry    = regexp.MustCompile(`^[-*] \[[^\]]+\]\((https://github\.com/[^)\s]+)\) - \S.*$`)
	reRepoPath = regexp.MustCompile(`^https://github\.com/([A-Za-z0-9-]+)/([A-Za-z0-9._-]+?)/?$`)
)

type repoRef struct{ Owner, Name string }

func (r repoRef) String() string { return r.Owner + "/" + r.Name }

// onlyReadme reports whether the changed file list is exactly README.md.
func onlyReadme(files []string) bool {
	return len(files) == 1 && files[0] == "README.md"
}

// addedRepos reads a unified diff patch and returns the GitHub repositories
// added by it. Anything other than pure additions of well formed entries is
// an error, so the caller does nothing.
func addedRepos(patch string) ([]repoRef, error) {
	var repos []repoRef
	seen := map[string]bool{}
	for _, line := range strings.Split(patch, "\n") {
		switch {
		case line == "", strings.HasPrefix(line, "@@"), strings.HasPrefix(line, " "),
			line == `\ No newline at end of file`:
			continue
		case strings.HasPrefix(line, "-"):
			return nil, fmt.Errorf("patch removes a line: %q", line)
		case strings.HasPrefix(line, "+"):
			entry := strings.TrimPrefix(line, "+")
			if strings.TrimSpace(entry) == "" {
				return nil, fmt.Errorf("patch adds a blank line")
			}
			m := reEntry.FindStringSubmatch(entry)
			if m == nil {
				return nil, fmt.Errorf("added line is not a GitHub package entry: %q", entry)
			}
			p := reRepoPath.FindStringSubmatch(m[1])
			if p == nil {
				return nil, fmt.Errorf("link is not a repository root: %q", m[1])
			}
			ref := repoRef{p[1], strings.TrimSuffix(p[2], ".git")}
			key := strings.ToLower(ref.String())
			if seen[key] {
				return nil, fmt.Errorf("repository %s added twice", ref)
			}
			seen[key] = true
			repos = append(repos, ref)
		default:
			return nil, fmt.Errorf("unexpected patch line: %q", line)
		}
	}
	if len(repos) == 0 {
		return nil, fmt.Errorf("patch adds no package entry")
	}
	return repos, nil
}

// oldEnough reports whether a repository created at createdAt is at least
// minRepoAgeMonths calendar months old at now.
func oldEnough(createdAt, now time.Time) bool {
	if createdAt.IsZero() || createdAt.After(now) {
		return false
	}
	return !now.Before(addMonthsClamped(createdAt, minRepoAgeMonths))
}

// addMonthsClamped adds calendar months to t, clamping to the last day of the
// target month instead of rolling over (Sep 30 + 5 months is Feb 28, not Mar 2).
func addMonthsClamped(t time.Time, months int) time.Time {
	y, m, d := t.Date()
	first := time.Date(y, m+time.Month(months), 1, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	last := first.AddDate(0, 1, -1).Day()
	if d > last {
		d = last
	}
	return first.AddDate(0, 0, d-1)
}

type review struct {
	User     string
	State    string
	CommitID string
}

// approvalStatus returns the maintainers whose latest review approves the
// head commit, and an error if any maintainer's latest review requests
// changes. isMaintainer must be true only for write, maintain or admin.
func approvalStatus(reviews []review, head string, isMaintainer func(string) (bool, error)) ([]string, error) {
	latest := map[string]review{}
	var order []string
	for _, r := range reviews {
		if r.State == "COMMENTED" || r.State == "PENDING" {
			continue
		}
		if _, ok := latest[r.User]; !ok {
			order = append(order, r.User)
		}
		latest[r.User] = r
	}
	var approvers []string
	for _, u := range order {
		r := latest[u]
		ok, err := isMaintainer(u)
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}
		switch r.State {
		case "CHANGES_REQUESTED":
			return nil, fmt.Errorf("%s requested changes", u)
		case "APPROVED":
			if r.CommitID == head {
				approvers = append(approvers, u)
			}
		}
	}
	return approvers, nil
}

type checkRun struct {
	Name       string
	Status     string
	Conclusion string
}

// checksGreen requires every required check to be present and successful and
// every other check run on the head commit to be finished without failing.
func checksGreen(runs []checkRun, required []string, self string) error {
	got := map[string]bool{}
	for _, c := range runs {
		if c.Name == self {
			continue
		}
		if c.Status != "completed" {
			return fmt.Errorf("check %q is %s", c.Name, c.Status)
		}
		switch c.Conclusion {
		case "success":
			got[c.Name] = true
		case "skipped", "neutral":
		default:
			return fmt.Errorf("check %q concluded %s", c.Name, c.Conclusion)
		}
	}
	for _, r := range required {
		if !got[r] {
			return fmt.Errorf("required check %q has not passed", r)
		}
	}
	return nil
}
