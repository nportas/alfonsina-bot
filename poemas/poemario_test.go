package poemas_test

import (
	"testing"

	"github.com/nportas/alfonsina-bot/palabras"
	"github.com/nportas/alfonsina-bot/poemas"
)

func TestNoDebeGenerarPoesiaConVersosVacios(t *testing.T) {

	// Inicialización
	predictor := palabras.NewPredictorMock()
	predictor.ResponderConFrase(" ")
	poemario := poemas.NewPoemario(predictor)

	// Operación
	poesia := poemario.GenerarPoesiaAPartirDe("amor")

	// Validación
	if len(poesia.Estrofas) != 0 {
		t.Logf("No se esperaba una poesía pero se generó '%v'", poesia)
		t.FailNow()
	}

}

/*
func TestNoDebeGenerarVersoConConectoresComoPalabraFinal(t *testing.T) {

	// Inicialización
	predictor := palabras.NewPredictorMock()
	predictor.ResponderConFrase("fuego en tu sombra todo lo real me")
	poemario := poemas.NewPoemario(predictor)

	// Operación
	poesia := poemario.GenerarPoesiaAPartirDe("fuego")

	// Validación
	palabrasDelVerso := poesia.Estrofas[0].Versos[0].Separar(" ")
	if palabrasDelVerso[len(palabrasDelVerso)-1] == "me" {
		t.Logf("Se esperaba un verso sin conectores como palabra final pero se generó '%s'", poesia.Estrofas[0].Versos[0].ToString())
		t.FailNow()
	}

}*/

func TestEsVacio(t *testing.T) {

	// Inicialización
	verso := poemas.Verso(" ")

	// Operación
	esVacio := verso.EsVacio()

	// Validación
	if esVacio == false {
		t.Logf("Se esperaba true pero se obtuvo '%v'", esVacio)
		t.FailNow()
	}
}
