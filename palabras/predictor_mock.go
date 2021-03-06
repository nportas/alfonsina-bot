package palabras

// PredictorMock Mockea un predictor
type PredictorMock struct {
	fraseAResponder         string
	palabraAlAzarAResponder string
}

// NewPredictorMock construye un nuevo predictor mock
func NewPredictorMock() *PredictorMock {
	return &PredictorMock{}
}

// GenerarFrase genera una frase mock
func (p *PredictorMock) GenerarFrase(cantidadDePalabras int) string {
	return p.fraseAResponder
}

// GenerarFraseAPartirDe genera una frase mock
func (p *PredictorMock) GenerarFraseAPartirDe(primeraPalabra string, cantidadDePalabras int) string {
	return p.fraseAResponder
}

// ResponderConFrase permite modificar la frase con la que se va a responder
func (p *PredictorMock) ResponderConFrase(frase string) {
	p.fraseAResponder = frase
}

func (p *PredictorMock) ObtenerPalabraAlAzar() string {

	return p.palabraAlAzarAResponder
}
