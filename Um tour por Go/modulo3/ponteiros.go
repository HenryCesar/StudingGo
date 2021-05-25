// Go tem ponteiros. Um ponteiro guarda na memória o endereço de uma variável.

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i //[i] é o valor subjacente de [p]
	fmt.Println(*p)
	*p = 21 //[*p] (ponteiro p) está definindo o valor subjacente dele = [i]
	fmt.Println(i)

	p = &j       //[j] é o valor subjacente de [p] (valor que está dentro de p)
	*p = *p / 37 //o valor subjacente do ponteiro [p]
	fmt.Println(j)

}

/* Ao definir o valor de [i] a partir de "*p", você está efetuando um "desreferenciamento" ou
"indirecionamento". */
