/* Aqui vemos os métodos da Scale e Abs reescritos como funções.

Mais uma vez, tente remover o * a partir da linha 16.
Você pode ver porque as mudanças de comportamento ocorrem?
O que mais você precisa mudar para o exemplo compilar?  */

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
