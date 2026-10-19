package main

import (
	"regexp"
	"sort"
	"strings"
)

// prKind is what a README.md patch does to the list.
type prKind string

const (
	kindPackage    prKind = "package"    // adds at least one link that was not in the README before
	kindRemoval    prKind = "removal"    // only drops links
	kindFormatting prKind = "formatting" // same links before and after (whitespace, bullets, order, wording)
	kindNone       prKind = "none"       // README.md not touched
)

// reLink matches the target of a markdown link to an external page.
var reLink = regexp.MustCompile(`\]\((https?://[^)\s]+)\)`)

type classification struct {
	Kind    prKind
	Added   []string // links on added lines that no removed line had
	Removed []string // links on removed lines that no added line has
}

// classify reads the unified diff of README.md and decides whether the PR
// brings new links in. Only the patch text is read, nothing from the PR is
// executed. A link that shows up on both sides (a moved entry, a fixed bullet,
// a trimmed trailing space) is not new.
func classify(patch string) classification {
	added := map[string]bool{}
	removed := map[string]bool{}
	for _, line := range strings.Split(patch, "\n") {
		// The API patch has no file headers, so every +/- line is content.
		switch {
		case strings.HasPrefix(line, "+"):
			for _, l := range links(line[1:]) {
				added[l] = true
			}
		case strings.HasPrefix(line, "-"):
			for _, l := range links(line[1:]) {
				removed[l] = true
			}
		}
	}

	var c classification
	for l := range added {
		if !removed[l] {
			c.Added = append(c.Added, l)
		}
	}
	for l := range removed {
		if !added[l] {
			c.Removed = append(c.Removed, l)
		}
	}
	sort.Strings(c.Added)
	sort.Strings(c.Removed)

	switch {
	case len(c.Added) > 0:
		c.Kind = kindPackage
	case len(c.Removed) > 0:
		c.Kind = kindRemoval
	default:
		c.Kind = kindFormatting
	}
	return c
}

func links(s string) []string {
	var out []string
	for _, m := range reLink.FindAllStringSubmatch(s, -1) {
		out = append(out, normalizeLink(m[1]))
	}
	return out
}

// normalizeLink makes cosmetic URL edits compare equal.
func normalizeLink(u string) string {
	return strings.ToLower(strings.TrimRight(strings.TrimSpace(u), "/"))
}
