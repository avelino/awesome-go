package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/russross/blackfriday"
)

func readme() []byte {
	input, err := ioutil.ReadFile("./README.md")
	if err != nil {
		panic(err)
	}
	html := fmt.Sprintf("<body>%s</body>", blackfriday.MarkdownCommon(input))
	htmlByteArray := []byte(html)
	return htmlByteArray
}

func startQuery() *goquery.Document {
	buf := bytes.NewBuffer(readme())
	query, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		panic(err)
	}
	return query
}

type content struct {
	Body string
}

// GenerateHTML generate site html (index.html) from markdown file
func GenerateHTML() (err error) {
	// options
	readmePath := "./README.md"
	tplPath := "tmpl/tmpl.html"
	idxPath := "tmpl/index.html"
	input, _ := ioutil.ReadFile(readmePath)
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.LaxHTMLBlocks
	parser := parser.NewWithExtensions(extensions)

	body := string(markdown.ToHTML(input, parser, nil))
	c := &content{Body: body}
	t := template.Must(template.ParseFiles(tplPath))
	f, err := os.Create(idxPath)
	t.Execute(f, c)
	return
}
