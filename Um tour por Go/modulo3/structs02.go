package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2} // nomeando [v] como um tipo Vertex
	v.X = 4           // Adicionando o Valor 4 ao elemento X dentro do tipo Vertex
	fmt.Println(v.X)  // Imprimindo na tela o valor X, dentro do tipo Vertex(v) acessado via "." (ponto)
}

// Campos [Struct] são acessados através de um ponto.
