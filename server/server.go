package server

import (
	"net/http"

	"htmx-todo/templates"
)

func TodoServer(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/completed" {
		templates.Task("Drink Coffee").Render(r.Context(), w)
	}

	if r.URL.Path == "/active" {
		templates.Task("Learn GO").Render(r.Context(), w)
	}
}
