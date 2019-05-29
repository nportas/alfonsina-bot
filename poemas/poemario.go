package poemas

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/nportas/alfonsina-bot/palabras"
)

// Poemario representación de un poemario
type Poemario struct {
	predictor *palabras.Predictor
}

const (
	cantidadMinimaDeEstrofas                  = 1
	cantidadMaximaDeEstrofas                  = 8
	cantidadMinimaDeVersosParaMasDeUnaEstrofa = 2
	cantidadMaximaDeVersosParaMasDeUnaEstrofa = 4
	cantidadMinimaDeVersosParaSoloUnaEstrofa  = 5
	cantidadMaximaDeVersosParaSoloUnaEstrofa  = 20
	minPalabras                               = 2
	maxPalabras                               = 10
)

// NewPoemario construye un nuevo poemario a partir de un libro con el que se lo entrena
func NewPoemario(libroDeFrases *palabras.LibroDeFrases) *Poemario {
	predictor := palabras.NewPredictor(libroDeFrases, palabras.NewSelectorPonderado())
	return &Poemario{predictor}
}

// GenerarPoesiaAPartirDe genera una poesia que comienza con la palabra primeraPalabra
func (p *Poemario) GenerarPoesiaAPartirDe(primeraPalabra string) string {

	rand.Seed(time.Now().UnixNano())
	cantidadDeEstrofas := rand.Intn(cantidadMaximaDeEstrofas-cantidadMinimaDeEstrofas) + cantidadMinimaDeEstrofas

	var poema string

	for i := 1; i < cantidadDeEstrofas; i++ {
		var estrofa string
		cantidadDeVersos := p.cantidadDeVersos(cantidadDeEstrofas)
		for i := 1; i < cantidadDeVersos; i++ {
			verso, cantidadDePalabras := p.generarVersoAPartirDe(primeraPalabra)
			fmt.Println("----------verso: " + verso)
			if len(strings.TrimSpace(verso)) == 0 {
				verso = p.generarVerso(cantidadDePalabras)
				fmt.Println("----------entro a generar nuevo verso: " + verso)
			}
			estrofa = estrofa + verso
			palabras := strings.Split(verso, " ")
			if len(palabras) > 0 {
				primeraPalabra = palabras[cantidadDePalabras-1]
			}
		}
		poema = poema + estrofa + "/n"
	}

	return poema
}

func (p *Poemario) generarVersoAPartirDe(primeraPalabra string) (string, int) {
	cantidadDePalabras := rand.Intn(maxPalabras-minPalabras) + minPalabras
	return p.predictor.GenerarFraseAPartirDe(primeraPalabra, cantidadDePalabras) + "/n", cantidadDePalabras
}

func (p *Poemario) generarVerso(cantidadDePalabras int) string {
	return p.predictor.GenerarFrase(cantidadDePalabras) + "/n"
}

func (p *Poemario) cantidadDeVersos(cantidadDeEstrofas int) int {

	var minVersos, maxVersos int

	if cantidadDeEstrofas == 1 {
		minVersos = cantidadMinimaDeVersosParaSoloUnaEstrofa
		maxVersos = cantidadMaximaDeVersosParaSoloUnaEstrofa
	} else {
		minVersos = cantidadMinimaDeVersosParaMasDeUnaEstrofa
		maxVersos = cantidadMaximaDeVersosParaMasDeUnaEstrofa
	}

	return rand.Intn(maxVersos-minVersos) + minVersos
}
