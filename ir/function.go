package ir

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	ir_pb "github.com/unpackdev/protos/dist/go/ir"
	"github.com/unpackdev/solgo/ast"
)

// Function represents a function declaration in the IR.
type Function struct {
	unit                    *ast.Function
	Id                      int64             `json:"id"`
	NodeType                ast_pb.NodeType   `json:"node_type"`
	Kind                    ast_pb.NodeType   `json:"kind"`
	Name                    string            `json:"name"`
	Implemented             bool              `json:"implemented"`
	Visibility              ast_pb.Visibility `json:"visibility"`
	StateMutability         ast_pb.Mutability `json:"state_mutability"`
	Virtual                 bool              `json:"virtual"`
	ReferencedDeclarationId int64             `json:"referenced_declaration_id"`
	Modifiers               []*Modifier       `json:"modifiers"`
	Overrides               []*Override       `json:"overrides"`
	Parameters              []*Parameter      `json:"parameters"`
	Body                    *Body             `json:"body"`
	ReturnStatements        []*Parameter      `json:"return"`
}

// GetAST returns the AST (Abstract Syntax Tree) for the function declaration.
func (f *Function) GetAST() *ast.Function {
	return f.unit
}

// GetId returns the ID of the function declaration.
func (f *Function) GetId() int64 {
	return f.Id
}

// GetNodeType returns the NodeType of the function declaration.
func (f *Function) GetNodeType() ast_pb.NodeType {
	return f.NodeType
}

// GetName returns the name of the function declaration.
func (f *Function) GetName() string {
	return f.Name
}

// GetKind returns the kind of the function declaration.
func (f *Function) GetKind() ast_pb.NodeType {
	return f.Kind
}

// IsImplemented returns whether the function is implemented or not.
func (f *Function) IsImplemented() bool {
	return f.Implemented
}

// GetVisibility returns the visibility of the function.
func (f *Function) GetVisibility() ast_pb.Visibility {
	return f.Visibility
}

// GetStateMutability returns the state mutability of the function.
func (f *Function) GetStateMutability() ast_pb.Mutability {
	return f.StateMutability
}

// IsVirtual returns whether the function is virtual or not.
func (f *Function) IsVirtual() bool {
	return f.Virtual
}

// GetModifiers returns the modifiers of the function.
func (f *Function) GetModifiers() []*Modifier {
	return f.Modifiers
}

// GetOverrides returns the overrides of the function.
func (f *Function) GetOverrides() []*Override {
	return f.Overrides
}

// GetParameters returns the parameters of the function.
func (f *Function) GetParameters() []*Parameter {
	return f.Parameters
}

// GetReturnStatements returns the return statements of the function.
func (f *Function) GetReturnStatements() []*Parameter {
	return f.ReturnStatements
}

// GetReferencedDeclarationId returns the referenced declaration id of the function.
func (f *Function) GetReferencedDeclarationId() int64 {
	return f.ReferencedDeclarationId
}

// GetBody returns the body of the function.
func (f *Function) GetBody() *Body {
	return f.Body
}

// GetSrc returns the source code of the function.
func (f *Function) GetSrc() ast.SrcNode {
	return f.unit.GetSrc()
}

// ToProto returns the protocol buffer version of the function.
func (f *Function) ToProto() *ir_pb.Function {
	proto := &ir_pb.Function{
		Id:                      f.GetId(),
		NodeType:                f.GetNodeType(),
		Kind:                    f.GetKind(),
		Name:                    f.GetName(),
		Implemented:             f.IsImplemented(),
		Visibility:              f.GetVisibility(),
		StateMutability:         f.GetStateMutability(),
		Virtual:                 f.IsVirtual(),
		ReferencedDeclarationId: f.GetReferencedDeclarationId(),
		Modifiers:               make([]*ir_pb.Modifier, 0),
		Overrides:               make([]*ir_pb.Override, 0),
		Parameters:              make([]*ir_pb.Parameter, 0),
		Body:                    f.GetBody().ToProto(),
		Return:                  make([]*ir_pb.Parameter, 0),
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

// processFunction processes the function declaration and returns the Function.
func (b *Builder) processFunction(unit *ast.Function, parseBody bool) *Function {
	toReturn := &Function{
		unit:                    unit,
		Id:                      unit.GetId(),
		NodeType:                unit.GetType(),
		Kind:                    unit.GetKind(),
		Name:                    unit.GetName(),
		Implemented:             unit.IsImplemented(),
		Visibility:              unit.GetVisibility(),
		StateMutability:         unit.GetStateMutability(),
		Virtual:                 unit.IsVirtual(),
		ReferencedDeclarationId: unit.GetReferencedDeclaration(),
		Modifiers:               make([]*Modifier, 0),
		Overrides:               make([]*Override, 0),
		Parameters:              make([]*Parameter, 0),
		ReturnStatements:        make([]*Parameter, 0),
	}

	for _, modifier := range unit.GetModifiers() {
		toReturn.Modifiers = append(toReturn.Modifiers, &Modifier{
			unit:          modifier,
			Id:            modifier.GetId(),
			NodeType:      modifier.GetType(),
			Name:          modifier.GetName(),
			ArgumentTypes: modifier.GetArgumentTypes(),
		})
	}

	for _, oride := range unit.GetOverrides() {
		for _, overrideParameter := range oride.GetOverrides() {
			override := &Override{
				unit:                    oride,
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
			unit:            parameter,
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

	if parseBody {
		toReturn.Body = b.processFunctionBody(toReturn, unit.GetBody())
	}

	for _, returnStatement := range unit.GetReturnParameters().GetParameters() {
		param := &Parameter{
			unit:            returnStatement,
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
