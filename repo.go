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

type Content struct {
	Body string
}

func GenerateHtml() {
	// Update repo
	exec.Command("git", "checkout", "-f").Output()
	exec.Command("git", "pull").Output()

	input, _ := ioutil.ReadFile("./README.md")
	body := string(blackfriday.MarkdownCommon([]byte(string(input))))
	c := &Content{Body: body}

	t := template.Must(template.ParseFiles("tmpl/tmpl.html"))
	f, _ := os.Create("tmpl/index.html")
	t.Execute(f, c)
}

func HookHandler(w http.ResponseWriter, r *http.Request) {
	go GenerateHtml()
	w.Write([]byte("Done!\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hook", HookHandler)
	http.ListenAndServe(":9000", r)
}
