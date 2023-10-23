package ast

import (
	"encoding/json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// VariableDeclaration represents a variable declaration node in the abstract syntax tree.
type VariableDeclaration struct {
	*ASTBuilder

	Id           int64           `json:"id"`                      // Unique identifier of the variable declaration node.
	NodeType     ast_pb.NodeType `json:"node_type"`               // Type of the node.
	Src          SrcNode         `json:"src"`                     // Source location information.
	Assignments  []int64         `json:"assignments"`             // List of assignment identifiers.
	Declarations []*Declaration  `json:"declarations"`            // List of declaration nodes.
	InitialValue Node[NodeType]  `json:"initial_value,omitempty"` // Initial value node.
}

// NewVariableDeclarationStatement creates a new instance of VariableDeclaration with the provided ASTBuilder.
func NewVariableDeclarationStatement(b *ASTBuilder) *VariableDeclaration {
	return &VariableDeclaration{
		ASTBuilder:   b,
		Id:           b.GetNextID(),
		NodeType:     ast_pb.NodeType_VARIABLE_DECLARATION,
		Assignments:  make([]int64, 0),
		Declarations: make([]*Declaration, 0),
	}
}

// SetReferenceDescriptor sets the reference descriptors of the VariableDeclaration node.
func (v *VariableDeclaration) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	parentNodeId := v.GetSrc().GetParentIndex()

	if parentNodeId > 0 {
		if parentNode := v.GetTree().GetById(parentNodeId); parentNode != nil {
			if parentNode.GetTypeDescription() == nil {
				parentNode.SetReferenceDescriptor(refId, refDesc)
			}
		}
	}

	return true
}

// GetId returns the unique identifier of the variable declaration node.
func (v *VariableDeclaration) GetId() int64 {
	return v.Id
}

// GetType returns the type of the node.
func (v *VariableDeclaration) GetType() ast_pb.NodeType {
	return v.NodeType
}

// GetSrc returns the source location information of the variable declaration node.
func (v *VariableDeclaration) GetSrc() SrcNode {
	return v.Src
}

// GetAssignments returns a list of assignment identifiers associated with the variable declaration.
func (v *VariableDeclaration) GetAssignments() []int64 {
	return v.Assignments
}

// GetDeclarations returns a list of declaration nodes associated with the variable declaration.
func (v *VariableDeclaration) GetDeclarations() []*Declaration {
	return v.Declarations
}

// GetInitialValue returns the initial value node associated with the variable declaration.
func (v *VariableDeclaration) GetInitialValue() Node[NodeType] {
	return v.InitialValue
}

// GetTypeDescription returns the type description associated with the variable declaration.
func (v *VariableDeclaration) GetTypeDescription() *TypeDescription {
	if len(v.Declarations) > 0 {
		return v.Declarations[0].GetTypeDescription()
	}

	return &TypeDescription{}
}

// GetNodes returns a list of nodes associated with the variable declaration (initial value and declarations).
func (v *VariableDeclaration) GetNodes() []Node[NodeType] {
	toReturn := []Node[NodeType]{}

	if v.GetInitialValue() != nil {
		toReturn = append(toReturn, v.GetInitialValue())
	}

	for _, declaration := range v.GetDeclarations() {
		toReturn = append(toReturn, declaration)
	}

	return toReturn
}

func (v *VariableDeclaration) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if id, ok := tempMap["id"]; ok {
		if err := json.Unmarshal(id, &v.Id); err != nil {
			return err
		}
	}

	if nodeType, ok := tempMap["node_type"]; ok {
		if err := json.Unmarshal(nodeType, &v.NodeType); err != nil {
			return err
		}
	}

	if src, ok := tempMap["src"]; ok {
		if err := json.Unmarshal(src, &v.Src); err != nil {
			return err
		}
	}

	if assignments, ok := tempMap["assignments"]; ok {
		if err := json.Unmarshal(assignments, &v.Assignments); err != nil {
			return err
		}
	}

	if declarations, ok := tempMap["declarations"]; ok {
		if err := json.Unmarshal(declarations, &v.Declarations); err != nil {
			return err
		}
	}

	if expression, ok := tempMap["initial_value"]; ok {
		if err := json.Unmarshal(expression, &v.InitialValue); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(expression, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["node_type"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(expression, tempNodeType)
			if err != nil {
				return err
			}
			v.InitialValue = node
		}
	}

	return nil
}

// ToProto converts the VariableDeclaration node to its corresponding protobuf representation.
func (v *VariableDeclaration) ToProto() NodeType {
	proto := ast_pb.Variable{
		Id:           v.Id,
		NodeType:     v.NodeType,
		Src:          v.Src.ToProto(),
		Assignments:  v.Assignments,
		Declarations: make([]*ast_pb.Declaration, 0),
	}

	if v.GetInitialValue() != nil {
		proto.InitialValue = v.GetInitialValue().ToProto().(*v3.TypedStruct)
	}

	for _, declaration := range v.GetDeclarations() {
		proto.Declarations = append(
			proto.Declarations,
			declaration.ToProto().(*ast_pb.Declaration),
		)
	}

	return NewTypedStruct(&proto, "Variable")
}

// Parse parses the variable declaration statement context and populates the VariableDeclaration fields.
func (v *VariableDeclaration) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.VariableDeclarationStatementContext,
) {
	v.Src = SrcNode{
		Id:          v.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.GetId(),
	}

	if ctx.VariableDeclaration() != nil {
		declaration := NewDeclaration(v.ASTBuilder)
		declaration.ParseVariableDeclaration(unit, contractNode, fnNode, bodyNode, v, ctx.VariableDeclaration())
		v.Declarations = append(v.Declarations, declaration)
		v.Assignments = append(v.Assignments, declaration.GetId())
	}

	if ctx.VariableDeclarationTuple() != nil {
		for _, declarationCtx := range ctx.VariableDeclarationTuple().AllVariableDeclaration() {
			declaration := NewDeclaration(v.ASTBuilder)
			declaration.ParseVariableDeclaration(unit, contractNode, fnNode, bodyNode, v, declarationCtx)
			v.Declarations = append(v.Declarations, declaration)
			v.Assignments = append(v.Assignments, declaration.GetId())
		}
	}

	if ctx.Expression() != nil {
		expression := NewExpression(v.ASTBuilder)
		v.InitialValue = expression.Parse(unit, contractNode, fnNode, bodyNode, v, nil, v.GetId(), ctx.Expression())
	}

	v.currentVariables = append(v.currentVariables, v)
}
