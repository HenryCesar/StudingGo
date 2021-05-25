package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice da slice para ter o tamanho 0
	s = s[:0]
	printSlice(s)

	// Aumentando o tamanho
	s = s[:4]
	printSlice(s)

	// Abandonando os dois primeiros valores
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) //len = comprimento (lenght) e cap = capacidade (capacity)
}

/* Uma slice tem tanto um `tamanho` quanto uma `capacidade`.

O comprimento de uma slice é o número de elementos que ela contém.

A capacidade de uma slice é o número de elemento na matriz subjacente, contando a partir do primeiro elemento na slice.

O comprimento e acapcidade de uma slice [S] podem ser obtidos utilizando as expressões "len(s)" e "cap(s)".

Você pode estender o comprimento de uma slice "refatiando-a", desde que tenha capacidade suficiente. Tente alterar uma das operações
da slice no programa acima para esntendê-la além de sua capacidade e veja o que acontece. */
