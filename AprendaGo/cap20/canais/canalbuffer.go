package main

import "fmt"

func main() {
	canal := make(chan int, 1)
	canal <- 42
	fmt.Println(<-canal)
}
