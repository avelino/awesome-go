package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
)

// memory usage optimizations
const (
	emtyStr  = ""
	git      = "git"
	checkout = "checkout"
	force    = "-f"
	pull     = "pull"

	// options
	readmePath = "./README.md"
	tplPath    = "tmpl/tmpl.html"
	idxPath    = "tmpl/index.html"

	bfHTMLRendererOpts = 0 |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_FRACTIONS |
		blackfriday.HTML_SMARTYPANTS_DASHES |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

	bfMDOpts = 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS |
		blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
		blackfriday.EXTENSION_DEFINITION_LISTS |
		blackfriday.EXTENSION_AUTO_HEADER_IDS
)

var (
	doneResp = []byte("Done!\n")
)

type content struct {
	Body string
}

func generateHTML() {
	// Update repo
	exec.Command(git, checkout, force).Output()
	exec.Command(git, pull).Output()

	input, _ := ioutil.ReadFile(readmePath)
	body := string(
		blackfriday.Markdown(
			input,
			blackfriday.HtmlRenderer(
				bfHTMLRendererOpts,
				emtyStr,
				emtyStr,
			),
			bfMDOpts,
		),
	)
	c := &content{Body: body}

	t := template.Must(template.ParseFiles(tplPath))
	f, _ := os.Create(idxPath)
	t.Execute(f, c)
}

func hookHandler(w http.ResponseWriter, r *http.Request) {
	go generateHTML()
	w.Write(doneResp)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hook", hookHandler)
	http.ListenAndServe(":9000", r)
}
