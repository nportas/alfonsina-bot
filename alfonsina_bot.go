package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nportas/alfonsina-bot/api"
	"github.com/nportas/alfonsina-bot/palabras"
	"github.com/nportas/alfonsina-bot/poemas"
)

func main() {

	libro := palabras.NewLibroDeFrases()

	file, err := os.Open("alfonsina.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		frase := scanner.Text()

		if len(strings.TrimSpace(frase)) > 0 {
			libro.AgregarFrase(frase)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("------------Entrenamiento finalizado------------------")

	selector := palabras.NewSelectorPonderado()
	predictor := palabras.NewPredictor(libro, selector)
	poemario := poemas.NewPoemario(predictor)
	server := api.NewGinServer(poemario)
	server.Start()
}
