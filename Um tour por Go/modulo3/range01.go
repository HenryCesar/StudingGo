package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

/* 	O [range] do laço [for] itera sobre uma slice ou map.

	Ao variar sobre uma slice, dois valores são retornados par acada iteração. O primeiro é o índice,
o segundo uma cópia do elemento daquele índice. */
