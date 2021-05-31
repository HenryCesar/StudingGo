package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())

	go func() {
		fmt.Println("Eu sou a main goroutine.")
	}()

	wg.Add(1)

	go func1()
	func2()

	wg.Wait()
	fmt.Println("Acabou o código.")
}

func func1() {
	for i := 0; i < 3; i++ {
		fmt.Println("goroutine1:", i)
	}
	wg.Done()
}
func func2() {
	for i := 0; i < 5; i++ {
		fmt.Println("goroutine2:", 2)
	}
}
