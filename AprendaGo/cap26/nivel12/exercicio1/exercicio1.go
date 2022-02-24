package main

import (
	"fmt"

	"estudostest.com/m/AprendaGo/cap26/nivel12/exercicio1/quadrado"
)

func main() {
	fmt.Println("Utilizando package quadrado:")
	x := 10
	fmt.Println("10 ao quadrado é:", quadrado.UtilizandoQuadrado(x))
}
