package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// Acrescentando em slices nulas.
	s = append(s, 0)
	printSlice(s)

	// O slice cresce como quiser
	s = append(s, 1)
	printSlice(s)

	// Podemos adicionar mais de um elemento desta vez.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("lent=%d cap=%d %v\n", len(s), cap(s), s)
}

/* 	É comum acrescentar novos elementos em uma slice, então Go dispões de uma função padrão para isso,
o [append].
		func append(s []T, vs ...T) []T
	O primeiro parâmentro [s] do [append] é uma slice do tipo T, e o resto são valores de T para
acrescentar na slice.
	O valor resultante do [append] é uma slice contendo todos os elementos da slice original mais os
valores providos.
	Se a matriz anterior de [s] for muito pequena para caber todos os valores, uma matriz gigante será alocada.
A slice retornada apontará para a nova amtriz alocada.  */
