// Operadores lógicos condicionais

package main

import "fmt"

func main() {

	x := 4

	if x%2 == 0 && x%3 == 0 {
		fmt.Println("[x] é multiplo de 2 e 3")
	} else if x%2 == 0 && x%3 != 0 {
		fmt.Println("[x] é multiplo de 2")
	} else if x%2 != 0 && x%3 == 0 {
		fmt.Println("[x] é multiplo de 3")
	} else {
		fmt.Println("Não é múltiplo nem de 2 e nem de 3")

	}
}
