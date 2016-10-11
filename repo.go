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

type content struct {
	Body string
}

func generateHTML() {
	// Update repo
	exec.Command("git", "checkout", "-f").Output()
	exec.Command("git", "pull").Output()

	input, _ := ioutil.ReadFile("./README.md")
	body := string(blackfriday.MarkdownCommon(input))
	c := &content{Body: body}

	t := template.Must(template.ParseFiles("tmpl/tmpl.html"))
	f, _ := os.Create("tmpl/index.html")
	t.Execute(f, c)
}

func hookHandler(w http.ResponseWriter, r *http.Request) {
	go generateHTML()
	w.Write([]byte("Done!\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hook", hookHandler)
	http.ListenAndServe(":9000", r)
}
