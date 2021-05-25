package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42 //Inserindo um elemento no map
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48 //Atualizando um elemento no map
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}

// Inserir ou atualizar um elemento no map "m":
// 			m[key] = elem

//Recuperar um elemnto:
//			elem  = m[key]

//Excluir um elemento:
//			delete(m, key)

//Testar se uma chave está presente com dois valores:
//			elem, ok = m[key]
