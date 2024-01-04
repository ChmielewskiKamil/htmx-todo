package server

import (
	"net/http"

	"github.com/a-h/templ"
	"htmx-todo/templates"
)

func TodoServer(w http.ResponseWriter, r *http.Request) {
	component := templates.Task()

	http.Handle("/active/", templ.Handler(component))
}
