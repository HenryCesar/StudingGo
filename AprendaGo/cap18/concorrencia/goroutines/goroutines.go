package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	go func1()
	func2()

	wg.Wait()
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1:", i)
	}
	wg.Done()
}
func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2:", i)
	}
}
