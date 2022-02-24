package main

import "fmt"

func main() {
	x := 10

	switch {
	case x == 100:
		fmt.Println("x == 100")
	case x < 100:
		fmt.Println("x < 100")
	case x > 100:
		fmt.Println("x > 100")
	}
}
