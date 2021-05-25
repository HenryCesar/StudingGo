package main

import "fmt"

type veículo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veículo
	traçãoNasQuatro bool
}

type sedan struct {
	veículo
	modeloLuxo bool
}

func main() {

	carro1 := sedan{veículo{4, "azul"}, true}
	carro2 := caminhonete{veículo{4, "preto"}, true}

	fmt.Println("O carro1 tem", carro1.portas, "portas e é da cor", carro1.cor)
	fmt.Println("O carro1 é de luxo? R:", carro1.modeloLuxo)

	fmt.Println("O carro2 tem", carro2.portas, "portas e é da cor", carro2.cor)
	fmt.Println("O carro2 tem tração nas quadro rodas? R:", carro2.traçãoNasQuatro)

}
