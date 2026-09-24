package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var templates = map[string]*template.Template{}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	loadTemplates()

	posts, err := parsePosts("./posts")
	if err != nil {
		log.Panic("failed to parse blog posts: %v\n", err)
	}
	blog := Blog { Posts: posts }

	r.Get("/", homeHandler)
	r.Get("/resume", resumeHandler)
	r.Get("/projects", projectsHandler)
	r.Get("/blog", blog.blogHandler)
	r.Get("/contact", contactHandler)
	r.Get("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/robots.txt")
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}

func loadTemplates() {
	pages := []string{"home.html", "resume.html", "projects.html", "blog.html", "contact.html"}

	for _, page := range pages {
		path := filepath.Join("templates", page)
		tmpl := template.Must(template.ParseFiles("templates/layout.html", path))
		templates[page] = tmpl
	}
}

func render(w http.ResponseWriter, page string, data any) {
	tmpl, ok := templates[page]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tmpl.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}
