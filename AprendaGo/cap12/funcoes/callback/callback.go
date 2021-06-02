package main

import "fmt"

func main() {
	total := somenteImpares(soma, []int{3, 15, 16, 8, 6, 123, 76, 9}...)
	fmt.Println(total)
}

func soma(x ...int) int {
	res := 0

	for _, valor := range x {
		res += valor
	}
	return res
}

func somenteImpares(f func(x ...int) int, y ...int) int {
	var slice []int

	for _, v := range y {
		if v%3 == 0 {
			slice = append(slice, v)
		}
	}
	res := f(slice...)
	return res
}
