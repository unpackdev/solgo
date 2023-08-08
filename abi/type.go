package abi

import (
	"strings"

	"github.com/txpull/solgo/ast"
	"github.com/txpull/solgo/ir"
)

type Type struct {
	Name         string
	Type         string
	InternalType string
	Outputs      []Type
}

type TypeResolver struct {
	parser *ir.Builder
}

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
			for _, stateVar := range contract.GetStateVariables() {
				if stateVar.GetName() == typeName {
					panic("State var...")
				}
			}

			for _, enumVar := range contract.GetEnums() {
				if enumVar.GetName() == typeName {
					panic("Enum var...")
				}
			}

			for _, errorVar := range contract.GetErrors() {
				if errorVar.GetName() == typeName {
					panic("Error var...")
				}
			}

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
