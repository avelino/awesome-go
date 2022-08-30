package slug

import (
	"strings"

	"github.com/avelino/slugify"
)

func Generate(text string) string {
	// remove slashes to create slugs similar to GitHub's slugs on markdown parsing
	s := strings.ReplaceAll(text, "/", "")
	return slugify.Slugify(strings.TrimSpace(s))
}
