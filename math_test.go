package main

import "testing"

func TestSoma(t *testing.T) {

	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado de soma é inválido %d. Esperado: %d", total, 30)
	}
}