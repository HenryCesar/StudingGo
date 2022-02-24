package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Meu OS:", runtime.GOOS)
	fmt.Println("Meu ARCH:", runtime.GOARCH)
}
