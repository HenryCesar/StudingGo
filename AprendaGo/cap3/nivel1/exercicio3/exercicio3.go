// Utilizar a solução do exercício anterior
// Usar fmt.Sprintf para atribuir todos esses valores a um única variável.

package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Println(s)
}