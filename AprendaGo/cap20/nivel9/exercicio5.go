package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var contador int64 = 0

const qGoroutines int = 20

func main() {
	criarGoroutines(qGoroutines)
	wg.Wait()
	fmt.Println("Quantidade de GoRoutines:", qGoroutines, "Contador:", atomic.LoadInt64(&contador))
}

func criarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			wg.Done()
		}()
	}

}
