package main

import "fmt"

type máquina struct {
	tipo int
	nome string
}

type câmera struct {
	máquina
	rara bool
}

func main() {

	canon := câmera{máquina{tipo: 1, nome: "Canon Pistola"}, true}
	tecpix := câmera{máquina{tipo: 2000, nome: "TecPix 2000"}, false}

	fmt.Println("A câmera", canon.nome, "é rara?", canon.rara, "Ela é do tipo", canon.tipo)
	fmt.Println("A câmera", tecpix.nome, "é rara?", tecpix.rara, "Ela é do tipo", tecpix.tipo)

}
