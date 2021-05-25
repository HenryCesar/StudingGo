package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // possuí o tipo Vertex
	v2 = Vertex{X: 1}  // Y = 0 de modo implicito
	v3 = Vertex{}      // X = 0 e Y = 0 (implicito)
	p  = &Vertex{1, 2} // possuí o tipo *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

/* A struct literal indica um valor struct recém-alocado, ao enumerar os valores de seus campos.

Você pode listar apenas um subconjunto de campos usando o [Name:] sintaxe. (E a ordem dos campos nomeados é irrelevante.)

O prefixo especial [&] (Ê comercial) constrói um ponteiro para uma struct literal. */
