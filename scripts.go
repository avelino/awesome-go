package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/avelino/awesome-go/pkg/markdown"
)

type content struct {
	Body template.HTML
}

// ConvertAndRenderIndex generate site html (index.html) from markdown file
func ConvertAndRenderIndex(srcFilename, outFilename string) error {
	input, err := os.ReadFile(srcFilename)
	if err != nil {
		return err
	}

	body, err := markdown.ToHTML(input)
	if err != nil {
		return err
	}

	c := &content{Body: template.HTML(body)}
	f, err := os.Create(outFilename)
	if err != nil {
		return err
	}

	fmt.Printf("Write Index file: %s\n", outIndexFile)
	if err := tplIndex.Execute(f, c); err != nil {
		return err
	}

	return nil
}
