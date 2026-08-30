package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"kawiggles.com/pages"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", pages.HomeHandler)
	r.Get("/resume", pages.ResumeHandler)
	r.Get("/projects", pages.ProjectsHandler)
	r.Get("/blog", pages.BlogHandler)
	r.Get("/contact", pages.ContactHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}

