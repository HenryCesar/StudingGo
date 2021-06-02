package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

type ordenarPorConsumo []carro

func (x ordenarPorConsumo) Len() int {
	return len(x)
}

func (x ordenarPorConsumo) Less(i, j int) bool {
	return x[i].consumo > x[j].consumo
}

func (x ordenarPorConsumo) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type ordenarPorLucroProDonoDoPosto []carro

func (x ordenarPorLucroProDonoDoPosto) Len() int {
	return len(x)
}

func (x ordenarPorLucroProDonoDoPosto) Less(i, j int) bool {
	return x[i].consumo < x[j].consumo
}

func (x ordenarPorLucroProDonoDoPosto) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type ordenarPorPotencia []carro

func (x ordenarPorPotencia) Len() int {
	return len(x)
}

func (x ordenarPorPotencia) Less(i, j int) bool {
	return x[i].potencia < x[j].potencia
}

func (x ordenarPorPotencia) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func main() {

	carros := []carro{
		{"Chevete", 50, 8},
		{"Porsche", 300, 5},
		{"Fusca", 20, 30},
	}

	fmt.Println("Não ordenado:", carros)

	sort.Sort(ordenarPorPotencia(carros))
	fmt.Println("Ordenando por potencia:", carros)

	sort.Sort(ordenarPorConsumo(carros))
	fmt.Println("Ordenando por consumo:", carros)

	sort.Sort(ordenarPorLucroProDonoDoPosto(carros))
	fmt.Println("Ordenando por [Lucro pro Dono do Posto]:", carros)
}
