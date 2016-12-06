package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
)

func TestAlpha(t *testing.T) {
	query := startQuery()

	query.Find("body > ul").Each(func(_ int, s *goquery.Selection) {
		testList(t, s)
	})
}

func TestDuplicatedLinks(t *testing.T) {
	query := startQuery()
	links := make(map[string]bool, 0)

	query.Find("body  a").Each(func(_ int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			log.Printf("expected '%s' href", s)
			t.Fail()
		}

		if links[href] {
			log.Printf("duplicated link '%s'", href)
			t.Fail()
			return
		}

		links[href] = true
	})
}

func testList(t *testing.T, list *goquery.Selection) {
	list.Find("ul").Each(func(_ int, items *goquery.Selection) {
		testList(t, items)
		items.RemoveFiltered("ul")
	})
	checkAlphabeticOrder(t, list)
}

func readme() []byte {
	input, err := ioutil.ReadFile("./README.md")
	if err != nil {
		panic(err)
	}
	html := append([]byte("<body>"), blackfriday.MarkdownCommon(input)...)
	html = append(html, []byte("</body>")...)
	return html
}

func startQuery() *goquery.Document {
	buf := bytes.NewBuffer(readme())
	query, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		panic(err)
	}

	return query
}

func checkAlphabeticOrder(t *testing.T, s *goquery.Selection) {
	items := s.Find("li > a:first-child").Map(func(_ int, li *goquery.Selection) string {
		return strings.ToLower(li.Text())
	})

	sorted := make([]string, len(items))
	copy(sorted, items)
	sort.Strings(sorted)

	for k, item := range items {
		if item != sorted[k] {
			log.Printf("expected '%s' but actual is '%s'", sorted[k], item)
			t.Fail()
		}
	}
}
