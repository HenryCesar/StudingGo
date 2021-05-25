package main

import "fmt"

func main() {

	y := "Meu amor"

	func(s string) {
		fmt.Println(s, "existe?", false)
	}(y)
}
