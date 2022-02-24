package main

import (
	"fmt"
)

func main() {

	array1 := [5]int{1, 2, 3, 4, 5}

	for key, valor := range array1 {
		fmt.Println(key, valor)
	}

	fmt.Printf("%T", array1)
}
