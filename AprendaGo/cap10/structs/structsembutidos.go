package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {

	pessoa1 := pessoa{"Alfredo", 30}
	pessoa2 := profissional{pessoa: pessoa{"Maricota", 31}, titulo: "Pizzaoila", salario: 10000}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
}
