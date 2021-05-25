package main

import "fmt"

func main() {
	x := um()
	x()
}

func um() func() {
	return func() {
		fmt.Println("sempre tento complicar...")
	}
}
