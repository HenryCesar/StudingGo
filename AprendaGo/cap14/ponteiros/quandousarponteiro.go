package main

import "fmt"

func main() {

	x := 11
	estarecebeovalor(x)
	estarecebeiumponteiro(&x)

	fmt.Println(x)
}

func estarecebeovalor(x int) {
	x++
	fmt.Println("Na função:", x)
}

func estarecebeiumponteiro(x *int) {
	*x++
	fmt.Println("Na função:", *x)
}
