package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	// Esse exemplo usa png.Decode o qual consegue apenas descifrar imagens PNG.
	catFile, err := os.Open("C:/Users/henrique.almeida/Pictures/zeee.png")
	if err != nil {
		log.Fatal(err)
	}
	defer catFile.Close()

	// Considerando o uso geral de image.Decode o qual pode pesquisar e decodificar qualquer formato de imagem.
	img, err := png.Decode(catFile)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(img)

	levels := []string{" ", "░", "▒", "▓", "█"}

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := c.Y / 51 // 51 * 5 = 255
			if level == 5 {
				level--
			}
			fmt.Print(levels[level])
		}
		fmt.Print("\n")
	}
}
