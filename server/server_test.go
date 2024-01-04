package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"htmx-todo/server"
)

const (
	activeTasks    = "/active"
	completedTasks = "/completed"
)

func TestGETTasks(t *testing.T) {
	t.Run("it shows an active task name", func(t *testing.T) {
		request := newGetTaskRequest(activeTasks)
		response := httptest.NewRecorder()

		server.TodoServer(response, request)

		assertResponseBody(t, response.Body.String(), "<li>Learn GO</li>")
	})

	t.Run("it shows a completed task name", func(t *testing.T) {
		request := newGetTaskRequest(completedTasks)
		response := httptest.NewRecorder()

		server.TodoServer(response, request)

		assertResponseBody(t, response.Body.String(), "<li>Drink Coffee</li>")
	})
}

////////////////////////////////////////////////////////////////////
//                            HELPERS			                  //
////////////////////////////////////////////////////////////////////

func newGetTaskRequest(tasksStatus string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, tasksStatus, nil)
	return request
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got response: %q, wanted: %q", got, want)
	}
}
