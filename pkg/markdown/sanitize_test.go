package markdown

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestToHTMLDropsDangerousSchemes(t *testing.T) {
	input := []byte(strings.Join([]string{
		`[plain](javascript:alert(1))`,
		`[named](javascript&colon;alert(2)//)`,
		`[decimal](javascript&#58;alert(3)//)`,
		`[hex](javascript&#x3a;alert(4)//)`,
		`![image](javascript&colon;alert(5))`,
		`<a href="javascript:alert(6)">raw</a>`,
		`<a href="javascript&colon;alert(7)//">raw entity</a>`,
		`[ok](https://example.local)`,
		`<a href="https://example.local">badge</a>`,
	}, "\n\n"))

	got, err := ToHTML(input)
	if err != nil {
		t.Fatalf("ToHTML() error = %v", err)
	}
	doc, err := html.Parse(strings.NewReader(string(got)))
	if err != nil {
		t.Fatalf("parse html: %v", err)
	}
	var hrefs, srcs []string
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				switch attr.Key {
				case "href":
					hrefs = append(hrefs, attr.Val)
				case "src":
					srcs = append(srcs, attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)

	for _, href := range hrefs {
		if DangerousURL(href) {
			t.Errorf("dangerous href survived: %q", href)
		}
	}
	for _, src := range srcs {
		if DangerousURL(src) {
			t.Errorf("dangerous src survived: %q", src)
		}
	}
	if !strings.Contains(string(got), `href="https://example.local"`) {
		t.Errorf("safe links were removed:\n%s", got)
	}
}

func TestDangerousURL(t *testing.T) {
	cases := []struct {
		raw  string
		want bool
	}{
		{"https://example.local", false},
		{"http://example.local/path", false},
		{"mailto:hi@example.local", false},
		{"#contents", false},
		{"", false},
		{"javascript:alert(1)", true},
		{"JavaScript:alert(1)", true},
		{" javascript:alert(1)", true},
		{"java\tscript:alert(1)", true},
		{"data:text/html,hi", true},
		{"vbscript:msgbox(1)", true},
	}
	for _, tc := range cases {
		if got := DangerousURL(tc.raw); got != tc.want {
			t.Errorf("DangerousURL(%q) = %v, want %v", tc.raw, got, tc.want)
		}
	}
}
