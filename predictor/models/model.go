package models

// Model is an interface that any vulnerability detection model should implement.
type Model interface {
	// Train trains the model on the provided data.
	Train(data interface{}) error

	// Predict uses the model to predict vulnerabilities in the provided data.
	Predict(data interface{}) ([]interface{}, error)

	// Save saves the model to a specified location.
	Save(path string) error

	// Load loads the model from a specified location.
	Load(path string) error
}

func LoadModel(path string) (Model, error) {
	return nil, nil
}
