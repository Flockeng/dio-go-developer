package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	resultado := soma(2, 3)
	esperado := 5.0
	if resultado != esperado {
		t.Errorf("Soma incorreta, esperado %f, obteve %f", esperado, resultado)
	}
}

func TestSubtracao(t *testing.T) {
	resultado := subtracao(5, 3)
	esperado := 2.0
	if resultado != esperado {
		t.Errorf("Subtração incorreta, esperado %f, obteve %f", esperado, resultado)
	}
}

func TestMultiplicacao(t *testing.T) {
	resultado := multiplicacao(2, 3)
	esperado := 6.0
	if resultado != esperado {
		t.Errorf("Multiplicação incorreta, esperado %f, obteve %f", esperado, resultado)
	}
}

func TestDivisao(t *testing.T) {
	resultado, err := divisao(6, 3)
	esperado := 2.0
	if err != nil || resultado != esperado {
		t.Errorf("Divisão incorreta, esperado %f, obteve %f", esperado, resultado)
	}
}

func TestDivisaoPorZero(t *testing.T) {
	_, err := divisao(6, 0)
	if err == nil {
		t.Errorf("Esperado erro de divisão por zero, mas nenhum erro foi retornado")
	}
}
