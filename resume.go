package main

import (
	"net/http"
)

func resumeHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "resume.html", map[string]any{ "Title": "Resume", })
}
