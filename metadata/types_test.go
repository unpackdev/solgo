package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContractMetadata(t *testing.T) {
	meta := &ContractMetadata{
		Raw:     "raw",
		Version: 1,
		Compiler: struct {
			Version   string `json:"version"`
			Keccak256 string `json:"keccak256"`
		}{
			Version:   "0.8.0",
			Keccak256: "keccak256",
		},
		Language: "Solidity",
	}

	t.Run("Test ToProto", func(t *testing.T) {
		proto := meta.ToProto()
		assert.NotNil(t, proto)
		assert.Equal(t, meta.Raw, proto.Raw)
		assert.Equal(t, int32(meta.Version), proto.Version)
		assert.Equal(t, meta.Compiler.Version, proto.Compiler.Version)
		assert.Equal(t, meta.Compiler.Keccak256, proto.Compiler.Keccak256)
	})

	t.Run("Test AbiToJSON", func(t *testing.T) {
		abiJson, err := meta.AbiToJSON()
		assert.NoError(t, err)
		assert.NotNil(t, abiJson)
	})

	t.Run("Test ToJSON", func(t *testing.T) {
		jsonBytes, err := meta.ToJSON()
		assert.NoError(t, err)
		assert.NotNil(t, jsonBytes)
	})
}
