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
	pessoa2 := profissional{
		pessoa: pessoa{
			"Maricota", 31,
		},
		titulo:  "Pizzaoila",
		salario: 10000,
	}
	pessoa3 := pessoa{"Mauricio", 40}
	pessoa4 := profissional{pessoa{"Vanderlei", 50}, "Político", 100000}
	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
	fmt.Println(pessoa3)
	fmt.Println(pessoa4)
}
