// check-pr-diff validates the actual changes in a PR against CONTRIBUTING.md rules:
//   - PR modifies only README.md (for package additions)
//   - PR adds or removes exactly one item
//   - Added link matches the forge link declared in PR body
//   - Link text matches the repository/project name
//   - Description ends with a period and is non-promotional
//   - Category has minimum 3 items after the change
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	reForgeLink = regexp.MustCompile(`(?i)forge\s+link[^:]*:\s*(https?://(?:github\.com|gitlab\.com|bitbucket\.org)/\S+)`)
	reEntry     = regexp.MustCompile(`^- \[([^\]]+)\]\(([^)]+)\)\s+-\s+(.+)$`)
	reHeading   = regexp.MustCompile(`^#{2,3}\s+(.+)`)
)

// Words that indicate promotional language in descriptions.
var promotionalWords = []string{
	"best", "fastest", "ultimate", "world-class", "blazing",
	"revolutionary", "cutting-edge", "enterprise-grade", "next-generation",
	"state-of-the-art", "unparalleled", "unmatched", "superior",
	"number one", "#1", "award-winning", "top-rated",
}

type prEvent struct {
	PullRequest struct {
		Body string `json:"body"`
	} `json:"pull_request"`
}

type entry struct {
	name        string
	url         string
	description string
	raw         string
}

func main() {
	event := readEvent()
	body := event.PullRequest.Body
	forgeLink := captureMatch(body, reForgeLink)

	var (
		results  []string
		warnings []string
		hasFail  bool
	)

	// 1. Check which files were changed
	changedFiles := getChangedFiles()
	includesReadme := false
	for _, f := range changedFiles {
		if f == "README.md" {
			includesReadme = true
			break
		}
	}

	if !includesReadme {
		results = append(results, icon(false)+" **README change**: no changes to README.md detected")
		hasFail = true
		outputResults(results, warnings, hasFail)
		return
	}

	readmeOnly := len(changedFiles) == 1 && changedFiles[0] == "README.md"
	if readmeOnly {
		results = append(results, icon(true)+" **Files changed**: only README.md")
	} else {
		var others []string
		for _, f := range changedFiles {
			if f != "README.md" {
				others = append(others, f)
			}
		}
		warnings = append(warnings, fmt.Sprintf("%s **Extra files changed**: %s (expected only README.md for package additions)", warnIcon(false), strings.Join(others, ", ")))
	}

	// 2. Parse the diff
	diff := getDiff()
	if diff == "" {
		results = append(results, icon(false)+" **Diff**: could not read diff for README.md")
		hasFail = true
		outputResults(results, warnings, hasFail)
		return
	}

	added, removed := parseDiffEntries(diff)
	totalChanges := len(added) + len(removed)

	// 3. Single item check
	switch {
	case len(added) == 1 && len(removed) == 0:
		results = append(results, icon(true)+" **Single item**: one package added")
	case len(removed) == 1 && len(added) == 0:
		results = append(results, icon(true)+" **Single item**: one package removed")
	case len(added) == 1 && len(removed) == 1:
		warnings = append(warnings, warnIcon(false)+" **Changes**: 1 added + 1 removed (update or move — please confirm in PR description)")
	case totalChanges == 0:
		warnings = append(warnings, warnIcon(false)+" **Entries**: no package entries detected in diff (might be a category or formatting change)")
	default:
		results = append(results, fmt.Sprintf("%s **Single item**: %d added, %d removed (expected exactly 1 change per PR)", icon(false), len(added), len(removed)))
		hasFail = true
	}

	// 4. Validate added entries
	for _, e := range added {
		// 4a. Link matches forge link in PR body
		if forgeLink != "" {
			if normalizeURL(e.url) == normalizeURL(forgeLink) {
				results = append(results, icon(true)+" **Link consistency**: README link matches forge link in PR body")
			} else {
				results = append(results, fmt.Sprintf("%s **Link consistency**: README link `%s` does not match forge link `%s`", icon(false), e.url, forgeLink))
				hasFail = true
			}
		}

		// 4b. Link text matches repo name
		repoName := extractRepoName(e.url)
		if repoName != "" {
			if strings.EqualFold(e.name, repoName) {
				results = append(results, icon(true)+" **Link text**: matches repository name")
			} else {
				warnings = append(warnings, fmt.Sprintf("%s **Link text**: `%s` differs from repo name `%s` (verify this is intentional)", warnIcon(false), e.name, repoName))
			}
		}

		// 4c. Description ends with period
		if strings.HasSuffix(e.description, ".") || strings.HasSuffix(e.description, "!") {
			results = append(results, icon(true)+" **Description**: ends with punctuation")
		} else {
			results = append(results, icon(false)+" **Description**: must end with a period")
			hasFail = true
		}

		// 4d. Non-promotional check
		descLower := strings.ToLower(e.description)
		var promoFound []string
		for _, w := range promotionalWords {
			if strings.Contains(descLower, w) {
				promoFound = append(promoFound, fmt.Sprintf("%q", w))
			}
		}
		if len(promoFound) > 0 {
			warnings = append(warnings, fmt.Sprintf("%s **Promotional language**: description contains: %s", warnIcon(false), strings.Join(promoFound, ", ")))
		} else {
			results = append(results, icon(true)+" **Description tone**: no promotional language detected")
		}

		// 4e. Category minimum items
		cat, count := getCategoryItemCount("README.md", e.url)
		if count >= 3 {
			results = append(results, fmt.Sprintf("%s **Category size**: %s has %d items", icon(true), cat, count))
		} else if count > 0 {
			results = append(results, fmt.Sprintf("%s **Category size**: %s has only %d item(s) (minimum 3 required)", icon(false), cat, count))
			hasFail = true
		}
	}

	outputResults(results, warnings, hasFail)
}

