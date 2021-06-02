package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type arquiteto struct {
	pessoa
	especialização  string
	obrasconcluidas int
}

type dentista struct {
	pessoa
	especialização   string
	dentesarrancados int
}

type gente interface {
	oibomdia()
}

func (d dentista) oibomdia() {
	fmt.Printf("Oi, bom dia! Meu nome é %v e eu já arranquei %d dentes!\n", d.nome, d.dentesarrancados)
}

func (a arquiteto) oibomdia() {
	fmt.Printf("Oi, bom dia! Meu nome é %v e eu já fiz %d obras!\n", a.nome, a.obrasconcluidas)
}

func serhumano(g gente) {
	fmt.Println("--- func serhumano ---")
	switch g.(type) {
	case arquiteto:
		fmt.Println(g.(arquiteto).nome, "é uma pessoa que trabalha com", g.(arquiteto).especialização)
	case dentista:
		fmt.Println(g.(dentista).nome, "é uma pessoa que trabalha com", g.(dentista).especialização)
	}
	fmt.Println("Ele diz:")
	g.oibomdia()
}

func main() {

	alfredo := arquiteto{
		pessoa: pessoa{
			nome:      "Alfredo",
			sobrenome: "da Silva",
			idade:     30,
		},
		especialização:  "galpões",
		obrasconcluidas: 10,
	}

	rogério := dentista{
		pessoa: pessoa{
			nome:      "Rogério",
			sobrenome: "Pereira",
			idade:     60,
		},
		especialização:   "tortura",
		dentesarrancados: 7835,
	}

	fmt.Println(alfredo)
	alfredo.oibomdia()
	serhumano(alfredo)

	fmt.Println("-----")

	fmt.Println(rogério)
	rogério.oibomdia()
	serhumano(rogério)
}
