package solgo

// ContractInfo contains information about a contract
type ContractInfo struct {
	Comments   []string // Comments associated with the contract
	License    string   // License information of the contract
	Pragmas    []string // Pragmas specified in the contract
	Imports    []string // Imported dependencies of the contract
	Name       string   // Name of the contract
	Implements []string // Interfaces implemented by the contract
}
