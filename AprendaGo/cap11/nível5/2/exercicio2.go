package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	mapa := map[string]pessoa{
		"Souza": {
			nome:      "Marcia",
			sobrenome: "Souza",
			sabores: []string{
				"Napolitano",
				"Flocos",
				"Melão"},
		},
		"De Luca": {
			nome:      "Lucas",
			sobrenome: "de Luca",
			sabores:   []string{"Flocos", "Chocolate", "Creme"},
		},
	}

	for _, valor := range mapa {
		fmt.Println("Meu nome é", valor.nome, " e meus sorvetes favoritos são:")
		for _, v := range valor.sabores {
			fmt.Println("-", v)
		}
	}
}
