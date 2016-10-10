package main

import (
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

type Content struct {
	Body string
}

func GenerateHtml() {
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
