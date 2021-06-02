// Maps: range & deletando

package main

import (
	"fmt"
)

func main() {
	qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		18:  "idade de ir para festa",
	}

	fmt.Println(qualquercoisa)

	total := 0

	for key, _ := range qualquercoisa {
		total += key
	}
	// for i, v := range qualquercoisa {
	// 	fmt.Println(i, v)
	// }

	fmt.Println(total)

	delete(qualquercoisa, 123)
	fmt.Println(qualquercoisa)
}
