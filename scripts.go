package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/avelino/awesome-go/pkg/markdown"
)

// renderIndex generate site html (index.html) from markdown file
func renderIndex(srcFilename, outFilename string) error {
	input, err := os.ReadFile(srcFilename)
	if err != nil {
		return err
	}

	body, err := markdown.ToHTML(input)
	if err != nil {
		return err
	}

	f, err := os.Create(outFilename)
	if err != nil {
		return err
	}

	fmt.Printf("Write Index file: %s\n", outIndexFile)
	data := map[string]interface{}{
		"Body": template.HTML(body),
	}
	if err := tplIndex.Execute(f, data); err != nil {
		return err
	}

	return nil
}
