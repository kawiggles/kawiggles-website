package main

import (
	"net/http"
)

func contactHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "contact.html", map[string]any{ "Title": "Contact", })
}
