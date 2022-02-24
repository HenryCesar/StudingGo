package main

import "fmt"

func main() {
	s := "Hello, vscode"
	fmt.Printf("%v\t[%T]\n", s, s)

	//Slice de byte
	sb := []byte(s)
	fmt.Printf("%v\t[%T]\n", sb, sb)

	//String literal uso de ``
	sl := `Stringdoidonauhul`
	fmt.Printf("%v\t[%T]\n", sl, sl)

	fmt.Println(" ")

	//Uso de %#U, %#x por caractere
	fmt.Println("Caractere por caractere")
	for _, v := range sb {
		fmt.Printf("%v - [%T]\t%#U - %#x\n", v, v, v, v)
	}

	fmt.Println(" ")

	s2 := "ascii éøâ"
	//Uso de %#U, %#x por byte
	fmt.Println("Byte por Byte")
	for i := 0; i < len(s2); i++ {
		fmt.Printf("%v - [%T]\t%#U - %#x\n", s2[i], s2[i], s2[i], s2[i])
	}
}
