package main

import "fmt"

var (
	anoDeNascimento = 2001
	anoAtual        = 2021
)

func main() {
	i := anoDeNascimento
	for {
		if i < anoAtual {
			i++
			fmt.Println(i)
		} else {
			break
		}
	}
}
