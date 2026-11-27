package markdown

import (
	"html/template"
	"strings"
	"testing"
)

func TestToHTML(t *testing.T) {
	input := []byte(
		`## some headline
followed by some paragraph with [a link](https://example.local)
and some list:
- first
- second
  - nested on second level
	- nested on third level
	- ~~strikethrough~~
  - yet another second level item, **but** with a [a link](https://example.local)
- end

### h3 headline/header

<a href="https://example.local">embedded HTML is allowed</a>
  `,
	)
	expected := []byte(
		`<h2 id="some-headline">some headline</h2>
<p>followed by some paragraph with <a href="https://example.local">a link</a>
and some list:</p>
<ul>
<li>first</li>
<li>second
<ul>
<li>nested on second level
<ul>
<li>nested on third level</li>
<li><del>strikethrough</del></li>
</ul>
</li>
<li>yet another second level item, <strong>but</strong> with a <a href="https://example.local">a link</a></li>
</ul>
</li>
<li>end</li>
</ul>
<h3 id="h3-headlineheader">h3 headline/header</h3>
<p><a href="https://example.local">embedded HTML is allowed</a></p>`,
	)

	got, err := ToHTML(input)
	if err != nil {
		t.Errorf("ToHTML() error = %v", err)
		return
	}
	if strings.TrimSpace(string(got)) != strings.TrimSpace(string(expected)) {
		t.Errorf("ToHTML() got = %v, want %v", string(got), string(expected))
	}
}

// TestToHTMLSanitizesInlineHTML guards the allow-list: inline HTML in the source
// is rendered, but scripts, event handlers and executable URL schemes are stripped.
func TestToHTMLSanitizesInlineHTML(t *testing.T) {
	input := []byte(`## headline

<script>alert("xss")</script>

<img src="https://example.local/x.png" onerror="alert(1)">

<a href="javascript:alert(1)">executable</a>

<a href="https://example.local/">safe</a>
`)

	got, err := ToHTML(input)
	if err != nil {
		t.Fatalf("ToHTML() error = %v", err)
	}

	for _, dangerous := range []string{"<script", "onerror", "javascript:"} {
		if strings.Contains(string(got), dangerous) {
			t.Errorf("ToHTML() kept %q in the output:\n%s", dangerous, got)
		}
	}

	for _, want := range []string{
		`<img src="https://example.local/x.png"`,
		`<a href="https://example.local/">safe</a>`,
	} {
		if !strings.Contains(string(got), want) {
			t.Errorf("ToHTML() dropped %q from the output:\n%s", want, got)
		}
	}
}

// TestToHTMLIsTrustedByHTMLTemplate checks that the sanitized output is already
// template.HTML, so html/template emits it verbatim instead of escaping it again.
func TestToHTMLIsTrustedByHTMLTemplate(t *testing.T) {
	got, err := ToHTML([]byte("[safe](https://example.local/) <script>alert(1)</script>"))
	if err != nil {
		t.Fatalf("ToHTML() error = %v", err)
	}

	var out strings.Builder
	tpl := template.Must(template.New("body").Parse("{{.Body}}"))
	if err := tpl.Execute(&out, map[string]any{"Body": got}); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}

	if out.String() != string(got) {
		t.Errorf("template re-escaped the body:\n got %s\nwant %s", out.String(), got)
	}
	if !strings.Contains(out.String(), `<a href="https://example.local/">safe</a>`) {
		t.Errorf("link was not rendered as markup: %s", out.String())
	}
	if strings.Contains(out.String(), "<script") {
		t.Errorf("sanitizer let a script through: %s", out.String())
	}
}

func TestToHTMLDuplicateHeadingIDs(t *testing.T) {
	got, err := ToHTML([]byte("## Tools\n\n## Tools\n\n## Tools\n"))
	if err != nil {
		t.Fatalf("ToHTML() error = %v", err)
	}
	for _, id := range []string{`id="tools"`, `id="tools-1"`, `id="tools-2"`} {
		if !strings.Contains(string(got), id) {
			t.Errorf("ToHTML() = %v, want it to contain %s", string(got), id)
		}
	}
}

func TestIDGeneratorPutOnZeroValue(t *testing.T) {
	g := &IDGenerator{}
	g.Put([]byte("tools"))
	if got := string(g.Generate([]byte("Tools"), 0)); got != "tools-1" {
		t.Errorf("Generate() after Put = %q, want %q", got, "tools-1")
	}
}
