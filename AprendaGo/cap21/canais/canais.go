package main

import "fmt"

func main() {
	canal := make(chan int)

	go func() {
		canal <- 42
	}()
	fmt.Println(<-canal)
}
