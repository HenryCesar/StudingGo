package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ordenarPorIdade []user

func (x ordenarPorIdade) Len() int           { return len(x) }
func (x ordenarPorIdade) Less(i, j int) bool { return x[i].Age < x[j].Age }
func (x ordenarPorIdade) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ordenarPorSobrenome []user

func (x ordenarPorSobrenome) Len() int           { return len(x) }
func (x ordenarPorSobrenome) Less(i, j int) bool { return x[i].Last < x[j].Last }
func (x ordenarPorSobrenome) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println("Não ordenado:", users)
	// your code goes here
	for _, v := range users { //Strings bonitinhas
		sort.Strings(v.Sayings)
	}

	sort.Sort(ordenarPorIdade(users))
	fmt.Println("Ordenado por idade:")
	arrumar(users)

	sort.Sort(ordenarPorSobrenome(users))
	fmt.Println("Ordenado por Sobrenome")
	arrumar(users)

}

func arrumar(x []user) {
	for i, v := range x {
		fmt.Println(i, "- Idade:", v.Age, "  Nome:", v.First, " Sobrenome:", v.Last) // Printa as informações do "user"
		for _, y := range v.Sayings {
			fmt.Println("\t", y) //Printa os Sayings
		}
	}
}
