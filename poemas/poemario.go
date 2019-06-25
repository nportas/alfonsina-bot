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
	cantidadMaximaDeEstrofas                  = 1
	cantidadMinimaDeVersosParaMasDeUnaEstrofa = 4
	cantidadMaximaDeVersosParaMasDeUnaEstrofa = 7
	cantidadMinimaDeVersosParaSoloUnaEstrofa  = 3
	cantidadMaximaDeVersosParaSoloUnaEstrofa  = 6
	minPalabras                               = 3
	maxPalabras                               = 6
)

var palabrasNoFinales = []string{"con", "las", "los"}

// NewPoemario construye un nuevo poemario a partir de un libro con el que se lo entrena
func NewPoemario(generadorDeFrases palabras.GeneradorDeFrases) *Poemario {
	return &Poemario{generadorDeFrases}
}

// GenerarPoesia genera una poesia que comienza con una palabra obtenida al azar
func (p *Poemario) GenerarPoesia() *Poema {
	palabraInicial := p.generador.ObtenerPalabraAlAzar()
	return p.GenerarPoesiaAPartirDe(palabraInicial)
}

// GenerarPoesiaAPartirDe genera una poesia que comienza con la palabra primeraPalabra
func (p *Poemario) GenerarPoesiaAPartirDe(primeraPalabra string) *Poema {

	rand.Seed(time.Now().UnixNano())
	cantidadDeEstrofas := 1

	poema := new(Poema)

	for i := 0; i < cantidadDeEstrofas; i++ {
		estrofa := new(Estrofa)
		cantidadDeVersos := p.cantidadDeVersos(cantidadDeEstrofas)
		for j := 0; j < cantidadDeVersos; j++ {
			verso := p.generarVersoAPartirDe(primeraPalabra, j)
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

func (p *Poemario) generarVersoAPartirDe(primeraPalabra string, numeroDeVerso int) *Verso {
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

		palabrasDelVerso = p.eliminarPalabrasRepetidas(palabrasDelVerso, primeraPalabra, numeroDeVerso)
		palabrasDelVerso = p.eliminarUltimaPalabraSiEsMuyCorta(palabrasDelVerso)

		for _, p := range palabrasDelVerso {
			verso.Palabras = append(verso.Palabras, p)
		}
	}

	return verso
}

func (p *Poemario) eliminarPalabrasRepetidas(palabrasDelVerso []string, primeraPalabra string, i int) []string {
	// Permite repetir palabras solo en los versos pares
	if (palabrasDelVerso)[0] == primeraPalabra && i%2 != 0 {
		palabrasDelVerso = palabrasDelVerso[1:]
	}

	return palabrasDelVerso
}

func (p *Poemario) eliminarUltimaPalabraSiEsMuyCorta(palabrasDelVerso []string) []string {
	// Solo permite como última palabra, palabras de 4 o más letras
	if len(palabrasDelVerso[len(palabrasDelVerso)-1]) < 4 {
		palabrasDelVerso = palabrasDelVerso[0 : len(palabrasDelVerso)-1]
	}
	return palabrasDelVerso
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

	palabrasDelVerso := verso.Palabras

	if len(palabrasDelVerso) > 1 {
		return palabrasDelVerso[rand.Intn(len(verso.Palabras)-1)]
	}

	return palabrasDelVerso[0]
}

func (v *Verso) EsVacio() bool {
	return len(v.Palabras) == 0
}
