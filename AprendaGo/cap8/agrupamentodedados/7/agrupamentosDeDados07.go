//Slice multi-dimensional

package main

import (
	"fmt"
)

func main() {
	ss := [][]int{
		//Índice: 0  1  2  4  5  6	  		// Índice:
		{1, 2, 3, 4, 5, 6},       // 0
		{7, 8, 9, 10, 11, 12},    // 1
		{13, 14, 15, 16, 17, 18}, // 2
	}

	fmt.Println(ss[2][4])

	ss2 := [][][][]int{
		{
			[][]int{
				{1, 2, 3, 4, 5, 6},
			},
			[][]int{
				{10, 20, 30, 40, 50, 60},
			},
		},
		{
			[][]int{
				{7, 8, 9, 10, 11, 12},
			},
			[][]int{
				{70, 80, 90, 100, 110, 120},
			},
		},
	}
	fmt.Println("slice multi-dimensional gigante", +ss2[1][0][0][5])
}
