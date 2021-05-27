/* Os canais podem ser bufferizados.
Fornecendo o tamanho do buffer como o segundo argumento para make para inicializar um canal bufferizado:

ch := make(chan int, 100)

Envia para um bloco de canais bufferizados apenas quando o buffer está cheio.
Recebe bloco quando o buffer está vazio.*/

package main

import "fmt"

func main() {
	ch := make(chan int, 12)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
