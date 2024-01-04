package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"htmx-todo/server"
)

func TestGETTasks(t *testing.T) {
	t.Run("it shows an active task name", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/active", nil)
		response := httptest.NewRecorder()

		server.TodoServer(response, request)

		got := response.Body.String()
		want := "<li>Learn GO</li>"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("it shows a completed task name", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/completed", nil)
		response := httptest.NewRecorder()

		server.TodoServer(response, request)

		got := response.Body.String()
		want := "<li>Drink Coffee</li>"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
