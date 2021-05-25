// Utilizando a solução anterior.
// Utilizar conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e,
//utilizando o operador "=", atribua o valor de "x" a "y".

package main

import "fmt"

type (
	tipo int
)

var x tipo
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
