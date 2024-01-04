package main

import (
	"log"
	"net/http"

	"htmx-todo/server"
)

func main() {
	handler := http.HandlerFunc(server.TodoServer)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", handler))
}
