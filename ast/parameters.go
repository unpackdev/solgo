package ast

import (
	"sync/atomic"

	"github.com/txpull/solgo/parser"
)

type ParametersList struct {
	ID         int64       `json:"id"`
	NodeType   string      `json:"node_type"`
	Parameters []Parameter `json:"parameters"`
	Src        Src         `json:"src"`
}

type Parameter struct {
	Constant         bool              `json:"constant"`
	ID               int64             `json:"id"`
	Mutability       string            `json:"mutability"`
	Name             string            `json:"name"`
	NodeType         string            `json:"node_type"`
	Scope            int64             `json:"scope"`
	Src              Src               `json:"src"`
	StateVariable    bool              `json:"state_variable"`
	StorageLocation  string            `json:"storage_location"`
	TypeDescriptions *TypeDescriptions `json:"type_descriptions"`
	TypeName         *TypeName         `json:"type_name"`
	Visibility       string            `json:"visibility"`
}

type FunctionReturnParameters struct {
	ID         int64         `json:"id"`
	NodeType   string        `json:"node_type"`
	Parameters []interface{} `json:"parameters"`
	Src        Src           `json:"src"`
}

type TypeName struct {
	ID               int64             `json:"id"`
	Name             string            `json:"name"`
	NodeType         string            `json:"node_type"`
	Src              Src               `json:"src"`
	TypeDescriptions *TypeDescriptions `json:"type_descriptions"`
}

type TypeDescriptions struct {
	TypeIdentifier string `json:"type_identifier"`
	TypeString     string `json:"type_string"`
}

func (b *ASTBuilder) traverseParameterList(node Node, parameterCtx parser.IParameterListContext) *ParametersList {
	if parameterCtx.IsEmpty() {
		return nil
	}

	id := atomic.AddInt64(&b.nextID, 1) - 1

	parametersList := &ParametersList{
		ID:         id,
		Parameters: make([]Parameter, 0),
		Src: Src{
			Line:   parameterCtx.GetStart().GetLine(),
			Start:  parameterCtx.GetStart().GetStart(),
			End:    parameterCtx.GetStop().GetStop(),
			Length: parameterCtx.GetStop().GetStop() - parameterCtx.GetStart().GetStart() + 1,
			Index:  node.Src.Index,
		},
		NodeType: NodeTypeParameterList.String(),
	}

	for _, paramCtx := range parameterCtx.AllParameterDeclaration() {
		if paramCtx.IsEmpty() {
			continue
		}

		parameterID := atomic.AddInt64(&b.nextID, 1) - 1
		pNode := Parameter{
			ID: parameterID,
			Src: Src{
				Line:   paramCtx.GetStart().GetLine(),
				Start:  paramCtx.GetStart().GetStart(),
				End:    paramCtx.GetStop().GetStop(),
				Length: paramCtx.GetStop().GetStop() - paramCtx.GetStart().GetStart() + 1,
				Index:  parametersList.ID,
			},
			Scope: node.ID,
			Name: func() string {
				if paramCtx.Identifier() != nil {
					return paramCtx.Identifier().GetText()
				}
				return ""
			}(),
			NodeType: NodeTypeVariableDeclaration.String(),
			// Just hardcoding it here to internal as I am not sure how
			// could it be possible to be anything else.
			// @TODO: Check if it is possible to be anything else.
			Visibility: NodeTypeVisibilityInternal.String(),
			// Mutable is the default state for parameter declarations.
			Mutability: NodeTypeMutabilityMutable.String(),
		}

		if paramCtx.GetType_().ElementaryTypeName() != nil {
			typeNameID := atomic.AddInt64(&b.nextID, 1) - 1
			pNode.NodeType = NodeTypeElementaryTypeName.String()
			typeCtx := paramCtx.GetType_().ElementaryTypeName()
			normalizedTypeName, normalizedTypeIdentifier := normalizeTypeDescription(
				typeCtx.GetText(),
			)
			pNode.TypeName = &TypeName{
				ID:   typeNameID,
				Name: typeCtx.GetText(),
				Src: Src{
					Line:   paramCtx.GetStart().GetLine(),
					Start:  paramCtx.GetStart().GetStart(),
					End:    paramCtx.GetStop().GetStop(),
					Length: paramCtx.GetStop().GetStop() - paramCtx.GetStart().GetStart() + 1,
					Index:  pNode.ID,
				},
				NodeType: NodeTypeElementaryTypeName.String(),
				TypeDescriptions: &TypeDescriptions{
					TypeIdentifier: normalizedTypeIdentifier,
					TypeString:     normalizedTypeName,
				},
			}
		}

		parametersList.Parameters = append(parametersList.Parameters, pNode)
	}

	return parametersList
}
