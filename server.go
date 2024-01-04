package main

import (
	"fmt"
	"net/http"
)

func TodoServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Drink coffee")
}
