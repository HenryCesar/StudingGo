package main

import (
	"fmt"
)

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	cliente1 := cliente{
		nome:      "Henrique",
		sobrenome: "Almeida",
		fumante:   false,
	}

	cliente2 := cliente{"Joana", "Pereira", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
}
