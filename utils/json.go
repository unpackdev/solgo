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
