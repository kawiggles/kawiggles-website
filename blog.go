package main

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v3"
)

type Blog struct {
	Posts	[]Post
}

type Post struct {
	Slug 	string
	Title 	string 		`yaml:"title"`
	Date	time.Time	`yaml:"date"`
	Tags	[]string	`yaml:"tags"`
	HTML	template.HTML
}

func (b *Blog) blogHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "blog.html", map[string]any{
		"Title":	"Blog",
		"Posts":	b.Posts,
	})
}

func parsePosts(dir string) ([]Post, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var posts []Post

	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())
		post, err := parsePost(path)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func parsePost(path string) (Post, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Post{}, err
	}
	parts := strings.SplitN(string(raw), "---", 3)

	var post Post
	err = yaml.Unmarshal([]byte(parts[1]), &post)
	if err != nil {
		return Post{}, err
	}

	var buf bytes.Buffer
	err = goldmark.Convert([]byte(parts[2]), &buf)
	if err != nil {
		return Post{}, err
	}

	post.HTML = template.HTML(buf.String())
	post.Slug = strings.TrimSuffix(filepath.Base(path), ".md")

	return post, nil
}