// --- Event reading ---

func readEvent() prEvent {
	var ev prEvent
	path := os.Getenv("GITHUB_EVENT_PATH")
	if path == "" {
		return ev
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return ev
	}
	_ = json.Unmarshal(data, &ev)
	return ev
}

func captureMatch(s string, re *regexp.Regexp) string {
	m := re.FindStringSubmatch(s)
	if len(m) < 2 {
		return ""
	}
	return strings.TrimSpace(m[1])
}

// --- Git helpers ---

func getDiff() string {
	base := os.Getenv("GITHUB_BASE_REF")
	if base == "" {
		base = "main"
	}
	out, err := exec.Command("git", "diff", "origin/"+base+"...HEAD", "--", "README.md").Output()
	if err == nil && len(out) > 0 {
		return string(out)
	}
	out, err = exec.Command("git", "diff", "HEAD~1", "--", "README.md").Output()
	if err == nil {
		return string(out)
	}
	return ""
}

func getChangedFiles() []string {
	base := os.Getenv("GITHUB_BASE_REF")
	if base == "" {
		base = "main"
	}
	out, err := exec.Command("git", "diff", "--name-only", "origin/"+base+"...HEAD").Output()
	if err == nil && len(out) > 0 {
		return splitLines(string(out))
	}
	out, err = exec.Command("git", "diff", "--name-only", "HEAD~1").Output()
	if err == nil {
		return splitLines(string(out))
	}
	return nil
}

func splitLines(s string) []string {
	var lines []string
	for _, l := range strings.Split(strings.TrimSpace(s), "\n") {
		l = strings.TrimSpace(l)
		if l != "" {
			lines = append(lines, l)
		}
	}
	return lines
}

// --- Diff parsing ---

func parseDiffEntries(diff string) (added, removed []entry) {
	for _, line := range strings.Split(diff, "\n") {
		if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
			content := strings.TrimSpace(line[1:])
			if e, ok := parseEntry(content); ok {
				added = append(added, e)
			}
		}
		if strings.HasPrefix(line, "-") && !strings.HasPrefix(line, "---") {
			content := strings.TrimSpace(line[1:])
			if e, ok := parseEntry(content); ok {
				removed = append(removed, e)
			}
		}
	}
	return
}

func parseEntry(line string) (entry, bool) {
	m := reEntry.FindStringSubmatch(line)
	if m == nil {
		return entry{}, false
	}
	return entry{name: m[1], url: m[2], description: m[3], raw: line}, true
}

// --- URL helpers ---

func normalizeURL(u string) string {
	u = strings.TrimRight(u, "/")
	return strings.ToLower(u)
}

func extractRepoName(rawURL string) string {
	parts := strings.Split(strings.Trim(rawURL, "/"), "/")
	if len(parts) >= 2 {
		return parts[len(parts)-1]
	}
	return ""
}

// --- README parsing ---

func getCategoryItemCount(readmePath, entryURL string) (category string, count int) {
	data, err := os.ReadFile(readmePath)
	if err != nil {
		return "unknown", -1
	}

	var currentCat string
	var catItems int
	var foundCat string

	for _, line := range strings.Split(string(data), "\n") {
		trimmed := strings.TrimSpace(line)
		if m := reHeading.FindStringSubmatch(trimmed); m != nil {
			if foundCat != "" {
				break // passed the target category
			}
			currentCat = m[1]
			catItems = 0
		}
		if reEntry.MatchString(trimmed) {
			catItems++
			if strings.Contains(trimmed, entryURL) {
				foundCat = currentCat
			}
		}
	}

	if foundCat == "" {
		return "unknown", -1
	}
	return foundCat, catItems
}

// --- Output ---

func outputResults(results, warnings []string, hasFail bool) {
	var lines []string
	lines = append(lines, "## PR Diff Validation", "")

	if len(results) > 0 {
		lines = append(lines, "### Content checks", "")
		lines = append(lines, results...)
		lines = append(lines, "")
	}

	if len(warnings) > 0 {
		lines = append(lines, "### Warnings", "")
		lines = append(lines, warnings...)
		lines = append(lines, "")
	}

	if hasFail {
		lines = append(lines, "---")
		lines = append(lines, "> **Action needed:** one or more content checks failed. Please review the [contribution guidelines](https://github.com/avelino/awesome-go/blob/main/CONTRIBUTING.md).")
	}

	lines = append(lines, "")
	lines = append(lines, "_Automated diff validation — does not replace maintainer review._")

	comment := strings.Join(lines, "\n")
	setOutput("diff_comment", comment)
	setOutput("diff_fail", boolStr(hasFail))
}

func setOutput(name, value string) {
	path := os.Getenv("GITHUB_OUTPUT")
	if path == "" {
		fmt.Printf("%s=%s\n", name, value)
		return
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	fmt.Fprintf(f, "%s<<EOF\n%s\nEOF\n", name, value)
}

func icon(ok bool) string {
	if ok {
		return "\u2705"
	}
	return "\u274C"
}

func warnIcon(ok bool) string {
	if ok {
		return "\u2705"
	}
	return "\u26A0\uFE0F"
}

func boolStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
