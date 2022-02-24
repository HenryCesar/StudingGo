package main

import "fmt"

func main() {

	x := 10

	func(x int) {
		fmt.Println(x, "vezes 8735 é:")
		fmt.Println(x * 8735)
	}(x)

}
