// [If] [Else if] [Else]
package main

import "fmt"

func main() {

	x := 10

	if x > 100 {
		fmt.Println("[X] é maior que 100.")
	} else if x < 100 {
		fmt.Println("[X] é menor que 100.")
	} else {
		fmt.Println("[X] não é maior e nem menor que 100.")
	}

}
