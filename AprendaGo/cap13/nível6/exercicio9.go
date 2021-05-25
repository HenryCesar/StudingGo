package main

import "fmt"

func soma(x ...int) int {
	var res int
	for _, v := range x {
		res += v
	}
	return res
}

func impares(f func(...int) int, itens ...int) int {
	var slice []int

	for _, v := range itens {
		if v%3 == 0 {
			slice = append(slice, v)
		}
	}

	res := f(slice...)
	return res
}

func main() {

	x := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	y := impares(soma, x...)

	fmt.Println("A soma dos nros de [x] da:", soma(x...), "a soma dos impares de [x] da:", y)
}
