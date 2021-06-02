//Switch
package main

import "fmt"

func main() {

	x := 0

	switch {
	case x < 5:
		fmt.Println("[X] < 5")
	case x == 5:
		fmt.Println("[X] = 5")
	case x > 5:
		fmt.Println("[X] > 5")
	default:
		fmt.Println("ERRO: Não foi alocado conteúdo à variável [X]")
	}

}
