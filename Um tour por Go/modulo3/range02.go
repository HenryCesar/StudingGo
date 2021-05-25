package main

import "fmt"

func main() {
	pow := make([]int, 10) //Alocando matriz de comprimento 10
	for i := range pow {   //Criando uma função de potência
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow { //Rodando laço 10 vezes
		fmt.Printf("%d\n", value)
	}
}

/* Você pode ignorar o índice ou valor, atribuindo o _.

Se você só quiser o índice, deixe inteiramente sem o ", value". */
