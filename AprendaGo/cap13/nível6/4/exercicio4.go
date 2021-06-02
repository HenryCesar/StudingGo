package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) serhumano() {
	fmt.Println("Esse ser humano se chama:", p.nome, p.sobrenome)
}

func main() {

	rafael := pessoa{
		"Rafael",
		"Andrades",
		20,
	}
	rafael.serhumano()

}
