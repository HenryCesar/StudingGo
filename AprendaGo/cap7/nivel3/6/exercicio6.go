package main

import "fmt"

func main() {
	fmt.Println("Demonstrando o [if]")

	for x := 10; x <= 100; x++ {
		if x < 100 {
			fmt.Printf("\n[if] %v < 100\nYes!", x)
		}
		if x > 100 {
			fmt.Printf("\n[if] %v > 100\nYes!", x)
		}
		if x == 100 {
			fmt.Printf("\n[if] %v == 100\nYes!", x)
		}
	}
}
