package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func main() {
	catFile, err := os.Open("C:/Users/henrique.almeida/Pictures/minha_foto.png")
	if err != nil {
		log.Fatal(err)
	}
	defer catFile.Close()

	imData, imType, err := image.Decode(catFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(imData)
	fmt.Println(imType)

	cat, err := png.Decode(catFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cat)
}
