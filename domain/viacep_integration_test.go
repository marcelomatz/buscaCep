package domain_test

import (
	"http_buscaCep/domain"
	"testing"
)

// Chama a função BuscaCep com um valor de "CEP" conhecido.
// Verifica se houve erro na execução. Se houve algum erro, o teste falha imediatamente.
// Em seguida, o teste verifica se o retorno de BuscaCep é nil. Se for nil, é um erro inesperado, então o teste falha.
// Finalmente, o teste verifica se o retorno de BuscaCep não é vazio, o que significa que a função BuscaCep encontrou
// e retornou os detalhes do CEP. Se o objeto voltar vazio (IsEmpty é true), então o teste falha.
func TestBuscaCepIntegration(t *testing.T) {
	cep := "01001001" // Use um CEP existente para o teste
	c, err := domain.BuscaCep(cep)
	if err != nil {
		t.Fatalf("BuscaCep error: %v", err)
	}
	if c == nil {
		t.Fatalf("BuscaCep returned nil ViaCep")
	}
	if c.IsEmpty() {
		t.Errorf("BuscaCep returned empty ViaCep")
	}
}
