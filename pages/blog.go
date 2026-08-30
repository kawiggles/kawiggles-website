package pages

import (
	"html/template"
	"net/http"
	"time"
)

type Post struct {
	Slug 	string
	Title 	string 		`yaml:"title"`
	Date	time.Time	`yaml:"date"`
	Tags	[]string	`yaml:"tags"`
	HTML	template.HTML
}

func BlogHandler(w http.ResponseWriter, r *http.Request) {
}
