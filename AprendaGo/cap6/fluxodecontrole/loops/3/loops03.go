//A declaração [for]
package main

import "fmt"

func main() {
	// x := 0

	// for x < 10 {
	// 	x++
	// 	fmt.Printf("[%v] é menor que 10\n", x)
	// }

	x := 0

	for {
		if x < 10 {
			fmt.Printf("[%v] é menor que 10\n", x)
			x++
		} else {
			fmt.Printf("[%v] é maior que 10!\n", x)
			break
		}
	}

}
