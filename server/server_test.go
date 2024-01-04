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

		assertResponseBody(t, response.Body.String(), "<li>Learn GO</li>")
	})

	t.Run("it shows a completed task name", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/completed", nil)
		response := httptest.NewRecorder()

		server.TodoServer(response, request)

		assertResponseBody(t, response.Body.String(), "<li>Drink Coffee</li>")
	})
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got response: %q, wanted: %q", got, want)
	}
}
