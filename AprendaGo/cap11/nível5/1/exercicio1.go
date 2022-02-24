package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sorvete   []string
}

func main() {

	pessoa1 := pessoa{
		nome:      "Lucas",
		sobrenome: "de Luca",
		sorvete:   []string{"Flocos", "Chocolate", "Creme"},
	}

	pessoa2 := pessoa{
		nome:      "Marcia",
		sobrenome: "Souza",
		sorvete: []string{
			"Napolitano",
			"Flocos",
			"Melão"},
	}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.sorvete {
		fmt.Println("\t", v)
	}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, v := range pessoa2.sorvete {
		fmt.Println("\t", v)
	}
}
