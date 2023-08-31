package handlers_test

import (
	"http_buscaCep/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestBuscaCepHandle faz uma requisição HTTP GET para o servidor HTTP criado pelo teste de integração.
// Neste caso, o parâmetro cep é informado, então o servidor deve retornar um status 200 OK.
func TestBuscaCepHandle(t *testing.T) {
	req, err := http.NewRequest("GET", "/?cep=01001000", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	w := httptest.NewRecorder()
	handlers.BuscaCepHandle(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.StatusCode)
	}
}

// TestBuscaCepHandle_NoCep faz uma requisição HTTP GET para o servidor HTTP criado pelo teste de integração.
// Neste caso, o parâmetro cep não é informado, então o servidor deve retornar um erro 400 Bad Request.
func TestBuscaCepHandle_NoCep(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	w := httptest.NewRecorder()
	handlers.BuscaCepHandle(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status BadRequest; got %v", resp.StatusCode)
	}
}

// TestBuscaCepHandle_InvalidCep faz uma requisição HTTP GET para o servidor HTTP criado pelo teste de integração.
// Neste caso, o parâmetro cep é informado, mas é um CEP inválido, então o servidor deve retornar um erro 500.
func TestBuscaCepHandle_InvalidLengthCep(t *testing.T) {
	req, err := http.NewRequest("GET", "/?cep=1234", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	w := httptest.NewRecorder()
	handlers.BuscaCepHandle(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status BadRequest; got %v", resp.StatusCode)
	}
}
