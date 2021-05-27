/* Temos visto como canais são bons para a comunicação entre goroutines.
Mas e se nós não precisássemos de comunicação?
E se nós apenas quiséssemos ter certeza que uma única goroutine pode acessar uma variável de cada vez para evitar conflitos?

Este conceito é chamdo exlusão mútua, e o nome convencional para a estrutuda de dados que a fornce é mutex

A biblioteca padrão do Go que fornece exclusão mútua com sync.Mutex e seus dois métodos:
	Lock
	Unlock

Podemos definir um bloco de código a ser executado em exclusão mútua pelo que o rodeia com uma chamada para Lock e Unlock,
como mostrado no método Inc.

Nós também podemos usar defer para garantir que o mutex será desbloqueado como no método Value.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

//SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

//Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	//Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

//Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	//Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))

}
