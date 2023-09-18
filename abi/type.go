package abi

import (
	"strings"

	"github.com/unpackdev/solgo/ast"
	"github.com/unpackdev/solgo/ir"
)

// Type represents a type within the Ethereum ABI.
type Type struct {
	Name         string
	Type         string
	InternalType string
	Outputs      []Type
}

// TypeResolver provides methods to resolve and discover types within the ABI.
type TypeResolver struct {
	parser *ir.Builder
}

// ResolveType determines the type of a given typeName based on its identifier.
// It returns a string representation of the type.
func (t *TypeResolver) ResolveType(typeName *ast.TypeDescription) string {
	if strings.HasPrefix(typeName.GetIdentifier(), "t_mapping") {
		return "mapping"
	}

	if strings.HasPrefix(typeName.GetIdentifier(), "t_struct") {
		return "struct"
	}

	if strings.HasPrefix(typeName.GetIdentifier(), "t_contract") {
		return "contract"
	}

	if strings.HasPrefix(typeName.GetIdentifier(), "t_enum") {
		return "enum"
	}

	if strings.HasPrefix(typeName.GetIdentifier(), "t_error") {
		return "error"
	}

	return normalizeTypeName(typeName.GetString())
}

// ResolveMappingType resolves the input and output types for a given mapping type.
// It returns slices of MethodIO for inputs and outputs respectively.
func (t *TypeResolver) ResolveMappingType(typeName *ast.TypeDescription) ([]MethodIO, []MethodIO) {
	inputs := make([]MethodIO, 0)
	outputs := make([]MethodIO, 0)
	matched, inputList, outputList := parseMappingType(typeName.GetString())

	if matched {
		for _, input := range inputList {
			dType := t.discoverType(input)
			inputs = append(inputs, MethodIO{
				Name:         dType.Name,
				Type:         dType.Type,
				InternalType: dType.InternalType,
			})
		}

		for _, output := range outputList {
			dType := t.discoverType(output)
			if len(dType.Outputs) > 0 {
				for _, out := range dType.Outputs {
					outputs = append(outputs, MethodIO{
						Name:         out.Name,
						Type:         out.Type,
						InternalType: out.InternalType,
					})
				}
			} else {
				outputs = append(outputs, MethodIO{
					Name:         dType.Name,
					Type:         dType.Type,
					InternalType: dType.InternalType,
				})
			}
		}
	}

	return inputs, outputs
}

// ResolveStructType resolves the type of a given struct and returns its MethodIO representation.
func (t *TypeResolver) ResolveStructType(typeName *ast.TypeDescription) MethodIO {
	nameCleaned := strings.Replace(typeName.GetString(), "struct ", "", -1)
	nameCleaned = strings.TrimLeft(nameCleaned, "[]")
	nameParts := strings.Split(nameCleaned, ".")

	toReturn := MethodIO{
		Name:       nameParts[1],
		Components: make([]MethodIO, 0),
		Type:       "tuple",
	}

	for _, contract := range t.parser.GetRoot().GetContracts() {
		for _, structVar := range contract.GetStructs() {
			if structVar.GetName() == toReturn.Name {
				for _, member := range structVar.GetMembers() {
					dType := t.discoverType(member.GetTypeDescription().GetString())
					if len(dType.Outputs) > 0 {
						for _, out := range dType.Outputs {
							toReturn.Components = append(toReturn.Components, MethodIO{
								Name:         member.GetName(),
								Type:         out.Type,
								InternalType: member.GetTypeDescription().GetString(),
							})
						}
					} else {
						toReturn.Components = append(toReturn.Components, MethodIO{
							Name:         member.GetName(),
							Type:         dType.Type,
							InternalType: member.GetTypeDescription().GetString(),
						})
					}
				}
			}
		}
	}

	return toReturn
}

// discoverType determines the type of a given typeName.
// It searches through the contracts and their structs to find a matching type.
// Returns a Type representation of the discovered type.
// @WARN: This function will probably need more work to handle more complex types.
func (t *TypeResolver) discoverType(typeName string) Type {
	toReturn := Type{
		Outputs: make([]Type, 0),
	}

	discoveredType, found := normalizeTypeNameWithStatus(typeName)
	if found {
		toReturn.Type = discoveredType
		toReturn.InternalType = discoveredType
		return toReturn
	} else {
		for _, contract := range t.parser.GetRoot().GetContracts() {
			for _, structVar := range contract.GetStructs() {
				if structVar.GetName() == typeName {
					for _, member := range structVar.GetMembers() {
						toReturn.Outputs = append(toReturn.Outputs, Type{
							Name:         member.GetName(),
							Type:         normalizeTypeName(member.GetTypeDescription().GetString()),
							InternalType: member.GetTypeDescription().GetString(),
						})
					}
				}
			}
		}
	}

	return toReturn
}
