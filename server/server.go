package server

import (
	"net/http"

	"htmx-todo/templates"
)

func TodoServer(w http.ResponseWriter, r *http.Request) {
	templates.Task(GetTask(r.URL.String())).Render(r.Context(), w)
}

func GetTask(taskStatus string) string {
	if taskStatus == "/completed" {
		return "Drink Coffee"
	}

	if taskStatus == "/active" {
		return "Learn GO"
	}

	return ""
}
