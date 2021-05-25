package main

import "fmt"

const (
	x = 10
	z = 20
)

var (
	y = 20
	c int
	d float64
)

func main() {
	c = x
	fmt.Println("c = x =", c)
	c = y
	fmt.Println("c = y =", c)
	d = x
	fmt.Println("d = x =", d)

	fmt.Println("x, z =", x, z)
}
