package main

import (
	"github.com/go-martini/martini"
	"github.com/russross/blackfriday"
	"net/http"
)


func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	}) 

	m.Post("/generate", func(r *http.Request) []byte {
		body := r.FormValue("body")
		return blackfriday.MarkdownBasic([]byte(body))
	})

	m.Run()
}