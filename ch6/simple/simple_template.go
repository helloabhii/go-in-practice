package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}

type Page struct {
	Title, Content string
}

func displayPage(rw http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin' da castle.",
	}
	t := template.Must(template.ParseFiles("simple/simple.html"))
	t.Execute(rw, p)
}
