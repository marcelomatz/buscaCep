package domain_test

import (
	"http_buscaCep/backend/domain"
	"testing"
)

func TestIsEmpty(t *testing.T) {

	t.Run("empty structure", func(t *testing.T) {
		emptyViaCep := domain.ViaCep{}
		if emptyViaCep.IsEmpty() != true {
			t.Logf("Testing with empty structure: %+v", emptyViaCep)
			t.Errorf("Expected true for an empty ViaCep structure")
		}
	})

	t.Run("non-empty structure", func(t *testing.T) {
		nonEmptyViaCep := domain.ViaCep{Cep: "12345678", Logradouro: "Rua Imaginária"}
		if nonEmptyViaCep.IsEmpty() != false {
			t.Logf("Testing with non-empty structure: %+v", nonEmptyViaCep)
			t.Errorf("Expected false for a non-empty ViaCep structure")
		}
	})
}
