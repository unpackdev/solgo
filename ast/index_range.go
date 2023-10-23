package ast

import (
	"encoding/json"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// IndexRange represents an Index Range expression in the AST.
type IndexRange struct {
	*ASTBuilder

	Id               int64              `json:"id"`
	NodeType         ast_pb.NodeType    `json:"node_type"`
	Src              SrcNode            `json:"src"`
	LeftExpression   Node[NodeType]     `json:"left_expression"`
	RightExpression  Node[NodeType]     `json:"right_expression"`
	TypeDescriptions []*TypeDescription `json:"type_descriptions"`
}

// NewIndexRange creates a new instance of IndexRange with initialized values.
func NewIndexRangeAccessExpression(b *ASTBuilder) *IndexRange {
	return &IndexRange{
		ASTBuilder:       b,
		Id:               b.GetNextID(),
		NodeType:         ast_pb.NodeType_INDEX_RANGE_ACCESS,
		TypeDescriptions: make([]*TypeDescription, 0),
	}
}

// SetReferenceDescriptor is used to set reference descriptions for the IndexRange node.
// However, this function always returns false.
func (b *IndexRange) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the unique identifier of the IndexRange node.
func (f *IndexRange) GetId() int64 {
	return f.Id
}

// GetType returns the node type of the IndexRange.
func (f *IndexRange) GetType() ast_pb.NodeType {
	return f.NodeType
}

// GetSrc returns the source information of the IndexRange node.
func (f *IndexRange) GetSrc() SrcNode {
	return f.Src
}

// GetTypeDescription returns the type description associated with the IndexRange.
func (f *IndexRange) GetTypeDescription() *TypeDescription {
	return f.TypeDescriptions[0]
}

// GetNodes returns the list of nodes within the IndexRange.
func (f *IndexRange) GetNodes() []Node[NodeType] {
	return []Node[NodeType]{f.LeftExpression, f.RightExpression}
}

// GetLeftExpression returns the left expression of the IndexRange.
func (f *IndexRange) GetLeftExpression() Node[NodeType] {
	return f.LeftExpression
}

// GetRightExpression returns the right expression of the IndexRange.
func (f *IndexRange) GetRightExpression() Node[NodeType] {
	return f.RightExpression
}

func (f *IndexRange) UnmarshalJSON(data []byte) error {
	var tempMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	if id, ok := tempMap["id"]; ok {
		if err := json.Unmarshal(id, &f.Id); err != nil {
			return err
		}
	}

	if nodeType, ok := tempMap["node_type"]; ok {
		if err := json.Unmarshal(nodeType, &f.NodeType); err != nil {
			return err
		}
	}

	if src, ok := tempMap["src"]; ok {
		if err := json.Unmarshal(src, &f.Src); err != nil {
			return err
		}
	}

	if leftExpression, ok := tempMap["left_expression"]; ok {
		if err := json.Unmarshal(leftExpression, &f.LeftExpression); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(leftExpression, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["node_type"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(leftExpression, tempNodeType)
			if err != nil {
				return err
			}
			f.LeftExpression = node
		}
	}

	if rightExpression, ok := tempMap["right_expression"]; ok {
		if err := json.Unmarshal(rightExpression, &f.RightExpression); err != nil {
			var tempNodeMap map[string]json.RawMessage
			if err := json.Unmarshal(rightExpression, &tempNodeMap); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNodeMap["node_type"], &tempNodeType); err != nil {
				return err
			}

			node, err := unmarshalNode(rightExpression, tempNodeType)
			if err != nil {
				return err
			}
			f.RightExpression = node
		}
	}

	if typeDescriptions, ok := tempMap["type_descriptions"]; ok {
		if err := json.Unmarshal(typeDescriptions, &f.TypeDescriptions); err != nil {
			return err
		}
	}

	return nil
}

// ToProto converts the IndexRange node to its Protocol Buffers representation.
func (f *IndexRange) ToProto() NodeType {
	proto := ast_pb.IndexRange{
		Id:              f.GetId(),
		NodeType:        f.GetType(),
		Src:             f.GetSrc().ToProto(),
		LeftExpression:  f.GetLeftExpression().ToProto().(*v3.TypedStruct),
		RightExpression: f.GetRightExpression().ToProto().(*v3.TypedStruct),
		TypeDescription: f.GetTypeDescription().ToProto(),
	}

	return NewTypedStruct(&proto, "IndexRange")
}

// Parse parses the IndexRange expression from the provided context and constructs the IndexRange node.
func (f *IndexRange) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.IndexRangeAccessContext,
) Node[NodeType] {
	f.Id = f.GetNextID()
	f.Src = SrcNode{
		Id:     f.GetNextID(),
		Line:   int64(ctx.GetStart().GetLine()),
		Column: int64(ctx.GetStart().GetColumn()),
		Start:  int64(ctx.GetStart().GetStart()),
		End:    int64(ctx.GetStop().GetStop()),
		Length: int64(ctx.GetStop().GetStop() - ctx.GetStart().GetStart() + 1),
		ParentIndex: func() int64 {
			if vDeclar != nil {
				return vDeclar.GetId()
			}

			if expNode != nil {
				return expNode.GetId()
			}

			if bodyNode != nil {
				return bodyNode.GetId()
			}

			if fnNode != nil {
				return fnNode.GetId()
			}

			return contractNode.GetId()
		}(),
	}

	expression := NewExpression(f.ASTBuilder)

	f.LeftExpression = expression.Parse(unit, contractNode, fnNode, bodyNode, vDeclar, f, f.GetId(), ctx.Expression(0))
	f.RightExpression = expression.Parse(unit, contractNode, fnNode, bodyNode, vDeclar, f, f.GetId(), ctx.Expression(1))

	f.TypeDescriptions = append(f.TypeDescriptions, f.LeftExpression.GetTypeDescription())
	f.TypeDescriptions = append(f.TypeDescriptions, f.RightExpression.GetTypeDescription())

	return f
}
