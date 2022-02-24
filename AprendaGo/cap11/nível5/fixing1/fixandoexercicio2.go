package main

import "fmt"

type funcionario struct {
	nome      string
	sobrenome string
	ra        int
}

func main() {

	mapafuncionario := make(map[string]funcionario)

	mapafuncionario["Almeida"] = funcionario{
		nome:      "Henrique",
		sobrenome: "Almeida",
		ra:        190804,
	}
	mapafuncionario["Torres"] = funcionario{
		nome:      "Luan",
		sobrenome: "Torres",
		ra:        204230,
	}

	for _, valor := range mapafuncionario {
		fmt.Println(valor.nome, valor.sobrenome, "-", valor.ra)
	}
}
