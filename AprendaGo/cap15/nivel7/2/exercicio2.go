package main

import (
	"fmt"
)

type pessoa struct {
	altura int
}

func (p *pessoa) mudeMe() {
	p.altura++
}

func main() {
	alt := &pessoa{180}

	fmt.Println("Altura atual:", alt.altura)
	fmt.Println("Teste:", alt)

	alt.mudeMe()
	fmt.Println("Mudando altura:", alt.altura)

}
