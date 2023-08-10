package solc

// Version struct represents a solc versions installed on the current operating system.
type Version struct {
	Release string // The release version, for example: 0.5.0.
	Current bool   // Whether this version is the current version in use.
}
