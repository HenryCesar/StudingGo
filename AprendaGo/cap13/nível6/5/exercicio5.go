package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}
type circulo struct {
	raio float64
}

type info interface {
	area()
}

func (q quadrado) area() {
	res := q.lado * q.lado
	fmt.Println("A área do quadrado é:", res)
}

func (c circulo) area() {
	res := 2 * math.Pi * c.raio
	fmt.Println("A área do circulo é:", res)
}

func chamarArea(i info) {
	i.area()
}

func main() {

	x := quadrado{
		lado: 15.0,
	}

	y := circulo{
		raio: 5.25,
	}

	chamarArea(x)
	chamarArea(y)

}
