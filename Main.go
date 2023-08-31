package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ViaCep struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
	Ibge       string `json:"ibge"`
	Gia        string `json:"gia"`
	Ddd        string `json:"ddd"`
	Siafi      string `json:"siafi"`
}

// isEmpty retorna true se todas as propriedades da estrutura ViaCep estiverem vazias
func (v ViaCep) isEmpty() bool {
	return v.Cep == "" && v.Logradouro == "" && v.Bairro == "" && v.Localidade == "" && v.Uf == "" && v.Ibge == "" && v.Gia == "" && v.Ddd == "" && v.Siafi == ""
}

func main() {
	http.HandleFunc("/", BuscaCepHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("ListenAndServe error: %v", err)
	}
}

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
	viaCep, err := BuscaCep(cepParam)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting CEP details: %v", err), http.StatusInternalServerError)
		return
	}
	if viaCep.isEmpty() {
		http.Error(w, "CEP not found", http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(viaCep)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		return
	}
}

func BuscaCep(cep string) (*ViaCep, error) {
	url := "https://viacep.com.br/ws/" + cep + "/json/"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request error: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body error: %w", err)
	}
	var viaCep ViaCep
	err = json.Unmarshal(body, &viaCep)
	if err != nil {
		return nil, fmt.Errorf("JSON unmarshalling error: %w", err)
	}
	return &viaCep, nil
}
