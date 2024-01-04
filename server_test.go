package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTodoServer(t *testing.T) {
	t.Run("it shows a task name", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		TodoServer(response, request)

		got := response.Body.String()
		want := "Drink coffee"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
