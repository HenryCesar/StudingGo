// Empilhando [defer]'s
/*
	Chamadas de funções adiadas são empurradas por uma pilha. Quando a função retorna, as chamadas de
adiantamento são executadas na ordem em que a última a entrar é a prima a sair

*/

package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
