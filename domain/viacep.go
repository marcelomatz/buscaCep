package domain

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
	if viaCep.isEmpty() {
		return nil, fmt.Errorf("CEP %s not found", cep)
	}
	return &viaCep, nil
}
