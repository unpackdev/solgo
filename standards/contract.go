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

// ConfidenceCheck performs a general confidence check of the contract standard against a provided contract matcher.
// It evaluates the entire contract standard to determine how closely it matches the criteria specified in the contract matcher.
// This method returns a Discovery struct that details the overall matching confidence and a boolean indicating if a match was found.
func (e *Contract) ConfidenceCheck(contract *ContractMatcher) (Discovery, bool) {
	return ConfidenceCheck(e, contract)
}

// FunctionConfidenceCheck performs a confidence check on a specific function within the contract standard against a provided
// function matcher. It assesses whether the function in question matches the criteria defined in the function matcher,
// returning a FunctionDiscovery struct that details the matching confidence and a boolean indicating if a match was found.
func (e *Contract) FunctionConfidenceCheck(fn *Function) (FunctionDiscovery, bool) {
	return FunctionConfidenceCheck(e, fn)
}

// TokenCount returns the number of tokens associated with the standard.
func (e *Contract) TokenCount() int {
	return TokenCount(e.Standard)
}

// FunctionTokenCount calculates the total number of tokens present in a given function of the contract standard.
// It searches for the function by its name within the contract's associated functions. If found, it calculates the token count
// based on the function's inputs, outputs, and their respective types. The function name is also considered as an initial token.
// If the function name is not found within the contract's functions, it returns 0.
func (e *Contract) FunctionTokenCount(fnName string) int {
	for _, fn := range e.Standard.Functions {
		if fn.Name == fnName {
			return FunctionTokenCount(fn)
		}
	}
	return 0
}

// ToProto returns a protobuf representation of the standard.
func (e *Contract) ToProto() *eip_pb.ContractStandard {
	return e.Standard.ToProto()
}

// GetABI returns the ABI of the standard.
func (e *Contract) GetABI() string {
	return e.Standard.ABI
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
