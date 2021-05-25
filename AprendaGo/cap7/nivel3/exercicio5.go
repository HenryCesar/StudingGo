package main

import "fmt"

func main() {
	fmt.Println("Resto da divisão por 4:")
	for x := 10; x <= 100; x++ {
		fmt.Printf("\n%v / 4 = %v", x, x%4)
	}
}
