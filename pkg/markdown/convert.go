package markdown

import (
	"bytes"

	"github.com/avelino/awesome-go/pkg/slug"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

// sanitizer is shared by ToHTML; bluemonday policies are safe for concurrent use.
var sanitizer = newSanitizer()

// ToHTML converts markdown byte slice to a HTML byte slice.
//
// README.md relies on inline HTML (the logo, the sponsors table, the collapsible
// contents list), so the renderer runs with html.WithUnsafe and the output is
// untrusted. It is passed through an allow-list sanitizer before being returned.
func ToHTML(markdown []byte) ([]byte, error) {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(), // generate heading IDs for content navigation
		),
		goldmark.WithRendererOptions(
			html.WithXHTML(),
			html.WithUnsafe(), // allow inline HTML
		),
	)

	ctx := parser.NewContext(
		parser.WithIDs(&IDGenerator{}), // register custom ID generator
	)

	var buf bytes.Buffer
	if err := md.Convert(markdown, &buf, parser.WithContext(ctx)); err != nil {
		return nil, err
	}

	return sanitizer.SanitizeBytes(buf.Bytes()), nil
}

// newSanitizer builds the allow-list used on the rendered markdown.
//
// The starting point is bluemonday's user-generated-content policy, which keeps
// the list's real markup (links, headings, lists, tables, images) and strips
// scripts, event handlers and dangerous URL schemes. The additions below restore
// the presentational attributes README.md already uses, so the homepage markup
// does not change.
func newSanitizer() *bluemonday.Policy {
	p := bluemonday.UGCPolicy()
	// UGCPolicy appends rel="nofollow" to every link; the list's own links keep
	// their original form.
	p.RequireNoFollowOnLinks(false)
	p.AllowAttrs("align").Globally()
	p.AllowAttrs("colspan", "rowspan").OnElements("td", "th")
	p.AllowAttrs("cellpadding", "cellspacing", "border").OnElements("table")
	return p
}

// IDGenerator for goldmark to provide IDs more similar to GitHub's IDs on markdown parsing
type IDGenerator struct {
	used map[string]bool
}

// Generate an ID
func (g *IDGenerator) Generate(value []byte, _ ast.NodeKind) []byte {
	return []byte(slug.Generate(string(value)))
}

// Put an ID to the list of already used IDs
func (g *IDGenerator) Put(value []byte) {
	g.used[util.BytesToReadOnlyString(value)] = true
}
