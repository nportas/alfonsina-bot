package poemas_test

import (
	"strings"
	"testing"

	"github.com/nportas/alfonsina-bot/palabras"
	"github.com/nportas/alfonsina-bot/poemas"
)

func TestVersoVacio(t *testing.T) {

	// Inicialización
	libroDeFrases := palabras.NewLibroDeFrases()
	libroDeFrases.AgregarFrase("boca muerta que fuiste boca viva")
	poemario := poemas.NewPoemario(libroDeFrases)

	// Operación
	fraseGenerada := poemario.GenerarPoesiaAPartirDe("amor")

	// Validación
	if len(strings.TrimSpace(fraseGenerada)) == 0 {
		t.Logf("Se esperaba un verso pero se generó '%v'", fraseGenerada)
		t.FailNow()
	}

}
