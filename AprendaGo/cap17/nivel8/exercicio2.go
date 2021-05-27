package main

import (
	"encoding/json"
	"fmt"
)

type info struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	Contabancaria float64
}

func main() {
	sb := []byte(`{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente Secreto","Contabancaria":100000.5}`)
	var jamesbond info
	err := json.Unmarshal(sb, &jamesbond)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(jamesbond.Nome)
	fmt.Println(jamesbond)

}
