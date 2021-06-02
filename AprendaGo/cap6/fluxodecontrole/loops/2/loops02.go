//Repetição hierárquica
package main

import (
	"fmt"
)

func main() {
	for mes := 1; mes <= 12; mes++ {
		fmt.Println("Mês: ", mes)
		for dia := 0; dia <= 30; dia++ {
			fmt.Println("Dia: ", dia)
		}
		fmt.Println(" ")
	}
}
