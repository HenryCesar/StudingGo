package main

import "fmt"

func main() {
	defer fmt.Println("Escrevi primeiro, mas tem defer.")

	fmt.Println("Escrevi depois, mas não tem defer.")

}
