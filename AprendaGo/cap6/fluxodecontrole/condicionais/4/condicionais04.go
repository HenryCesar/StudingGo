// Switch pt.2

package main

import "fmt"

var (
	x, y interface{}
)

func main() {

	x = true

	switch x.(type) {
	case int:
		fmt.Println("é um int")
	case bool:
		fmt.Println("é um bool")
	case string:
		fmt.Println("é um string")
	case float64:
		fmt.Println("é um float64")
	}

	switch y := 2; {
	case y == 1:
		fmt.Println("número:", y)
	case y == 2:
		fmt.Println("número:", y)
	case y == 3:
		fmt.Println("número:", y)
	case y == 4:
		fmt.Println("número:", y)
	}
}
