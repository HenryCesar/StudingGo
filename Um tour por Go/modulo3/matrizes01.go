package main

import "fmt"

func main() {
	var a [2]string // O tipo [2]string é uma matriz de valor 2 do tipo string (valor sujeita a quantidade de itens)
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

/* O tamanho da matriz, "var a [2]string" (uma variável a com uma matriz de 2 strings), é parte de seu tipo,
logo, matrizes não podem ser redimensionadas. Isso parece uma limitação, mas não se preocupe,
Go dispõe de uma maneira conveniente para trabalhar com matrizes. */
