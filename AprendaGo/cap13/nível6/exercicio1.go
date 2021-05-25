package main

import "fmt"

func main() {

	fmt.Println(retornaint())
	fmt.Println(retornaambos())
}

func retornaint() int {
	return 10
}
func retornaambos() (int, string) {
	return 20, "vinte"
}
