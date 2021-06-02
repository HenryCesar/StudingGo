// Slice the slice

package main

import (
	"fmt"
)

func main() {

	sabores := []string{"peperoni", "mozzarela", "abacaxi", "quatro queijos", "marg"}

	fatia := sabores[:]
	sabores = append(sabores[:2], sabores[3:]...)

	fmt.Println(fatia)
	fmt.Println(sabores)
}
