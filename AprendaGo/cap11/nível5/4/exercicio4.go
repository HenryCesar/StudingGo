package main

import "fmt"

func main() {

	pessoa1 := struct {
		nome   string
		função string
	}{
		nome:   "Henrique",
		função: "Estagiário",
	}

	fmt.Println(pessoa1)
}
