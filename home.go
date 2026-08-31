package main

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "home.html", map[string]any{ "Title": "Home", })
}
