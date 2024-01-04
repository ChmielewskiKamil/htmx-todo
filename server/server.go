package server

import (
	"net/http"

	"htmx-todo/templates"
)

type TaskStore interface {
	GetTask(taskStatus string) string
}

type TaskServer struct {
	Store TaskStore
}

func (s *TaskServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	task := GetTask(r.URL.String())

	if task == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	templates.Task(task).Render(r.Context(), w)
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
