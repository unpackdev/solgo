package ast

import (
	"encoding/json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// TryStatement represents a try-catch statement in the AST.
type TryStatement struct {
	*ASTBuilder

	Id               int64            `json:"id"`                // Unique identifier for the TryStatement node.
	NodeType         ast_pb.NodeType  `json:"node_type"`         // Type of the AST node.
	Src              SrcNode          `json:"src"`               // Source location information.
	Body             *BodyNode        `json:"body"`              // Body of the try block.
	Kind             ast_pb.NodeType  `json:"kind"`              // Kind of try statement.
	Returns          bool             `json:"returns"`           // True if the try statement returns.
	ReturnParameters *ParameterList   `json:"return_parameters"` // Return parameters of the try statement.
	Expression       Node[NodeType]   `json:"expression"`        // Expression within the try block.
	Clauses          []Node[NodeType] `json:"clauses"`           // List of catch clauses.
	Implemented      bool             `json:"implemented"`       // True if the try statement is implemented.
}

// NewTryStatement creates a new TryStatement node with a given ASTBuilder.
func NewTryStatement(b *ASTBuilder) *TryStatement {
	return &TryStatement{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_TRY_STATEMENT,
		Kind:       ast_pb.NodeType_TRY,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the TryStatement node.
func (t *TryStatement) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the ID of the TryStatement node.
func (t *TryStatement) GetId() int64 {
	return t.Id
}

// GetType returns the NodeType of the TryStatement node.
func (t *TryStatement) GetType() ast_pb.NodeType {
	return t.NodeType
}

// GetSrc returns the SrcNode of the TryStatement node.
func (t *TryStatement) GetSrc() SrcNode {
	return t.Src
}

// GetBody returns the body of the TryStatement node.
func (t *TryStatement) GetBody() *BodyNode {
	return t.Body
}

// GetKind returns the kind of the try statement.
func (t *TryStatement) GetKind() ast_pb.NodeType {
	return t.Kind
}

// GetImplemented returns true if the try statement is implemented.
func (t *TryStatement) IsImplemented() bool {
	return t.Implemented
}

// GetTypeDescription returns the TypeDescription of the TryStatement node.
func (t *TryStatement) GetTypeDescription() *TypeDescription {
	return &TypeDescription{
		TypeString:     "try",
		TypeIdentifier: "$_t_try",
	}
}

// GetNodes returns the child nodes of the TryStatement node.
func (t *TryStatement) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, t.Body)
	toReturn = append(toReturn, t.Expression)
	toReturn = append(toReturn, t.Clauses...)
	toReturn = append(toReturn, t.ReturnParameters.GetNodes()...)
	return toReturn
}

// GetExpression returns the expression within the try block.
func (t *TryStatement) GetExpression() Node[NodeType] {
	return t.Expression
}

// GetClauses returns the list of catch clauses.
func (t *TryStatement) GetClauses() []Node[NodeType] {
	return t.Clauses
}

// GetReturns returns true if the try statement returns.
func (t *TryStatement) GetReturns() bool {
	return t.Returns
}

// GetReturnParameters returns the return parameters of the try statement.
func (t *TryStatement) GetReturnParameters() *ParameterList {
	return t.ReturnParameters
}

// MarshalJSON marshals the TryStatement node into a JSON byte slice.
func (t *TryStatement) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if id, ok := tempMap["id"]; ok {
		if err := json.Unmarshal(id, &t.Id); err != nil {
			return err
		}
	}

	if nodeType, ok := tempMap["node_type"]; ok {
		if err := json.Unmarshal(nodeType, &t.NodeType); err != nil {
			return err
		}
	}

	if src, ok := tempMap["src"]; ok {
		if err := json.Unmarshal(src, &t.Src); err != nil {
			return err
		}
	}

	if body, ok := tempMap["body"]; ok {
		if err := json.Unmarshal(body, &t.Body); err != nil {
			return err
		}
	}

	if kind, ok := tempMap["kind"]; ok {
		if err := json.Unmarshal(kind, &t.Kind); err != nil {
			return err
		}
	}

	if returns, ok := tempMap["returns"]; ok {
		if err := json.Unmarshal(returns, &t.Returns); err != nil {
			return err
		}
	}

	if implemented, ok := tempMap["implemented"]; ok {
		if err := json.Unmarshal(implemented, &t.Implemented); err != nil {
			return err
		}
	}

	if returnParameters, ok := tempMap["return_parameters"]; ok {
		if err := json.Unmarshal(returnParameters, &t.ReturnParameters); err != nil {
			return err
		}
	}

	if expression, ok := tempMap["expression"]; ok {
		if err := json.Unmarshal(expression, &t.Expression); err != nil {
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
			t.Expression = node
		}
	}

	if clauses, ok := tempMap["clauses"]; ok {
		var nodes []json.RawMessage
		if err := json.Unmarshal(clauses, &nodes); err != nil {
			return err
		}

		for _, tempNode := range nodes {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(tempNode, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["node_type"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(tempNode, tempNodeType)
			if err != nil {
				return err
			}
			t.Clauses = append(t.Clauses, node)
		}
	}

	return nil
}

// ToProto returns a protobuf representation of the TryStatement node.
func (t *TryStatement) ToProto() NodeType {
	proto := ast_pb.Try{
		Id:       t.GetId(),
		NodeType: t.GetType(),
		Kind:     t.GetKind(),
		Src:      t.GetSrc().ToProto(),
		Returns:  t.GetReturns(),
	}

	if t.GetExpression() != nil {
		proto.Expression = t.GetExpression().ToProto().(*v3.TypedStruct)
	}

	if t.GetClauses() != nil {
		for _, clause := range t.GetClauses() {
			proto.Clauses = append(proto.Clauses, clause.ToProto().(*v3.TypedStruct))
		}
	}

	if t.GetBody() != nil {
		proto.Body = t.GetBody().ToProto().(*ast_pb.Body)
	}

	if t.GetReturnParameters() != nil {
		proto.ReturnParameters = t.GetReturnParameters().ToProto()
	}

	return NewTypedStruct(&proto, "Try")
}

// Parse parses a try-catch statement context into the TryStatement node.
func (t *TryStatement) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	ctx *parser.TryStatementContext,
) Node[NodeType] {
	t.Src = SrcNode{
		Id:          t.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: bodyNode.Id,
	}

	expression := NewExpression(t.ASTBuilder)
	t.Expression = expression.Parse(unit, contractNode, fnNode, bodyNode, nil, t, t.GetId(), ctx.Expression())

	if ctx.Block() != nil && !ctx.Block().IsEmpty() {
		bodyNode := NewBodyNode(t.ASTBuilder, false)
		bodyNode.ParseBlock(unit, contractNode, t, ctx.Block())
		t.Body = bodyNode

		if ctx.Block().AllUncheckedBlock() != nil {
			for _, uncheckedCtx := range ctx.Block().AllUncheckedBlock() {
				bodyNode := NewBodyNode(t.ASTBuilder, false)
				bodyNode.ParseUncheckedBlock(unit, contractNode, t, uncheckedCtx)
				t.Body.Statements = append(t.Body.Statements, bodyNode)
			}
		}

		// Very naive implementation check but it works for now until someone starts to complain.
		if len(bodyNode.GetNodes()) > 0 {
			t.Implemented = true
		}
	}

	for _, clauseCtx := range ctx.AllCatchClause() {
		clause := NewCatchClauseStatement(t.ASTBuilder)
		t.Clauses = append(t.Clauses, clause.Parse(
			unit, contractNode, fnNode, bodyNode, t, clauseCtx.(*parser.CatchClauseContext),
		))
	}

	if ctx.Returns() != nil {
		t.Returns = true
	}

	returnParams := NewParameterList(t.ASTBuilder)
	if ctx.GetReturnParameters() != nil {
		returnParams.Parse(unit, t, ctx.GetReturnParameters())
	} else {
		returnParams.Src = t.Src
		returnParams.Src.ParentIndex = t.Id
	}
	t.ReturnParameters = returnParams
	return t
}
