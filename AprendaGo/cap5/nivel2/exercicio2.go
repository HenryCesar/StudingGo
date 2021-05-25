package main

import "fmt"

func main() {
	fmt.Println("Demonstrando operadores. Com 10 e 100")
	a := (10 == 100)
	b := (10 != 100)
	c := (10 <= 100)
	d := (10 < 100)
	e := (10 > 100)

	fmt.Printf("==: %v\n", a)
	fmt.Printf("!=: %v\n", b)
	fmt.Printf("<=: %v\n", c)
	fmt.Printf(">=: %v\n", d)
	fmt.Printf("> : %v\n", e)
}
