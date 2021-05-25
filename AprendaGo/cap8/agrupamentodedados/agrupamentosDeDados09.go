//Maps

package main

import (
	"fmt"
)

func main() {

	amigos := map[string]int{
		"alfredo": 5551234,
		"joana":   9996674,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["joana"])

	amigos["gopher"] = 444444

	fmt.Println(amigos)
	fmt.Println(amigos["gopher"])

	if será, ok := amigos["fantasma"]; !ok {
		fmt.Println("não tem!")
	} else {
		fmt.Println(será, ok)
	}

}
