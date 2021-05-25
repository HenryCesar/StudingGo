package main

import (
	"fmt"
)

func main() {

	key := map[string][]string{
		"almeida_henrique": {"jogar", "ler", "estudar"},
		"torres_caique":    {"estudar", "pesquisar", "ouvir musica"},
		"algarra_michell":  {"encher o saco", "jogar", "encher mais o saco ainda"},
	}

	key["escagion_erick"] = []string{"sumir", "dormir", "rir"}
	delete(key, "algarra_michell")

	for i, v := range key {
		fmt.Println(i)
		for x, y := range v {
			fmt.Println("\t", x, " - ", y)
		}
	}

}
