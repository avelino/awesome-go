package markdown

import (
	"bytes"
	"io"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// urlAttributes are the attributes that carry a navigation or resource URL.
var urlAttributes = map[string]struct{}{
	"href":       {},
	"src":        {},
	"srcset":     {},
	"action":     {},
	"poster":     {},
	"formaction": {},
}

// DangerousURL reports whether raw is a URL whose scheme can execute script
// or load a non-web resource. html.WithUnsafe() is required for README badges,
// and it also disables goldmark's own scheme filter, so this check runs on
// the rendered attribute value (entity references already resolved).
func DangerousURL(raw string) bool {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" || strings.HasPrefix(trimmed, "#") {
		return false
	}
	// Strip characters browsers ignore inside a scheme.
	cleaned := strings.Map(func(r rune) rune {
		switch r {
		case '\t', '\n', '\r', '\f', ' ':
			return -1
		default:
			return r
		}
	}, trimmed)
	u, err := url.Parse(cleaned)
	if err != nil {
		return true
	}
	switch strings.ToLower(u.Scheme) {
	case "", "http", "https", "mailto":
		return false
	default:
		return true
	}
}

// sanitizeURLs blanks dangerous URL attributes in a goldmark HTML fragment.
// Tokens that are safe are copied unchanged so existing markup is preserved.
func sanitizeURLs(src []byte) ([]byte, error) {
	z := html.NewTokenizer(bytes.NewReader(src))
	var out bytes.Buffer
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			if err := z.Err(); err != io.EOF {
				return nil, err
			}
			return out.Bytes(), nil
		}
		raw := append([]byte(nil), z.Raw()...)
		if tt == html.StartTagToken || tt == html.SelfClosingTagToken {
			if rewritten, ok := rewriteDangerousTag(z, tt); ok {
				out.Write(rewritten)
				continue
			}
		}
		out.Write(raw)
	}
}

func rewriteDangerousTag(z *html.Tokenizer, tt html.TokenType) ([]byte, bool) {
	name, hasAttr := z.TagName()
	if !hasAttr {
		return nil, false
	}
	tag := string(name)
	var attrs []html.Attribute
	changed := false
	for {
		key, val, more := z.TagAttr()
		attr := html.Attribute{Key: string(key), Val: string(val)}
		if _, isURL := urlAttributes[strings.ToLower(attr.Key)]; isURL && DangerousURL(attr.Val) {
			attr.Val = ""
			changed = true
		}
		attrs = append(attrs, attr)
		if !more {
			break
		}
	}
	if !changed {
		return nil, false
	}
	var buf bytes.Buffer
	buf.WriteByte('<')
	buf.WriteString(tag)
	for _, attr := range attrs {
		buf.WriteByte(' ')
		buf.WriteString(attr.Key)
		buf.WriteString(`="`)
		buf.WriteString(html.EscapeString(attr.Val))
		buf.WriteByte('"')
	}
	if tt == html.SelfClosingTagToken {
		buf.WriteString(" /")
	}
	buf.WriteByte('>')
	return buf.Bytes(), true
}
