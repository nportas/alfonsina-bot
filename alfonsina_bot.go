package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	_ "github.com/gin-gonic/gin"
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

		if frase != "" {
			libro.AgregarFrase(frase)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("------------Entrenamiento finalizado------------------")

	poemario := poemas.NewPoemario(libro)
	server := api.NewGinServer(poemario)
	server.Start()
}
