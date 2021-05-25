// Usar var para declarar três variáveis. (Package-level scope)
// Não atribuir valores a estas variáveis.

package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
