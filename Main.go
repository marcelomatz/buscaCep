package main

import (
	"http_buscaCep/backend/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/busca", handlers.BuscaCepHandle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
