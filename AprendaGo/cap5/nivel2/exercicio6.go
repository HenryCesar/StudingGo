package main

import "fmt"

const (
	_ = iota + 2021
	b
	c
	d
	e
)

func main() {
	fmt.Println("Proximos 4 anos:", b, c, d, e)
}
