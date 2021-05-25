// for range

package main

import "fmt"

func main() {
	slice := []string{"banana", "maçã", "jaca", "pêssego"}

	for índice, valor := range slice {
		fmt.Println("No índice", índice, "temos o valor", valor)
	}

	slice = append(slice, "melancia")
	fmt.Println(" ")

	for índice, valor := range slice {
		fmt.Println("No índice", índice, "temos o valor", valor)
	}
}
