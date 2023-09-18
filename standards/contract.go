package standards

import eip_pb "github.com/unpackdev/protos/dist/go/eip"

// Contract represents the contract standard.
type Contract struct {
	// Standard holds the details of the contract standard.
	Standard ContractStandard
}

// GetName returns the name of the standard.
func (e *Contract) GetName() string {
	return e.Standard.Name
}

// GetType returns the type of the standard.
func (e *Contract) GetType() Standard {
	return e.Standard.Type
}

// GetUrl returns the URL of the standard.
func (e *Contract) GetUrl() string {
	return e.Standard.Url
}

// GetFunctions returns the functions associated with the standard.
func (e *Contract) GetFunctions() []Function {
	return e.Standard.Functions
}

// GetEvents returns the events associated with the standard.
func (e *Contract) GetEvents() []Event {
	return e.Standard.Events
}

// GetStandard returns the complete contract standard.
func (e *Contract) GetStandard() ContractStandard {
	return e.Standard
}

// IsStagnant returns a boolean indicating whether the standard is stagnant.
func (e *Contract) IsStagnant() bool {
	return e.Standard.Stagnant
}

func (e *Contract) ConfidenceCheck(contract *ContractMatcher) (Discovery, bool) {
	return ConfidenceCheck(e, contract)
}

// TokenCount returns the number of tokens associated with the standard.
func (e *Contract) TokenCount() int {
	return TokenCount(e.Standard)
}

// ToProto returns a protobuf representation of the standard.
func (e *Contract) ToProto() *eip_pb.ContractStandard {
	return e.Standard.ToProto()
}

// String returns the name of the standard.
func (e *Contract) String() string {
	return e.GetName()
}

// NewContract initializes and returns an instance of the standard.
// It sets up the standard with its name, type, associated functions, and events.
func NewContract(standard ContractStandard) EIP {
	return &Contract{Standard: standard}
}

// GetContractByStandard returns the contract standard by its type.
func GetContractByStandard(standard Standard) (EIP, error) {
	if standard, ok := standards[standard]; ok {
		return NewContract(standard), nil
	}
	return nil, ErrStandardNotFound
}
