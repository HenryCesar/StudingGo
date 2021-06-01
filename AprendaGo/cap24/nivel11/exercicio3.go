package main

import "fmt"

type erroEspecial struct {
	variavel string
}

func (e erroEspecial) Error() string {
	return fmt.Sprintf("ERROR: %v\n", e.variavel)
}

func usandoErro(e erroEspecial) {
	fmt.Println("Valor de erroEspecial como parâmetro:", e, "Agora como variavel:", e.variavel)
}

func main() {
	atributoError := erroEspecial{"SHINELIKEADIAMOND!"}

	usandoErro(atributoError)
}
