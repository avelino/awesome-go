package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/avelino/awesome-go/pkg/markdown"
)

func readme() []byte {
	input, err := os.ReadFile("./README.md")
	if err != nil {
		panic(err)
	}
	html, err := markdown.ConvertMarkdownToHTML(input)
	if err != nil {
		panic(err)
	}
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

type content struct {
	Body template.HTML
}

// GenerateHTML generate site html (index.html) from markdown file
func GenerateHTML() (err error) {
	// options
	readmePath := "./README.md"
	tplPath := "tmpl/tmpl.html"
	idxPath := "tmpl/index.html"
	input, _ := ioutil.ReadFile(readmePath)
	body, _ := markdown.ConvertMarkdownToHTML(input)
	c := &content{Body: template.HTML(body)}
	t := template.Must(template.ParseFiles(tplPath))
	f, err := os.Create(idxPath)
	t.Execute(f, c)
	return
}
