package ast

import (
	"github.com/goccy/go-json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// ForStatement represents a for loop statement in the AST.
type ForStatement struct {
	*ASTBuilder

	Id          int64           `json:"id"`          // Unique identifier for the ForStatement node.
	NodeType    ast_pb.NodeType `json:"nodeType"`    // Type of the AST node.
	Src         SrcNode         `json:"src"`         // Source location information.
	Initialiser Node[NodeType]  `json:"initialiser"` // Initialiser expression.
	Condition   Node[NodeType]  `json:"condition"`   // Condition expression.
	Closure     Node[NodeType]  `json:"closure"`     // Closure expression.
	Body        *BodyNode       `json:"body"`        // Body of the for loop.
}

// NewForStatement creates a new ForStatement node with a given ASTBuilder.
func NewForStatement(b *ASTBuilder) *ForStatement {
	return &ForStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_FOR_STATEMENT,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the ForStatement node.
// We don't need to do any reference description updates here, at least for now...
func (f *ForStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the ID of the ForStatement node.
func (f *ForStatement) GetId() int64 {
	return f.Id
}

// GetType returns the NodeType of the ForStatement node.
func (f *ForStatement) GetType() ast_pb.NodeType {
	return f.NodeType
}

// GetSrc returns the SrcNode of the ForStatement node.
func (f *ForStatement) GetSrc() SrcNode {
	return f.Src
}

// GetInitialiser returns the initialiser expression.
func (f *ForStatement) GetInitialiser() Node[NodeType] {
	return f.Initialiser
}

// GetCondition returns the condition expression.
func (f *ForStatement) GetCondition() Node[NodeType] {
	return f.Condition
}

// GetClosure returns the closure expression.
func (f *ForStatement) GetClosure() Node[NodeType] {
	return f.Closure
}

// GetBody returns the body of the for loop.
func (f *ForStatement) GetBody() *BodyNode {
	return f.Body
}

// GetNodes returns the child nodes of the ForStatement node.
func (f *ForStatement) GetNodes() []Node[NodeType] {
	toReturn := []Node[NodeType]{f.Initialiser, f.Condition, f.Closure}
	toReturn = append(toReturn, f.Body.GetNodes()...)
	return toReturn
}

// GetTypeDescription returns the TypeDescription of the ForStatement node.
func (f *ForStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeString:     "for",
		TypeIdentifier: "$_t_for",
	}
}

// MarshalJSON marshals the ForStatement node into a JSON byte slice.
func (f *ForStatement) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if id, ok := tempMap["id"]; ok {
		if err := json.Unmarshal(id, &f.Id); err != nil {
			return err
		}
	}

	if nodeType, ok := tempMap["nodeType"]; ok {
		if err := json.Unmarshal(nodeType, &f.NodeType); err != nil {
			return err
		}
	}

	if src, ok := tempMap["src"]; ok {
		if err := json.Unmarshal(src, &f.Src); err != nil {
			return err
		}
	}

	if initialiser, ok := tempMap["initialiser"]; ok {
		if err := json.Unmarshal(initialiser, &f.Initialiser); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(initialiser, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["nodeType"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(initialiser, tempNodeType)
			if err != nil {
				return err
			}
			f.Initialiser = node
		}
	}

	if closure, ok := tempMap["closure"]; ok {
		if err := json.Unmarshal(closure, &f.Closure); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(closure, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["nodeType"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(closure, tempNodeType)
			if err != nil {
				return err
			}
			f.Closure = node
		}
	}

	if condition, ok := tempMap["condition"]; ok {
		if err := json.Unmarshal(condition, &f.Condition); err != nil {
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
			f.Condition = node
		}
	}

	if body, ok := tempMap["body"]; ok {
		if err := json.Unmarshal(body, &f.Body); err != nil {
			return err
		}
	}

	return nil
}

// ToProto returns a protobuf representation of the ForStatement node.
func (f *ForStatement) ToProto() NodeType {
	proto := ast_pb.For{
		Id:       f.GetId(),
		NodeType: f.GetType(),
		Src:      f.GetSrc().ToProto(),
	}

	if f.GetInitialiser() != nil {
		proto.Initialiser = f.GetInitialiser().ToProto().(*v3.TypedStruct)
	}

	if f.GetCondition() != nil {
		proto.Condition = f.GetCondition().ToProto().(*v3.TypedStruct)
	}

	if f.GetClosure() != nil {
		proto.Closure = f.GetClosure().ToProto().(*v3.TypedStruct)
	}

	if f.GetBody() != nil {
		proto.Body = f.GetBody().ToProto().(*ast_pb.Body)
	}

	return NewTypedStruct(&proto, "For")
}

// Parse parses a for loop statement context into the ForStatement node.
// Documentation: https://docs.soliditylang.org/en/v0.8.19/grammar.html#a4.SolidityParser.forStatement
func (f *ForStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.ForStatementContext,
) Node[NodeType] {
	f.Src = SrcNode{
		Line:        int64(ctx.GetStart().GetLine()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.Id,
	}

	if ctx.SimpleStatement() != nil {
		statement := NewSimpleStatement(f.ASTBuilder)
		f.Initialiser = statement.Parse(
			unit, contractNode, fnNode, bodyNode, f, f.GetId(), ctx.SimpleStatement().(*parser.SimpleStatementContext),
		)
	}

	if ctx.ExpressionStatement() != nil {
		f.Condition = parseExpressionStatement(
			f.ASTBuilder,
			unit, contractNode, fnNode,
			bodyNode, f, f.GetId(), ctx.ExpressionStatement().(*parser.ExpressionStatementContext),
		)
	}

	if ctx.Expression() != nil {
		expression := NewExpression(f.ASTBuilder)
		f.Closure = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, f, f.GetId(), ctx.Expression())
	}

	if ctx.Statement() != nil && ctx.Statement().Block() != nil && !ctx.Statement().Block().IsEmpty() {
		bodyNode := NewBodyNode(f.ASTBuilder, false)
		bodyNode.ParseBlock(unit, contractNode, f, ctx.Statement().Block())
		f.Body = bodyNode

		if ctx.Statement().Block() != nil && ctx.Statement().Block().AllUncheckedBlock() != nil {
			for _, uncheckedCtx := range ctx.Statement().Block().AllUncheckedBlock() {
				bodyNode := NewBodyNode(f.ASTBuilder, false)
				bodyNode.ParseUncheckedBlock(unit, contractNode, f, uncheckedCtx)
				f.Body.Statements = append(f.Body.Statements, bodyNode)
			}
		}
	} else {
		bodyNode := NewBodyNode(f.ASTBuilder, false)
		f.Body = bodyNode
	}

	return f
}
