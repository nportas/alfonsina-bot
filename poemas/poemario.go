package poemas

import (
	"math/rand"
	"strings"
	"time"

	"github.com/nportas/alfonsina-bot/palabras"
)

// Poemario representación de un poemario
type Poemario struct {
	generador palabras.GeneradorDeFrases
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
func NewPoemario(generadorDeFrases palabras.GeneradorDeFrases) *Poemario {
	return &Poemario{generadorDeFrases}
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
			estrofa = estrofa + verso
			versoSinEnter := strings.Replace(verso, "/n", "", -1)
			palabrasDelVerso := strings.Split(versoSinEnter, " ")
			if len(palabrasDelVerso) > 0 {
				primeraPalabra = palabrasDelVerso[cantidadDePalabras-1]
			}
		}
		poema = poema + estrofa + "/n"
	}

	return poema
}

func (p *Poemario) generarVersoAPartirDe(primeraPalabra string) (string, int) {
	cantidadDePalabras := rand.Intn(maxPalabras-minPalabras) + minPalabras
	verso := p.generador.GenerarFraseAPartirDe(primeraPalabra, cantidadDePalabras)
	i := 0

	for len(strings.TrimSpace(verso)) == 0 && i < 1000 {
		verso = p.generador.GenerarFrase(cantidadDePalabras)
		i++
	}

	return verso + "/n", cantidadDePalabras
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
