package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13} // Matriz literal de int para [s]

	s = s[1:4] //Slice subscrita a matriz de [s] nas posições 1 a 3. (3 5 7)
	fmt.Println(s)

	s = s[:2] //Slice subscrita a slice de [s] omitindo o limite baixo e solicitando o item de indice 1. (3 5)
	fmt.Println(s)

	s = s[1:] //Slice subscrita a slice de [s] solicitando item de indice 1 e omitindo o limite alto. (3)
	fmt.Println(s)
}

/* Ao "fatiar", você pode omitir as posições de limite altas ou baixas para usar seus padrões em seu lugar.

O padrão é igual a zero para o limite baixo e o comprimento da slice para o limite alto.

Para a matriz:
	var a [10]int
Essas expressões das slices são equivalentes:
	a[0:10]
	a[:10]
	a[0:]
	a[:] */
