package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Profissão string
}

func main() {
	henrique := Pessoa{
		"Henrique",
		"Almeida",
		19,
		"Estagiário",
	}
	caique := Pessoa{
		"Caique",
		"Torres",
		20,
		"Fullstack Junior",
	}

	h, err := json.Marshal(henrique)
	if err != nil {
		fmt.Println(err)
	}
	c, err := json.Marshal(caique)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(h))
	fmt.Println(string(c))
}
