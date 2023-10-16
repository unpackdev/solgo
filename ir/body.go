package ir

import (
	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	ir_pb "github.com/unpackdev/protos/dist/go/ir"
	"github.com/unpackdev/solgo/ast"
)

// Statement is an interface that defines common methods for all statement-like
// AST nodes. It ensures consistency in extracting information from different node types.
type Statement interface {
	GetId() int64
	GetNodeType() ast_pb.NodeType
	GetKind() ast_pb.NodeType
	GetTypeDescription() *ast_pb.TypeDescription
	GetNodes() []Statement
	ToProto() *v3.TypedStruct
}

// Expression is an interface that abstracts operations on expression-like AST nodes.
// It is useful in ensuring consistent interaction with different expression nodes.
type Expression interface {
	GetId() int64
	GetType() ast_pb.NodeType
	GetName() string
	GetTypeDescription() *ast.TypeDescription
	GetReferencedDeclaration() int64
}

// Body represents a generic body of a construct, which can contain multiple statements.
type Body struct {
	Unit       *ast.BodyNode   `json:"ast"` // Original AST node reference
	Id         int64           `json:"id"`
	NodeType   ast_pb.NodeType `json:"node_type"`
	Kind       ast_pb.NodeType `json:"kind"`
	Statements []Statement     `json:"statements"`
}

// GetAST returns the original AST node reference for the body.
func (e *Body) GetAST() *ast.BodyNode {
	return e.Unit
}

// GetId returns the unique identifier of the body node.
func (e *Body) GetId() int64 {
	return e.Id
}

// GetNodeType returns the type of the body node.
func (e *Body) GetNodeType() ast_pb.NodeType {
	return e.NodeType
}

// GetKind returns the kind of the body node.
func (e *Body) GetKind() ast_pb.NodeType {
	return e.Kind
}

// GetSrc returns the source location of the body node.
func (e *Body) GetSrc() ast.SrcNode {
	return e.Unit.GetSrc()
}

// GetNodes returns a list of statements contained within the body.
func (e *Body) GetNodes() []Statement {
	return e.Statements
}

// GetStatements returns a list of statements contained within the body.
func (e *Body) GetStatements() []Statement {
	return e.Statements
}

// ToProto converts the Body into its protocol buffer representation.
func (e *Body) ToProto() *ir_pb.Body {
	proto := &ir_pb.Body{
		Id:         e.GetId(),
		NodeType:   e.GetNodeType(),
		Kind:       e.GetKind(),
		Statements: make([]*v3.TypedStruct, 0),
	}

	for _, statement := range e.GetNodes() {
		proto.Statements = append(proto.Statements, statement.ToProto())
	}

	return proto
}

// processFunctionBody processes the body of a function and returns
// its intermediate representation. Currently, it only processes function call statements.
func (b *Builder) processFunctionBody(fn *Function, unit *ast.BodyNode) *Body {
	body := &Body{
		Unit:       unit,
		Id:         unit.GetId(),
		NodeType:   unit.GetType(),
		Kind:       unit.GetKind(),
		Statements: make([]Statement, 0),
	}

	for _, statement := range unit.GetNodes() {
		switch stmt := statement.(type) {
		case *ast.FunctionCall:
			body.Statements = append(body.Statements, b.processFunctionCall(fn, stmt))
		}
	}

	return body
}
