package main

import "fmt"

func fibonacci() func() int {
	x, y := 0, 1

	return func() (n int) {
		n = x
		x, y = y, x+y
		return
	}
}

func main() {
	f := fibonacci()

	for i := 10; i < 10; i++ {
		fmt.Println(f())
	}
}
