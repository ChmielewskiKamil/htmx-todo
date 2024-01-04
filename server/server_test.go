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

type StubTaskStore struct {
	// @TODO: This is under the assumption that there is only one task per status
	tasks map[string]string
}

func (s *StubTaskStore) GetTask(taskStatus string) string {
	return s.tasks[taskStatus]
}

func TestGETTasks(t *testing.T) {
	store := StubTaskStore{
		map[string]string{
			"/active":    "Learn GO",
			"/completed": "Drink Coffee",
		},
	}
	server := &server.TaskServer{&store}

	t.Run("it shows an active task name", func(t *testing.T) {
		request := newGetTaskRequest(activeTasks)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "<li>Learn GO</li>")
	})

	t.Run("it shows a completed task name", func(t *testing.T) {
		request := newGetTaskRequest(completedTasks)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "<li>Drink Coffee</li>")
	})

	t.Run("it returns 404 on missing url", func(t *testing.T) {
		request := newGetTaskRequest("/missingURL")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatusCode(t, response.Code, http.StatusNotFound)
	})
}

func assertStatusCode(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got status %d, wanted %d", got, want)
	}
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
