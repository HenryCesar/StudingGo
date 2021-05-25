package main

import "fmt"

func main() {

	nros := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(soma(nros...))

	nroslice := []int{20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	fmt.Println(somaslice(nroslice))

}

func soma(x ...int) int {
	res := 0
	for _, v := range x {
		res += v
	}

	return res
}

func somaslice(x []int) int {
	res := 0
	for _, v := range x {
		res += v
	}
	return res
}
