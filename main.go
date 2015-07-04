package main

import (
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("html", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.Handle("/lobby", &templateHandler{filename: "lobby.html"})
	http.Handle("/game", &templateHandler{filename: "game.html"})
	http.ListenAndServe(":8080", nil)
}
