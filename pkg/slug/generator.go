package slug

import (
	"strings"

	"github.com/avelino/slugify"
)

// Generate slugs similar to GitHub's slugs on markdown parsing
func Generate(text string) string {
	// FIXME: this is should be like regexp.Replace(`[^-a-zA-Z\d]+`, ``)
	s := strings.ReplaceAll(text, "/", "")
	return slugify.Slugify(strings.TrimSpace(s))
}
