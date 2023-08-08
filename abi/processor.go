package abi

import (
	"github.com/txpull/solgo/ast"
	"github.com/txpull/solgo/ir"
)

type Contract []*Method

func (b *Builder) processContract(contract *ir.Contract) *Contract {
	toReturn := Contract{}

	// Let's process state variables.
	for _, stateVar := range contract.GetStateVariables() {
		method := b.processStateVariable(stateVar)
		toReturn = append(toReturn, method)
	}

	// Let's process events.
	for _, event := range contract.GetEvents() {
		method := b.processEvent(event)
		toReturn = append(toReturn, method)
	}

	// Let's process errors.
	for _, errorNode := range contract.GetErrors() {
		method := b.processError(errorNode)
		toReturn = append(toReturn, method)
	}

	// Let's process constructor.
	if contract.GetConstructor() != nil {
		toReturn = append(
			toReturn,
			b.processConstructor(contract.GetConstructor()),
		)
	}

	// Let's process functions.
	for _, function := range contract.GetFunctions() {
		method := b.processFunction(function)
		toReturn = append(toReturn, method)
	}

	if contract.GetFallback() != nil {
		toReturn = append(
			toReturn,
			b.processFallback(contract.GetFallback()),
		)
	}

	if contract.GetReceive() != nil {
		toReturn = append(
			toReturn,
			b.processReceive(contract.GetReceive()),
		)
	}

	/*

		// Let's process structs.
		for _, structVar := range contract.GetStructs() {
			panic(structVar)
		}

		// Let's process enums.
		for _, enumVar := range contract.GetEnums() {
			panic(enumVar)
		}

	*/

	return &toReturn
}

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
