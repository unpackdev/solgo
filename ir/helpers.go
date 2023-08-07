package ir

import (
	"encoding/json"
	"fmt"
	"os"
)

// nolint:unused
func (b *Builder) dumpNode(whatever interface{}) {
	j, _ := json.MarshalIndent(whatever, "", "\t")
	fmt.Println(string(j))
	os.Exit(1)
}

// nolint:unused
func (b *Builder) dumpNodeNoExit(whatever interface{}) {
	j, _ := json.MarshalIndent(whatever, "", "\t")
	fmt.Println(string(j))
}
