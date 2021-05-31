package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}
type humano interface {
	falar()
}

func (p *pessoa) falar() {
	fmt.Println("Olá, eu sou", p.nome, ", eu tenho", p.idade, "anos de idade.")
}
func dizerAlgumaCoisa(h humano) {
	h.falar()
}

func main() {
	mario := pessoa{"Mario", 20}
	julio := pessoa{"Julio", 32}
	carlos := pessoa{"Carlos", 12}

	mario.falar()

	(&julio).falar()

	dizerAlgumaCoisa(&carlos)

}
