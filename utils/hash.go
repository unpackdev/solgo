package utils

import "golang.org/x/crypto/sha3"

// Keccak256 returns the Keccak256 hash of the input data.
func Keccak256(data []byte) []byte {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(data)
	return hasher.Sum(nil)
}
