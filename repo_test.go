package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
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

		checkGoreportcard(t, href)
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

type Card struct {
	Checks []struct {
		Description   string        `json:"description"`
		FileSummaries []interface{} `json:"file_summaries"`
		Name          string        `json:"name"`
		Percentage    float64       `json:"percentage"`
		Weight        float64       `json:"weight"`
	} `json:"checks"`
	Files       int    `json:"files"`
	Grade       string `json:"grade"`
	Issues      int    `json:"issues"`
	LastRefresh string `json:"last_refresh"`
	Repo        string `json:"repo"`
}

func checkGoreportcard(t *testing.T, href string) {
	if !strings.Contains(string(href), string("github.com")) {
		return
	}

	r, _ := regexp.Compile(`https?:\/\/[^\/]*\/([^\/]*)\/([^\/]*).*`)
	_r := r.FindStringSubmatch(href)
	repo := fmt.Sprintf("%s/%s", _r[1], _r[2])

	request := gorequest.New()
	_, body, _ := request.Post("http://goreportcard.com/checks").
		Type("form").
		Send(`{"repo":"` + repo + `"}`).
		End()

	var card Card
	err := json.Unmarshal([]byte(body), &card)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range card.Checks {
		if v.Percentage != 1 {
			fmt.Println("Percentage", v.Percentage)
			log.Printf("'%s' not 100% in goreportcard.com", string(href))
			t.Fail()
		}
	}
}
