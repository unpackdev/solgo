package ir

import (
	"strings"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	ir_pb "github.com/unpackdev/protos/dist/go/ir"
	"github.com/unpackdev/solgo/ast"
)

// Pragma represents a Pragma in the Abstract Syntax Tree.
type Pragma struct {
	Unit     *ast.Pragma     `json:"ast"`
	Id       int64           `json:"id"`
	NodeType ast_pb.NodeType `json:"nodeType"`
	Literals []string        `json:"literals"`
	Text     string          `json:"text"`
}

// GetAST returns the underlying AST node for the Pragma.
func (p *Pragma) GetAST() *ast.Pragma {
	return p.Unit
}

// GetId returns the ID of the Pragma.
func (p *Pragma) GetId() int64 {
	return p.Id
}

// GetNodeType returns the AST node type of the Pragma.
func (p *Pragma) GetNodeType() ast_pb.NodeType {
	return p.NodeType
}

// GetLiterals returns the literals associated with the Pragma.
func (p *Pragma) GetLiterals() []string {
	return p.Literals
}

// GetText returns the text of the Pragma.
func (p *Pragma) GetText() string {
	return p.Text
}

// GetVersion extracts and returns the version information from the Pragma text.
func (p *Pragma) GetVersion() string {
	parts := strings.Split(p.Text, " ")
	return strings.Replace(parts[len(parts)-1], ";", "", -1)
}

// GetSrc returns the source code location associated with the Pragma.
func (p *Pragma) GetSrc() ast.SrcNode {
	return p.Unit.GetSrc()
}

// ToProto converts the Pragma to its corresponding protobuf representation.
func (p *Pragma) ToProto() *ir_pb.Pragma {
	proto := &ir_pb.Pragma{
		Id:       p.GetId(),
		NodeType: p.GetNodeType(),
		Literals: p.GetLiterals(),
		Text:     p.GetText(),
	}

	return proto
}

// processPragma processes the given pragma and returns the corresponding Pragma object.
func (b *Builder) processPragma(unit *ast.Pragma) *Pragma {
	return &Pragma{
		Unit:     unit,
		Id:       unit.GetId(),
		NodeType: unit.GetType(),
		Literals: unit.GetLiterals(),
		Text:     unit.GetText(),
	}
}
