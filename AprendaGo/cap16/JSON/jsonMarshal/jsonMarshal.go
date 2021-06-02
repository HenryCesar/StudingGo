package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	Contabancaria float64
}

func main() {

	jamesbond := Pessoa{
		Nome:          "James",
		Sobrenome:     "Bond",
		Idade:         40,
		Profissao:     "Agente Secreto",
		Contabancaria: 100000.50,
	}

	darthvader := Pessoa{"Anakin", "Skywalker", 50, "Vilão Intergalático", 500000000000.83}

	j, err := json.Marshal(jamesbond)
	if err != nil {
		fmt.Println(err)
	}
	d, err := json.Marshal(darthvader)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(j))
	fmt.Println(string(d))

}
