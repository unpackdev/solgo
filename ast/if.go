package ast

import (
	"github.com/goccy/go-json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// IfStatement represents an if statement node in the abstract syntax tree.
type IfStatement struct {
	*ASTBuilder

	Id        int64           `json:"id"`        // Unique identifier of the if statement node.
	NodeType  ast_pb.NodeType `json:"nodeType"`  // Type of the node.
	Src       SrcNode         `json:"src"`       // Source location information.
	Condition Node[NodeType]  `json:"condition"` // Condition node.
	Body      Node[NodeType]  `json:"body"`      // Body node.
}

// NewIfStatement creates a new instance of IfStatement with the provided ASTBuilder.
func NewIfStatement(b *ASTBuilder) *IfStatement {
	return &IfStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_IF_STATEMENT,
	}
}

// SetReferenceDescriptor sets the reference descriptors of the IfStatement node.
func (i *IfStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the unique identifier of the if statement node.
func (i *IfStatement) GetId() int64 {
	return i.Id
}

// GetType returns the type of the node.
func (i *IfStatement) GetType() ast_pb.NodeType {
	return i.NodeType
}

// GetSrc returns the source location information of the if statement node.
func (i *IfStatement) GetSrc() SrcNode {
	return i.Src
}

// GetCondition returns the condition node of the if statement.
func (i *IfStatement) GetCondition() Node[NodeType] {
	return i.Condition
}

// GetTypeDescription returns the type description of the if statement.
func (i *IfStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeString:     "if",
		TypeIdentifier: "$_t_if",
	}
}

// GetNodes returns a list of nodes associated with the if statement (condition and body).
func (i *IfStatement) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{i.Condition, i.Body}
}

// GetBody returns the body node of the if statement.
func (i *IfStatement) GetBody() Node[NodeType] {
	return i.Body
}

// UnmarshalJSON unmarshals the JSON data into a IfStatement.
func (i *IfStatement) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if id, ok := tempMap["id"]; ok {
		if err := json.Unmarshal(id, &i.Id); err != nil {
			return err
		}
	}

	if nodeType, ok := tempMap["nodeType"]; ok {
		if err := json.Unmarshal(nodeType, &i.NodeType); err != nil {
			return err
		}
	}

	if src, ok := tempMap["src"]; ok {
		if err := json.Unmarshal(src, &i.Src); err != nil {
			return err
		}
	}

	if condition, ok := tempMap["condition"]; ok {
		if err := json.Unmarshal(condition, &i.Condition); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(condition, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["nodeType"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(condition, tempNodeType)
			if err != nil {
				return err
			}
			i.Condition = node
		}
	}

	if body, ok := tempMap["body"]; ok {
		if err := json.Unmarshal(body, &i.Body); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(body, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["nodeType"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(body, tempNodeType)
			if err != nil {
				return err
			}
			i.Body = node
		}
	}

	return nil
}

// ToProto converts the IfStatement node to its corresponding protobuf representation.
func (i *IfStatement) ToProto() NodeType {
	proto := ast_pb.If{
		Id:        i.GetId(),
		NodeType:  i.GetType(),
		Src:       i.GetSrc().ToProto(),
		Condition: i.GetCondition().ToProto().(*v3.TypedStruct),
	}

	if i.GetBody() != nil {
		proto.Body = i.GetBody().ToProto().(*ast_pb.Body)
	}

	return NewTypedStruct(&proto, "If")
}

// Parse parses the if statement context and populates the IfStatement fields.
func (i *IfStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.IfStatementContext,
) Node[NodeType] {
	i.Src = SrcNode{
		Line:        int64(ctx.GetStart().GetLine()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.Id,
	}

	expression := NewExpression(i.ASTBuilder)

	i.Condition = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, i, i.GetId(), ctx.Expression())

	body := NewBodyNode(i.ASTBuilder, false)
	if len(ctx.AllStatement()) > 0 {
		for _, statementCtx := range ctx.AllStatement() {
			if statementCtx.IsEmpty() {
				continue
			}

			if statementCtx.Block() != nil {
				body.ParseBlock(unit, contractNode, fnNode, statementCtx.Block())
				break
			}

			// There can be single-statement conditional...
			body.parseStatements(unit, contractNode, fnNode, statementCtx.GetChild(0))

			i.Body = body
		}

		i.Body = body
	} else {
		i.Body = body
	}

	return i
}
