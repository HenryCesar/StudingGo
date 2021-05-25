// Switch sem condição.
// Switch sem uma condição é o mesmo que [switch true].
// Essa estrutura pode ser uma maneira limpa para escrever longas cadeias if-then-else.

package main

import(
	"fmt"
	"time"
)

func main() {

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}