package main

import "fmt"

func main() {
	names := [4]string{ //Matriz com 4 string
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2] //Variável recebendo uma slice da posição 0 a 1 (0:2)
	b := names[1:3] //Variável recebendo uma slice da posição 1 a 2 (1:3)
	fmt.Println(a, b)

	b[0] = "XXX" //Renomeando a string da posição 0 dentro de b (b possui 2 posições sendo elas iguais a posição 1 e 2 de names)
	fmt.Println(a, b)
	fmt.Println(names)
}

/* Uma slice não armazena todos os daods, apenas descreve uma seção de uma matriz subjacente.

Alterando os elementos de uma slice modifica os elementos correspondentes da sua matriz subjacente.

Outras slices que partilham a mesma matriz subjacente vão ver essas mudanças.  */
