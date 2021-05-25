//Utilizando ASCII
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Iniciando impressão.")
	for x := 33; x <= 122; x++ {
		formatPrint(x)
	}
	fmt.Println("Done!")
}

func formatPrint(x int) {
	fmt.Printf("%c\n", x)
}

//%d\t%#x\t