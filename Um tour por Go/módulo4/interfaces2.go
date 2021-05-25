/* Um tipo implementa uma interface através da implementação dos métodos.
Não há declaração explícita de intenções, não há palavra-chave "implements".

Interfaces implícitas dissociam-se da definição de uma interface de sua implementação,
que pode então aparecer em qualquer pacote sem pré-arranjamento.  */

package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

//Esse método declara que o tipo T implementa a interface I,
//mas nós não precisamos de uma declaração explícita para fazer isso.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}
