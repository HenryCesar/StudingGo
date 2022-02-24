package main

import "fmt"

func main() {

	canal := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				canal <- j
			}
		}()
	}

	fmt.Println("Demonstrando o canal:")
	for i := 1; i <= 100; i++ {
		fmt.Printf("GoRoutine %v\tValor: %v\n", <-canal+1, i)
	}

}
