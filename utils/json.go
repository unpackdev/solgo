package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// ToJSON converts a Go struct to its JSON representation.
func ToJSON(data any) ([]byte, error) {
	return json.Marshal(data)
}

// FromJSON converts JSON data into a Go data structure of type T.
// It takes a slice of bytes representing the JSON data and a pointer to the type T where the data will be decoded.
// Returns an error if the decoding fails.
func FromJSON[T any](data []byte, target *T) error {
	err := json.Unmarshal(data, target)
	if err != nil {
		// Handle the error according to your application's requirements
		return fmt.Errorf("error unmarshaling from JSON: %w", err)
	}
	return nil
}

// ToProtoJSON converts a Go struct to its JSON representation.
func ToProtoJSON(data any) ([]byte, error) {
	return json.Marshal(data)
}

// ToJSONPretty returns a pretty-printed JSON representation of the provided interface.
// This function is primarily used for debugging purposes.
func ToJSONPretty(data any) ([]byte, error) {
	return json.MarshalIndent(data, "", "\t")
}

// DumpNodeWithExit prints a formatted JSON representation of the provided interface and exits the program.
// This function is primarily used for debugging purposes.
func DumpNodeWithExit(whatever any) {
	j, _ := json.MarshalIndent(whatever, "", "\t")
	fmt.Println(string(j))
	os.Exit(1)
}

// DumpNodeNoExit prints a formatted JSON representation of the provided interface without exiting the program.
// This function is primarily used for debugging purposes.
func DumpNodeNoExit(whatever any) {
	j, _ := json.MarshalIndent(whatever, "", "\t")
	fmt.Println(string(j))
}
