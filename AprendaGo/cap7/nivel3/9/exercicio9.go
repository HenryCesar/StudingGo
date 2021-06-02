package main

import "fmt"

func main() {

	esporteFavorito := "basquete"

	switch esporteFavorito {
	case "vólei":
		fmt.Println("Meu esporte favorito é vólei!")
	case "futebol":
		fmt.Println("Meu esporte favorito é futebol!")
	case "ping-pong":
		fmt.Println("Meu esporte favorito é ping-pong!")
	case "tênis":
		fmt.Println("Meu esporte favorito é tênis!")
	default:
		fmt.Println("O switch não acha meu esporte favorito! D:")
	}

}
