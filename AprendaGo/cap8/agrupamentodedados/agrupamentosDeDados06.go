// Slice: make

package main

import "fmt"

func main() {

	slice := make([]int, 5, 10)

	slice[0], slice[1], slice[2], slice[3] = 1, 2, 3, 4

	for i := 0; i < 5; i++ {
		slice = append(slice, 10)
		fmt.Println("Adicionando +10")
	}
	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println("Adicionando +1 valor e dobrando o cap")

	slice = append(slice, 10)
	fmt.Println(slice, len(slice), cap(slice))

}
