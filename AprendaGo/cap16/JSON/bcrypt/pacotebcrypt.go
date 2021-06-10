package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "21Julho1980"
	senhaerrada := "21Julho1990"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))

	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senhaerrada)))

}
