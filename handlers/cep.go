package handlers

import (
	"encoding/json"
	"fmt"
	"http_buscaCep/domain"
	"net/http"
)

func BuscaCepHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" || len(cepParam) != 8 {
		http.Error(w, "Invalid cep parameter", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	viaCep, err := domain.BuscaCep(cepParam)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting CEP details: %v", err), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(viaCep)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		return
	}
}
