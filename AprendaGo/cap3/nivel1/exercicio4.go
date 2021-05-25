// Criar um tipo, o subjacente deve ser int.
// Criar uma variável para esse tipo, como "x".

package main

import "fmt"

type(
	tipo int
)

var x tipo

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}