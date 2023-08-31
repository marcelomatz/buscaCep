package handlers

import (
	"fmt"
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica se a request foi para a rota correta
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Carrega o arquivo HTML e verifica se ocorreu algum erro
	index, err := os.ReadFile("frontend/index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading index.html: %v", err), http.StatusInternalServerError)
		return
	}

	// Escreve o conteúdo do HTML na resposta
	w.Write(index)
}
