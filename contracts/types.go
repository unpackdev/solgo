package contracts

// ContractInfo contains information about a contract
type ContractInfo struct {
	Comments        []string `json:"comments"`         // Comments associated with the contract
	License         string   `json:"license"`          // License information of the contract
	Pragmas         []string `json:"pragmas"`          // Pragmas specified in the contract
	Imports         []string `json:"imports"`          // Imported dependencies of the contract
	Name            string   `json:"name"`             // Name of the contract
	Implements      []string `json:"implements"`       // Interfaces implemented by the contract
	IsProxy         bool     `json:"is_proxy"`         // Whether the contract is a proxy
	ProxyConfidence int16    `json:"proxy_confidence"` // Confidence in the proxy detection
	IsContract      bool     `json:"is_contract"`      // Whether the contract is a contract
	IsInterface     bool     `json:"is_interface"`     // Whether the contract is an interface
	IsLibrary       bool     `json:"is_library"`       // Whether the contract is a library
	IsAbstract      bool     `json:"is_abstract"`      // Whether the contract is abstract
}
