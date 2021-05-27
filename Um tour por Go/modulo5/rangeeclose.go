/* Um remetente pode close um canal para indicar que os valores não serão mais enviados.
Receptores poem testar se um canal foi fechado através da tribuição de um segundo parâmetro para a expressão de recepção:
v, ok := <-ch

ok é false se não há mais valores a recebe e o canal está fechado.
O laço for i:= range c recebe valores do canal repetidamente até que seja fechado.

Nota: Apenas o remetente deve fechar um canal, nunca o receptor. O envio em um canal fechado irá causar um pânico.
Outra nota: Canais não são como arquivos, você geralmente não precisa fechá-los. O encerramento só é necessário quando
o receptor precisa saber que não há mais valores chegando, como para terminar um laço range.
*/
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
