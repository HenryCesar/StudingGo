package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 { //Função que efetua outra função com os números 3 e 4
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 { //Efetua uma expressão específica (raiz)
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) //hypot irá efetuar a raiz de 5*5 + 12*12

	fmt.Println(compute(hypot))    //compute irá efetuar o "hypot" de 3 e 4
	fmt.Println(compute(math.Pow)) //compute irá efetuar o "math.Pow" de 3 e 4

}

// Funções são valores também. Elas podem ser passadas assim como outros valores.

// Funções valores podem ser usadas como argumentos de funções e retornar valores.
