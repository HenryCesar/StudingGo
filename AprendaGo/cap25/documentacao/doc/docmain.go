package escrevendo

import "fmt"

// X é um número random.
const X = 10

// Doc não faz nada.
func Doc() {
	fmt.Println("Essa função não faz nada.")
}

// Doc2 também não faz nada.
func Doc2() {
	fmt.Println("Essa outra função também não faz nada.")
}

// Doc3 mostra o valor de X, inutíl também.
func Doc3() {
	fmt.Println("Essa função não faz nada, minto! Ela da o valor de X:", X)
}
