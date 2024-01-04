package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"htmx-todo/server"
)

func TestTodoServer(t *testing.T) {
	t.Run("it shows a task name", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/active", nil)
		response := httptest.NewRecorder()

		server.TodoServer(response, request)

		got := response.Body.String()
		want := "<li>Drink Coffee</li>"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
