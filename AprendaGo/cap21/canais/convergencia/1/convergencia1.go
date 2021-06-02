package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)

	go envia(par, impar)
	go recebe(par, impar, converge)

	for v := range converge {
		fmt.Println("Valor Recebido:", v)
	}

}

func envia(p, i chan int) {
	x := 100
	for n := 0; n < x; n++ {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	close(p)
	close(i)
}

func recebe(p, i, c chan int) {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		for v := range p {
			c <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range i {
			c <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(c)
}
