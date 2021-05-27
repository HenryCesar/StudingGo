package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

/* Uma das interfaces mais uíqua é a Stringer definida pelo pacote fmt

type Stringer interface{
	String() string
}

A Stringer é um tipo que pode descrever-se como uma string.
O pacote fmt (e muitos outros) olham para essa interface
para mostrar valores. */
