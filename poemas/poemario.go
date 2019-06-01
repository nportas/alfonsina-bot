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

type Poema struct {
	Estrofas []*Estrofa
}

type Estrofa struct {
	Versos []*Verso
}

type Verso struct {
	Palabras []string
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

var palabrasNoFinales = []string{"con", "las", "los"}

// NewPoemario construye un nuevo poemario a partir de un libro con el que se lo entrena
func NewPoemario(generadorDeFrases palabras.GeneradorDeFrases) *Poemario {
	return &Poemario{generadorDeFrases}
}

// GenerarPoesiaAPartirDe genera una poesia que comienza con la palabra primeraPalabra
func (p *Poemario) GenerarPoesiaAPartirDe(primeraPalabra string) *Poema {

	rand.Seed(time.Now().UnixNano())
	cantidadDeEstrofas := rand.Intn(cantidadMaximaDeEstrofas-cantidadMinimaDeEstrofas) + cantidadMinimaDeEstrofas

	poema := new(Poema)

	for i := 1; i < cantidadDeEstrofas; i++ {
		estrofa := new(Estrofa)
		cantidadDeVersos := p.cantidadDeVersos(cantidadDeEstrofas)
		for i := 1; i < cantidadDeVersos; i++ {
			verso := p.generarVersoAPartirDe(primeraPalabra)
			if !verso.EsVacio() {
				estrofa.Versos = append(estrofa.Versos, verso)
				primeraPalabra = p.obtenerNuevaPrimeraPalabra(verso)
			} else {
				break
			}
		}
		if len(estrofa.Versos) > 0 {
			poema.Estrofas = append(poema.Estrofas, estrofa)
		}
	}

	return poema
}

func (p *Poemario) generarVersoAPartirDe(primeraPalabra string) *Verso {
	cantidadDePalabras := rand.Intn(maxPalabras-minPalabras) + minPalabras
	frase := p.generador.GenerarFraseAPartirDe(primeraPalabra, cantidadDePalabras)
	i := 0

	// Si no pudo generar nada intento generar la frase a partir de cualquier palabra
	// Hago solo 1000 intentos
	for len(strings.TrimSpace(frase)) == 0 && i < 1000 {
		frase = p.generador.GenerarFrase(cantidadDePalabras)
		i++
	}

	verso := new(Verso)

	// Paso las palabras del verso a un slice
	if len(strings.TrimSpace(frase)) > 0 {
		palabrasDelVerso := strings.Split(frase, " ")
		for i, p := range palabrasDelVerso {
			if i != len(palabrasDelVerso)-1 {
				verso.Palabras = append(verso.Palabras, p)
			} else if len(p) > 3 {
				verso.Palabras = append(verso.Palabras, p)
			}
		}
	}

	return verso
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

func (p *Poemario) obtenerNuevaPrimeraPalabra(verso *Verso) string {

	var primeraPalabra string
	palabrasDelVerso := verso.Palabras

	if len(palabrasDelVerso) > 0 {
		primeraPalabra = palabrasDelVerso[len(palabrasDelVerso)-1]
	}

	return primeraPalabra
}

func (v *Verso) EsVacio() bool {
	return len(v.Palabras) == 0
}
