package main

import (
	"log"
	"net/http"

	"htmx-todo/server"
)

type InMemoryTaskStore struct{}

func (s *InMemoryTaskStore) GetTask(taskStatus string) string {
	return "lol?"
}

func main() {
	server := &server.TaskServer{&InMemoryTaskStore{}}
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", server))
}
