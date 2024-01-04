package main

import (
	"log"
	"net/http"

	"htmx-todo/server"
)

func main() {
	server := &server.TaskServer{}
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", server))
}
