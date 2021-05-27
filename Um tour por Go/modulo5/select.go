/* A instrução select permite uma espera na goroutine sobre as operações de comunicação múltiplas.

O bloco select aguarda até que um de seus cases possam executar, então ele executa esse case.
Ele escolhe um ao acaso se vários estiverem prontos.
*/

package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)

	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)

}
