package pages

import (
	"net/http"
)

func ResumeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "resume.html")
}
