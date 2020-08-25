package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
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
