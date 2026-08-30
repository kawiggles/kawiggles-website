package pages

import (
	"net/http"
	"html/template"
	"path/filepath"
	"strings"
)

func renderTemplate(w http.ResponseWriter, path string) {
	if strings.ContainsAny(path, "{}*") {
		panic("Server does not permit URL parameters.")
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
