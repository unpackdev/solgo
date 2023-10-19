package abi

import (
	abi_pb "github.com/unpackdev/protos/dist/go/abi"
	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/ir"
)

// Contract represents a collection of Ethereum contract methods.
type Contract []*Method

func (c *Contract) GetMethodByType(typeName string) *Method {
	for _, method := range *c {
		if method.Type == typeName {
			return method
		}
	}

	return nil
}

func (c *Contract) GetMethodByName(name string) *Method {
	for _, method := range *c {
		if method.Name == name {
			return method
		}
	}

	return nil
}

// ToProto converts the Contract into its protocol buffer representation.
func (c *Contract) ToProto() *abi_pb.Contract {
	toReturn := &abi_pb.Contract{
		Methods: make([]*abi_pb.Method, 0),
	}

	for _, method := range *c {
		toReturn.Methods = append(toReturn.Methods, method.ToProto())
	}

	return toReturn
}

// processContract processes an IR contract and returns a Contract representation of it.
// It extracts state variables, events, errors, constructor, functions, fallback, and receive methods.
func (b *Builder) processContract(contract *ir.Contract) *Contract {
	toReturn := Contract{}

	// Process state variables.
	for _, stateVar := range contract.GetStateVariables() {
		method := b.processStateVariable(stateVar)
		toReturn = append(toReturn, method)
	}

	// Process events.
	for _, event := range contract.GetEvents() {
		method := b.processEvent(event)
		toReturn = append(toReturn, method)
	}

	// Process errors.
	for _, errorNode := range contract.GetErrors() {
		method := b.processError(errorNode)
		toReturn = append(toReturn, method)
	}

	// Process constructor.
	if contract.GetConstructor() != nil {
		toReturn = append(
			toReturn,
			b.processConstructor(contract.GetConstructor()),
		)
	}

	// Process functions.
	for _, function := range contract.GetFunctions() {
		method := b.processFunction(function)
		toReturn = append(toReturn, method)
	}

	// Process fallback.
	if contract.GetFallback() != nil {
		toReturn = append(
			toReturn,
			b.processFallback(contract.GetFallback()),
		)
	}

	// Process receive.
	if contract.GetReceive() != nil {
		toReturn = append(
			toReturn,
			b.processReceive(contract.GetReceive()),
		)
	}

	return &toReturn
}

// buildMethodIO constructs a MethodIO object based on the provided method and type description.
// It resolves the type of the method and sets the appropriate fields.
func (b *Builder) buildMethodIO(method MethodIO, typeDescr *ast.TypeDescription) MethodIO {
	typeName := b.resolver.ResolveType(typeDescr)

	switch typeName {
	case "mapping":
		inputList, outputList := b.resolver.ResolveMappingType(typeDescr)
		method.Inputs = append(method.Inputs, inputList...)
		method.Outputs = append(method.Outputs, outputList...)
	case "contract":
		method.Type = "address"
		method.InternalType = typeDescr.GetString()
	case "enum":
		method.Type = "uint8"
		method.InternalType = typeDescr.GetString()
	case "struct":
		return b.resolver.ResolveStructType(typeDescr)
	default:
		method.Type = typeName
		method.InternalType = typeDescr.GetString()
	}

	return method
}
