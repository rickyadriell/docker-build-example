package main

import (
	"log"
	"net/http"

	"github.com/codecplyre/docker-build-golang/golang/multistage/handler"
)

func main() {
	http.HandleFunc("/", handler.HomeHandler)
	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
