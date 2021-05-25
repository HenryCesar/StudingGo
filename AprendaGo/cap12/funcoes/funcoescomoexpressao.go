package main

import "fmt"

func main() {

	x := 386
	y := func(x int) int {
		return x * 1022
	}

	fmt.Println(x, "vezes 1022 é:", y(x))
}
