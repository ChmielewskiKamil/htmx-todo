package main

import "net/http"

func main() {
	handler := http.HandlerFunc(TodoServer)
	http.ListenAndServe("127.0.0.1:8080", handler)
}
