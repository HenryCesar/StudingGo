package main

import "fmt"

var x int

func main() {
	x = 200
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t%b\t%#x", y, y, y)
}
