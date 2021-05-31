package main

import "fmt"

func main() {
	canal := make(chan int)

	go loop100(100, canal)
	for v := range canal {
		fmt.Println("Valor de Canal:", v)
	}
}

func loop100(cap int, s chan<- int) {
	for i := 0; i < cap; i++ {
		s <- i
	}
	close(s)
}
