package predictor

// Predictor encapsulates the entire process of vulnerability prediction.
type Predictor struct {
}

// NewPredictor initializes and returns a new Predictor.
func NewPredictor(modelPath string) (*Predictor, error) {
	return &Predictor{}, nil
}
