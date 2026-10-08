// check-pr-diff validates the actual changes in a PR against CONTRIBUTING.md rules:
//   - PR modifies only README.md (for package additions)
//   - PR adds or removes exactly one item
//   - Added link matches the forge link declared in PR body
//   - Link text matches the repository/project name
//   - Description ends with a period and is non-promotional
//   - Added name and description contain no raw HTML, event attributes, or javascript/data/vbscript URLs
//   - Added entry URL uses http or https
//   - Category has minimum 3 items after the change
package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	reForgeLink = regexp.MustCompile(`(?i)forge\s+link[^:]*:\s*(https?://(?:github\.com|gitlab\.com|bitbucket\.org)/\S+)`)
	reHeading   = regexp.MustCompile(`^#{2,3}\s+(.+)`)
	reDescSep   = regexp.MustCompile(`^\s+-\s+(.+)$`)
	// Raw HTML tags. The tag name must follow "<" or "</" immediately so a
	// comparison ("a < b") is left alone. A colon after the name keeps
	// autolinks such as "<https://example.com>" from matching.
	reHTMLTag = regexp.MustCompile(`(?i)</?[a-z][a-z0-9-]*(?:\s[^<>]*|/[^<>]*)?>`)
	// HTML event handler attributes: onclick=, onerror=, onload=, and the rest.
	reEventAttr = regexp.MustCompile(`(?i)(?:^|[^a-z0-9])on[a-z]{3,}\s*=`)
	reBadScheme = regexp.MustCompile(`(?i)(?:^|[^a-z0-9])(?:javascript|data|vbscript)\s*:`)
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
		results = append(results, fix("Your PR should add or remove an entry in README.md.", "Edit README.md following the format: `- [project-name](url) - Short description.`"))
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
		warnings = append(warnings, fix("Package addition PRs should only modify README.md.", "If you need other changes, please open a separate PR."))
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
		results = append(results, fix("Each PR should add, remove, or change only **one** package.", "Please split this into separate PRs — one package per PR."))
		hasFail = true
	}

	// 4. Validate added entries
	for _, e := range added {
		if lines, bad := reviewEntrySafety(e); bad {
			results = append(results, lines...)
			hasFail = true
		} else {
			results = append(results, lines...)
		}

		// 4a. Link matches forge link in PR body
		if forgeLink != "" {
			if normalizeURL(e.url) == normalizeURL(forgeLink) {
				results = append(results, icon(true)+" **Link consistency**: README link matches forge link in PR body")
			} else {
				results = append(results, fmt.Sprintf("%s **Link consistency**: README link `%s` does not match forge link `%s`", icon(false), e.url, forgeLink))
				results = append(results, fix("The URL you added to README.md must match the forge link in your PR description.", fmt.Sprintf("Either update the README entry to use `%s`, or update `Forge link:` in your PR body to `%s`.", forgeLink, e.url)))
				hasFail = true
			}
		}

		// 4b. Link text matches repo name
		repoName := extractRepoName(e.url)
		if repoName != "" {
			if strings.EqualFold(e.name, repoName) {
				results = append(results, icon(true)+" **Link text**: matches repository name")
			} else {
				warnings = append(warnings, fmt.Sprintf("%s **Link text**: `%s` differs from repo name `%s`", warnIcon(false), e.name, repoName))
				warnings = append(warnings, fix("The link text should be the exact project name.", fmt.Sprintf("If the project name really is `%s`, this is fine. Otherwise change it to: `- [%s](%s) - ...`", e.name, repoName, e.url)))
			}
		}

		// 4c. Description ends with period
		if strings.HasSuffix(e.description, ".") || strings.HasSuffix(e.description, "!") {
			results = append(results, icon(true)+" **Description**: ends with punctuation")
		} else {
			results = append(results, icon(false)+" **Description**: must end with a period")
			results = append(results, fix("Add a period `.` at the end of the description.", fmt.Sprintf("Change to: `- [%s](%s) - %s.`", e.name, e.url, e.description)))
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
			warnings = append(warnings, fix("Descriptions should be factual and neutral, not promotional.", "Remove superlatives and marketing language. Good: `Lightweight HTTP router for Go.` Bad: `The fastest, best HTTP router ever.`"))
		} else {
			results = append(results, icon(true)+" **Description tone**: no promotional language detected")
		}

		// 4e. Category minimum items
		cat, count := getCategoryItemCount("README.md", e.url)
		if count >= 3 {
			results = append(results, fmt.Sprintf("%s **Category size**: %s has %d items", icon(true), cat, count))
		} else if count > 0 {
			results = append(results, fmt.Sprintf("%s **Category size**: %s has only %d item(s) (minimum 3 required)", icon(false), cat, count))
			results = append(results, fix("Categories must have at least 3 packages.", "Either add more packages to this category in the same PR, or add your package to an existing category that already has 3+ items."))
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
	head := os.Getenv("PR_HEAD_SHA")
	if head == "" {
		head = "HEAD"
	}
	out, err := exec.Command("git", "diff", "origin/"+base+"..."+head, "--", "README.md").Output()
	if err == nil && len(out) > 0 {
		return string(out)
	}
	out, err = exec.Command("git", "diff", head+"~1", "--", "README.md").Output()
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
	head := os.Getenv("PR_HEAD_SHA")
	if head == "" {
		head = "HEAD"
	}
	out, err := exec.Command("git", "diff", "--name-only", "origin/"+base+"..."+head).Output()
	if err == nil && len(out) > 0 {
		return splitLines(string(out))
	}
	out, err = exec.Command("git", "diff", "--name-only", head+"~1").Output()
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
	line = strings.TrimSpace(line)
	var rest string
	switch {
	case strings.HasPrefix(line, "- ["):
		rest = line[len("- ["):]
	case strings.HasPrefix(line, "* ["):
		rest = line[len("* ["):]
	default:
		return entry{}, false
	}
	name, rest, ok := strings.Cut(rest, "](")
	if !ok || name == "" || strings.Contains(name, "]") {
		return entry{}, false
	}
	// The link destination may contain balanced parentheses, as in
	// javascript:alert(1) or a Wikipedia title.
	depth := 1
	urlEnd := -1
	for i := 0; i < len(rest); i++ {
		switch rest[i] {
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 {
				urlEnd = i
			}
		}
		if urlEnd >= 0 {
			break
		}
	}
	if urlEnd < 0 {
		return entry{}, false
	}
	rawURL := rest[:urlEnd]
	m := reDescSep.FindStringSubmatch(rest[urlEnd+1:])
	if m == nil {
		return entry{}, false
	}
	return entry{name: name, url: rawURL, description: m[1], raw: line}, true
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
	var data []byte
	var err error
	head := os.Getenv("PR_HEAD_SHA")
	if head != "" {
		data, err = exec.Command("git", "show", head+":README.md").Output()
	} else {
		data, err = os.ReadFile(readmePath)
	}
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
		if _, ok := parseEntry(trimmed); ok {
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
	delimiter := githubOutputDelimiter(value)
	fmt.Fprintf(f, "%s<<%s\n%s\n%s\n", name, delimiter, value, delimiter)
}

// githubOutputDelimiter returns a per-call heredoc delimiter that is not a
// line of its own inside value. GitHub Actions ends a multiline GITHUB_OUTPUT
// value at the first line equal to the delimiter, so a fixed word such as EOF
// lets that line inject another step output.
func githubOutputDelimiter(value string) string {
	for range 5 {
		buf := make([]byte, 16)
		if _, err := rand.Read(buf); err != nil {
			break
		}
		delimiter := "ghadelim_" + hex.EncodeToString(buf)
		if !outputLineEquals(value, delimiter) {
			return delimiter
		}
	}
	return "ghadelim_fallback"
}

func outputLineEquals(value, delimiter string) bool {
	rest := value
	for {
		line, after, found := strings.Cut(rest, "\n")
		if line == delimiter {
			return true
		}
		if !found {
			return false
		}
		rest = after
	}
}

// reviewEntrySafety fails an added README entry whose name or description
// contains raw HTML, an event handler attribute, or a javascript/data/vbscript
// URL, and fails an entry URL that is not http(s).
func reviewEntrySafety(e entry) (lines []string, fail bool) {
	var problems []string
	if !isHTTPURL(e.url) {
		problems = append(problems, fmt.Sprintf("URL must use http or https, got `%s`", e.url))
	}
	problems = append(problems, textSafetyProblems("name", e.name)...)
	problems = append(problems, textSafetyProblems("description", e.description)...)
	if len(problems) == 0 {
		return []string{icon(true) + " **Entry safety**: name, description, and URL are plain http(s) content"}, false
	}
	for _, problem := range problems {
		lines = append(lines, icon(false)+" **Entry safety**: "+problem)
	}
	lines = append(lines, fix(
		"Names and descriptions cannot contain raw HTML, event handler attributes (onerror=, onclick=, ...), or javascript:, data:, or vbscript: URLs. The entry URL must be http or https.",
		"Use a plain project name, an http(s) link, and a text description. Example: `- [project](https://github.com/org/project) - Short description.`",
	))
	return lines, true
}

func textSafetyProblems(label, value string) []string {
	var problems []string
	if reHTMLTag.MatchString(value) {
		problems = append(problems, label+" contains a raw HTML tag")
	}
	if reBadScheme.MatchString(value) {
		problems = append(problems, label+" contains a javascript:, data:, or vbscript: URL")
	}
	if reEventAttr.MatchString(value) {
		problems = append(problems, label+" contains an HTML event attribute")
	}
	return problems
}

func isHTTPURL(raw string) bool {
	raw = strings.TrimSpace(raw)
	if raw == "" || strings.ContainsAny(raw, " \t\r\n") {
		return false
	}
	u, err := url.Parse(raw)
	if err != nil || u.Host == "" {
		return false
	}
	switch strings.ToLower(u.Scheme) {
	case "http", "https":
		return true
	default:
		return false
	}
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

func fix(problem, howToFix string) string {
	return fmt.Sprintf("  > **How to fix:** %s\n  > %s", problem, howToFix)
}
