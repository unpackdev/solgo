package ast

import (
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

type YulVariableNames struct {
	Id       int64           `json:"id"`
	Name     string          `json:"name"`
	NodeType ast_pb.NodeType `json:"node_type"`
	Src      SrcNode         `json:"src"`
}

type YulAssignment struct {
	*ASTBuilder

	Id            int64            `json:"id"`
	NodeType      ast_pb.NodeType  `json:"node_type"`
	Src           SrcNode          `json:"src"`
	VariableNames []*YulIdentifier `json:"variable_names"`
	Value         Node[NodeType]   `json:"value"`
}

func NewYulAssignment(b *ASTBuilder) *YulAssignment {
	return &YulAssignment{
		ASTBuilder:    b,
		Id:            b.GetNextID(),
		NodeType:      ast_pb.NodeType_YUL_ASSIGNMENT,
		VariableNames: make([]*YulIdentifier, 0),
	}
}

// SetReferenceDescriptor sets the reference descriptions of the YulAssignment node.
func (y *YulAssignment) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

func (y *YulAssignment) GetId() int64 {
	return y.Id
}

func (y *YulAssignment) GetType() ast_pb.NodeType {
	return y.NodeType
}

func (y *YulAssignment) GetSrc() SrcNode {
	return y.Src
}

func (y *YulAssignment) GetNodes() []Node[NodeType] {
	toReturn := make([]Node[NodeType], 0)
	toReturn = append(toReturn, y.Value)
	return toReturn
}

func (y *YulAssignment) GetTypeDescription() *TypeDescription {
	return &TypeDescription{}
}

func (y *YulAssignment) ToProto() NodeType {
	return ast_pb.Statement{}
}

func (y *YulAssignment) GetVariableNames() []*YulIdentifier {
	return y.VariableNames
}

// UnmarshalJSON unmarshals a given JSON byte array into a YulAssignment node.
func (f *YulAssignment) UnmarshalJSON(data []byte) error {
	return nil
}

func (y *YulAssignment) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	assemblyNode *Yul,
	statementNode *YulStatement,
	ctx *parser.YulAssignmentContext,
) Node[NodeType] {
	y.Src = SrcNode{
		Id:          y.GetNextID(),
		Line:        int64(ctx.GetStart().GetLine()),
		Column:      int64(ctx.GetStart().GetColumn()),
		Start:       int64(ctx.GetStart().GetStart()),
		End:         int64(ctx.GetStop().GetStop()),
		Length:      int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: assemblyNode.GetId(),
	}

	if ctx.AllYulPath() != nil {
		for _, path := range ctx.AllYulPath() {
			for _, identifier := range path.AllYulIdentifier() {
				y.VariableNames = append(y.VariableNames, &YulIdentifier{
					Id:       y.GetNextID(),
					NodeType: ast_pb.NodeType_YUL_IDENTIFIER,
					Src: SrcNode{
						Id:          y.GetNextID(),
						Line:        int64(identifier.GetSymbol().GetLine()),
						Column:      int64(identifier.GetSymbol().GetColumn()),
						Start:       int64(identifier.GetSymbol().GetStart()),
						End:         int64(identifier.GetSymbol().GetStop()),
						Length:      int64(identifier.GetSymbol().GetStop() - identifier.GetSymbol().GetStart() + 1),
						ParentIndex: y.GetId(),
					},
					Name: identifier.GetText(),
					NameLocation: SrcNode{
						Id:          y.GetNextID(),
						Line:        int64(identifier.GetSymbol().GetLine()),
						Column:      int64(identifier.GetSymbol().GetColumn()),
						Start:       int64(identifier.GetSymbol().GetStart()),
						End:         int64(identifier.GetSymbol().GetStop()),
						Length:      int64(identifier.GetSymbol().GetStop() - identifier.GetSymbol().GetStart() + 1),
						ParentIndex: y.GetId(),
					},
				})
			}
		}
	}

	if ctx.YulExpression() != nil {
		yExpression := NewYulExpressionStatement(y.ASTBuilder)
		y.Value = yExpression.Parse(
			unit, contractNode, fnNode, bodyNode, assemblyNode, statementNode,
			y, ctx.YulExpression().(*parser.YulExpressionContext),
		)
	}

	return y
}
