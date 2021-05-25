// Array
package main

import "fmt"

var x [5]int
var y [6]int

func main() {

	x[0] = 1
	x[1] = 10
	fmt.Println(x[0], x[1])
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y[0], y[1])
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println("Tamanho do array [x]:", len(x))
	fmt.Println("Tamanho do array [y]:", len(y))
}
