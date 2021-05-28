package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var contador int = 0

const qGoroutines int = 20

func main() {
	criarGoroutines(qGoroutines)
	wg.Wait()
	fmt.Println("Quantidade de GoRoutines:", qGoroutines, "Contador:", contador)
}

func criarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}

}
