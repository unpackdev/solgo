package ir

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	ir_pb "github.com/unpackdev/protos/dist/go/ir"
	"github.com/unpackdev/solgo/ast"
)

// Fallback represents a fallback function definition in the IR.
type Fallback struct {
	Unit             *ast.Fallback     `json:"ast"`
	Id               int64             `json:"id"`
	NodeType         ast_pb.NodeType   `json:"nodeType"`
	Name             string            `json:"name"`
	Kind             ast_pb.NodeType   `json:"kind"`
	Implemented      bool              `json:"implemented"`
	Visibility       ast_pb.Visibility `json:"visibility"`
	StateMutability  ast_pb.Mutability `json:"stateMutability"`
	Virtual          bool              `json:"virtual"`
	Modifiers        []*Modifier       `json:"modifiers"`
	Overrides        []*Override       `json:"overrides"`
	Parameters       []*Parameter      `json:"parameters"`
	ReturnStatements []*Parameter      `json:"return"`
}

// GetAST returns the AST (Abstract Syntax Tree) for the fallback function definition.
func (f *Fallback) GetAST() *ast.Fallback {
	return f.Unit
}

// GetId returns the ID of the fallback function definition.
func (f *Fallback) GetId() int64 {
	return f.Id
}

// GetNodeType returns the NodeType of the fallback function definition.
func (f *Fallback) GetNodeType() ast_pb.NodeType {
	return f.NodeType
}

// GetName returns the name of the fallback function definition.
func (f *Fallback) GetName() string {
	return f.Name
}

// GetKind returns the kind of the fallback function definition.
func (f *Fallback) GetKind() ast_pb.NodeType {
	return f.Kind
}

// IsImplemented returns whether the fallback function is implemented.
func (f *Fallback) IsImplemented() bool {
	return f.Implemented
}

// GetVisibility returns the visibility of the fallback function.
func (f *Fallback) GetVisibility() ast_pb.Visibility {
	return f.Visibility
}

// GetStateMutability returns the state mutability of the fallback function.
func (f *Fallback) GetStateMutability() ast_pb.Mutability {
	return f.StateMutability
}

// IsVirtual returns whether the fallback function is virtual.
func (f *Fallback) IsVirtual() bool {
	return f.Virtual
}

// GetModifiers returns the modifiers applied to the fallback function.
func (f *Fallback) GetModifiers() []*Modifier {
	return f.Modifiers
}

// GetOverrides returns the overrides applied to the fallback function.
func (f *Fallback) GetOverrides() []*Override {
	return f.Overrides
}

// GetParameters returns the parameters of the fallback function.
func (f *Fallback) GetParameters() []*Parameter {
	return f.Parameters
}

// GetReturnStatements returns the return statements of the fallback function.
func (f *Fallback) GetReturnStatements() []*Parameter {
	return f.ReturnStatements
}

// GetSrc returns the source code location of the fallback function.
func (f *Fallback) GetSrc() ast.SrcNode {
	return f.Unit.GetSrc()
}

// ToProto converts the Fallback to its protobuf representation.
func (f *Fallback) ToProto() *ir_pb.Fallback {
	proto := &ir_pb.Fallback{
		Id:              f.GetId(),
		NodeType:        f.GetNodeType(),
		Kind:            f.GetKind(),
		Name:            f.GetName(),
		Implemented:     f.IsImplemented(),
		Visibility:      f.GetVisibility(),
		StateMutability: f.GetStateMutability(),
		Virtual:         f.IsVirtual(),
		Modifiers:       make([]*ir_pb.Modifier, 0),
		Overrides:       make([]*ir_pb.Override, 0),
		Parameters:      make([]*ir_pb.Parameter, 0),
		Return:          make([]*ir_pb.Parameter, 0),
	}

	for _, modifier := range f.GetModifiers() {
		proto.Modifiers = append(proto.Modifiers, modifier.ToProto())
	}

	for _, overrides := range f.GetOverrides() {
		proto.Overrides = append(proto.Overrides, overrides.ToProto())
	}

	for _, parameter := range f.GetParameters() {
		proto.Parameters = append(proto.Parameters, parameter.ToProto())
	}

	for _, returnStatement := range f.GetReturnStatements() {
		proto.Return = append(proto.Return, returnStatement.ToProto())
	}

	return proto
}

// processFallback processes the fallback function definition unit and returns the Fallback.
func (b *Builder) processFallback(unit *ast.Fallback) *Fallback {
	toReturn := &Fallback{
		Unit:             unit,
		Id:               unit.GetId(),
		NodeType:         unit.GetType(),
		Kind:             unit.GetKind(),
		Name:             "fallback",
		Implemented:      unit.IsImplemented(),
		Visibility:       unit.GetVisibility(),
		StateMutability:  unit.GetStateMutability(),
		Virtual:          unit.IsVirtual(),
		Modifiers:        make([]*Modifier, 0),
		Overrides:        make([]*Override, 0),
		Parameters:       make([]*Parameter, 0),
		ReturnStatements: make([]*Parameter, 0),
	}

	for _, modifier := range unit.GetModifiers() {
		toReturn.Modifiers = append(toReturn.Modifiers, &Modifier{
			Unit:          modifier,
			Id:            modifier.GetId(),
			NodeType:      modifier.GetType(),
			Name:          modifier.GetName(),
			ArgumentTypes: modifier.GetArgumentTypes(),
		})
	}

	for _, oride := range unit.GetOverrides() {
		for _, overrideParameter := range oride.GetOverrides() {
			override := &Override{
				Unit:                    oride,
				Id:                      overrideParameter.GetId(),
				NodeType:                overrideParameter.GetType(),
				Name:                    overrideParameter.GetName(),
				ReferencedDeclarationId: overrideParameter.GetReferencedDeclaration(),
				TypeDescription:         overrideParameter.GetTypeDescription(),
			}
			toReturn.Overrides = append(toReturn.Overrides, override)
		}
	}

	for _, parameter := range unit.GetParameters().GetParameters() {
		param := &Parameter{
			Unit:            parameter,
			Id:              parameter.GetId(),
			NodeType:        parameter.GetType(),
			Name:            parameter.GetName(),
			Type:            parameter.GetTypeName().GetName(),
			TypeDescription: parameter.GetTypeDescription(),
		}

		if param.GetType() == "" && parameter.GetTypeName().GetPathNode() != nil {
			param.Type = parameter.GetTypeName().GetPathNode().Name
		}

		toReturn.Parameters = append(toReturn.Parameters, param)
	}

	for _, returnStatement := range unit.GetReturnParameters().GetParameters() {
		param := &Parameter{
			Unit:            returnStatement,
			Id:              returnStatement.GetId(),
			NodeType:        returnStatement.GetType(),
			Name:            returnStatement.GetName(),
			Type:            returnStatement.GetTypeName().GetName(),
			TypeDescription: returnStatement.GetTypeDescription(),
		}

		if param.GetType() == "" && returnStatement.GetTypeName().GetPathNode() != nil {
			param.Type = returnStatement.GetTypeName().GetPathNode().Name
		}

		toReturn.ReturnStatements = append(toReturn.ReturnStatements, param)
	}

	return toReturn
}
