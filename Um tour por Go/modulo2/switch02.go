/*
	Aqui temos exemplos de como executar um switch em conjunto a biblioteca Time
	Switch cases avaliam casos de cima par abaixo, parando quando um caso for bem-sucedido.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	t := time.Now()
	switch time.Saturday { //Utilizando uma condição específica
	case today + 0:
		fmt.Println("Today.")
	case today + 1 :
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	fmt.Println("\nToday is ")
	switch today {// Utilizando uma condição "aberta".
	case time.Saturday, time.Sunday: 
		fmt.Println("Weekend.")
	default:
		fmt.Println("Weekday")
	}

	fmt.Println("\nWhat's a clock?")
	switch { //Sem condição
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	case t.Hour() > 12:
		fmt.Println("It's after noon")
	}
}

// As condições devem ser interpretadas em conjunto ao contexto!!!