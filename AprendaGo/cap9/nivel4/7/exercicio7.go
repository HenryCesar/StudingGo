package main

import "fmt"

func main() {
	slice := [][]string{
		{
			"Henrique", "Almeida", "Jogar",
		},
		{
			"Caique", "Torres", "Estudar",
		},
		{
			"Michell", "Algarra", "Encher o saco",
		},
	}

	for k, v := range slice {
		fmt.Println(k, v)
	}

	for _, v := range slice {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}

	}
}
