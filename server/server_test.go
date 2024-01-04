package server_test

import (
	// "net/http"
	"net/http/httptest"
	"testing"
)

func TestTodoServer(t *testing.T) {
	t.Run("it shows a task name", func(t *testing.T) {
		// request, _ := http.NewRequest(http.MethodGet, "/active", nil)
		response := httptest.NewRecorder()

		// TodoServer(response, request)

		got := response.Body.String()
		want := "<ul><li>Drink Coffee</li></ul>"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
