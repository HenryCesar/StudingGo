package main

import "fmt"

func main() {
	x := 10

	y := &x
	fmt.Println("O endereço de", x, "é:", y)
}
