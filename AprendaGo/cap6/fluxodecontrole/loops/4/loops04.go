// Break & Continue
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			// faz isso quando o número não é par
			continue
		}
		fmt.Println(i)
	}
}
