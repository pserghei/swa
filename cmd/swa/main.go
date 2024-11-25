package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("."))
	mux.Handle("GET /", fs)

	log.Print("Starting server...")
	http.ListenAndServe(":8080", mux)
}
