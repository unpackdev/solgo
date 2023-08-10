package utils

import (
	"encoding/json"
)

// ToJSON converts a Go struct to its JSON representation.
func ToJSON(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

// ToProtoJSON converts a Go struct to its JSON representation.
func ToProtoJSON(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

// ToJSONPretty returns a pretty-printed JSON representation of the provided interface.
// This function is primarily used for debugging purposes.
func ToJSONPretty(data interface{}) ([]byte, error) {
	return json.MarshalIndent(data, "", "\t")
}
