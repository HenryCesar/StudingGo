// Mutação de Maps

package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["Answer"] = 42 //Inserir ou atualizar um elemento
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer") //Excluir um elemento
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] //Testar se a chave está presente com dois valores
	fmt.Println("The value:", v, "Present?", ok)
	// Se a "[key]" está em m, "ok" é "true". Se não, "ok" é "false".
	// elemento, ok = m[key]
}

//Recuperar um elemento: elem = m[key]
//Se elemento ou ok, não foram declarados você pode usar a seguinte declaração
//				"elem, ok := m[key]"
