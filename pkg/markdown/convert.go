package markdown

import (
	"bytes"

	"github.com/avelino/awesome-go/pkg/slug"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

// ToHTML converts markdown byte slice to a HTML byte slice
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

	return buf.Bytes(), nil
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
