package main

import "fmt"

const ebulicaoKelvin = 373.0
const pontoFusao = 273.0

func main() {
	tempKelvin := ebulicaoKelvin
	tempCelsius := tempKelvin - pontoFusao

	fmt.Printf("Ebulição da água em Kelvin = %g , ebulição da água em celsius = %g.", tempKelvin, tempCelsius)
}
