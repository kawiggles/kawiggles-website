package main

import (
	"net/http"
)

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "projects.html", map[string]any{ "Title": "Projects", })
}
