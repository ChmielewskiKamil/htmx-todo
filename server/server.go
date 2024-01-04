package server

import (
	"net/http"

	"htmx-todo/templates"
)

func TodoServer(w http.ResponseWriter, r *http.Request) {
	templates.Task().Render(r.Context(), w)
}
