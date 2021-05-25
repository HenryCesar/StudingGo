/* Um closure é uma função valor que referencia variáveis de fora de seu corpo.
A função pode acessar e atribuir nas variáveis referenciadas; nesse sentido a função é "limitada" às variáveis.  */

package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
