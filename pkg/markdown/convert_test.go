package markdown

import (
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
