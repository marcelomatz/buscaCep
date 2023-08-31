package main

import (
	"fmt"
	"http_buscaCep/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.BuscaCepHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("ListenAndServe error: %v", err)
	}
}
