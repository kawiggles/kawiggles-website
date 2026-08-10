package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "home.html")
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

func renderTemplate(w http.ResponseWriter, path string) {
	if strings.ContainsAny(path, "{}*") {
		panic("Server does not permit and URL parameters.")
	}

	fp := filepath.Join("templates", path)

	tmpl, err := template.ParseFiles("templates/layout.html", fp)
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, nil)
}

type SpotifyData struct {
}

func getSpotifyData() {
}
