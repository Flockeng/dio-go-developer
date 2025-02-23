package main

import (
	"errors"
	"fmt"
)

func soma(a, b float64) float64 {
	return a + b
}

func subtracao(a, b float64) float64 {
	return a - b
}

func multiplicacao(a, b float64) float64 {
	return a * b
}

func divisao(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero não é permitida")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Soma: ", soma(10, 5))
	fmt.Println("Subtração: ", subtracao(10, 5))
	fmt.Println("Multiplicação: ", multiplicacao(10, 5))
	resultado, err := divisao(10, 5)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Divisão: ", resultado)
	}
}
