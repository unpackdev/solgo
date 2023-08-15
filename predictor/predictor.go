package predictor

import (
	"github.com/txpull/solgo/predictor/models"
	"github.com/txpull/solgo/predictor/preprocessing"
	"github.com/txpull/solgo/predictor/vulnerabilities"
)

// Predictor encapsulates the entire process of vulnerability prediction.
type Predictor struct {
	Model            models.Model
	Tokenizer        *preprocessing.Tokenizer
	FeatureExtractor *preprocessing.FeatureExtractor
	Detector         *vulnerabilities.Detector
	Recommender      *vulnerabilities.Recommender
}

// NewPredictor initializes and returns a new Predictor.
func NewPredictor(modelPath string) (*Predictor, error) {
	model, err := models.LoadModel(modelPath)
	if err != nil {
		return nil, err
	}

	return &Predictor{
		Model:            model,
		Tokenizer:        preprocessing.NewTokenizer(),
		FeatureExtractor: preprocessing.NewFeatureExtractor(),
		Detector:         vulnerabilities.NewDetector(model),
		Recommender:      vulnerabilities.NewRecommender(),
	}, nil
}

// PredictVulnerabilities detects vulnerabilities in the provided Solidity source code.
func (p *Predictor) PredictVulnerabilities(code string) ([]interface{}, []interface{}, error) {
	tokens, err := p.Tokenizer.Tokenize(code)
	if err != nil {
		return nil, nil, err
	}

	features, err := p.FeatureExtractor.ExtractFeatures(tokens)
	if err != nil {
		return nil, nil, err
	}

	vulnerabilities, err := p.Detector.DetectVulnerabilities(features)
	if err != nil {
		return nil, nil, err
	}

	recommendations, err := p.Recommender.ProvideRecommendations(vulnerabilities)
	if err != nil {
		return nil, nil, err
	}

	return vulnerabilities, recommendations, nil
}
