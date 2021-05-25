//Switch é uma forma mais curta de escrever uma sequência de declarações
// [if] - [else]
// Diferente de C, C++, Java, JavaScript e PHP, o Go apenas executa o caso selecionado.

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os { //runtime.GOOS é uma constante utilizada para detectar o SO utilizado.
	case "darwin":
	 	fmt.Println("OS X.") //Declaração [break] é implícita pela linguagem
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}
}
