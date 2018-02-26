package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"text/template"

	"github.com/gorilla/mux"
	gfm "github.com/shurcooL/github_flavored_markdown"
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
	body := string(gfm.Markdown(input))
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
