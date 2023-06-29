package solgo

// ContractInfo contains information about a contract
type ContractInfo struct {
	Comments   []string `json:"comments"`   // Comments associated with the contract
	License    string   `json:"license"`    // License information of the contract
	Pragmas    []string `json:"pragmas"`    // Pragmas specified in the contract
	Imports    []string `json:"imports"`    // Imported dependencies of the contract
	Name       string   `json:"name"`       // Name of the contract
	Implements []string `json:"implements"` // Interfaces implemented by the contract
}
