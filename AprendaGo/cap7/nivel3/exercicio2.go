package main

import "fmt"

func main() {
	x := 65

	for ; x < 90; x++ {
		fmt.Printf("%#U\n", x)
	}
}
