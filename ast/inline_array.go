package ast

import (
	"encoding/json"
	"regexp"
	"strings"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// InlineArray represents a for loop statement in the AST.
type InlineArray struct {
	*ASTBuilder

	Id               int64              `json:"id"`                // Unique identifier for the InlineArray node.
	NodeType         ast_pb.NodeType    `json:"node_type"`         // Type of the AST node.
	Src              SrcNode            `json:"src"`               // Source location information.
	TypeDescriptions []*TypeDescription `json:"type_descriptions"` // Type descriptions of the InlineArray node.
	Expressions      []Node[NodeType]   `json:"expressions"`       // List of expressions in the InlineArray node.
	Empty            bool               `json:"empty"`             // Indicates whether the InlineArray node is empty.
	TypeDescription  *TypeDescription   `json:"type_description"`  // Type description of the InlineArray node.
}

// NewInlineArrayExpression creates a new InlineArray node with a given ASTBuilder.
func NewInlineArrayExpression(b *ASTBuilder) *InlineArray {
	return &InlineArray{
		ASTBuilder: b,
		Id:         b.GetNextID(),
		NodeType:   ast_pb.NodeType_INLINE_ARRAY,
	}
}

// SetReferenceDescriptor sets the reference descriptions of the InlineArray node.
// We don't need to do any reference description updates here, at least for now...
func (f *InlineArray) SetReferenceDescriptor(refId int64, refDesc *TypeDescription) bool {
	return false
}

// GetId returns the ID of the InlineArray node.
func (f *InlineArray) GetId() int64 {
	return f.Id
}

// GetType returns the NodeType of the InlineArray node.
func (f *InlineArray) GetType() ast_pb.NodeType {
	return f.NodeType
}

// GetSrc returns the SrcNode of the InlineArray node.
func (f *InlineArray) GetSrc() SrcNode {
	return f.Src
}

// IsEmpty returns true if the InlineArray node is empty.
func (f *InlineArray) IsEmpty() bool {
	return f.Empty
}

// GetExpressions returns the list of associated expressions.
func (f *InlineArray) GetExpressions() []Node[NodeType] {
	return f.Expressions
}

func (f *InlineArray) GetNodes() []Node[NodeType] {
	toReturn := []Node[NodeType]{}
	toReturn = append(toReturn, f.Expressions...)
	return toReturn
}

func (f *InlineArray) GetTypeDescriptions() []*TypeDescription {
	return f.TypeDescriptions
}

func (f *InlineArray) GetTypeDescription() *TypeDescription {
	return f.TypeDescription
}

func (f *InlineArray) UnmarshalJSON(data []byte) error {
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

	if typeDescriptions, ok := tempMap["type_descriptions"]; ok {
		if err := json.Unmarshal(typeDescriptions, &f.TypeDescriptions); err != nil {
			return err
		}
	}

	if empty, ok := tempMap["empty"]; ok {
		if err := json.Unmarshal(empty, &f.Empty); err != nil {
			return err
		}
	}

	if typeDescription, ok := tempMap["type_description"]; ok {
		if err := json.Unmarshal(typeDescription, &f.TypeDescription); err != nil {
			return err
		}
	}

	if expressions, ok := tempMap["expressions"]; ok {
		var nodes []json.RawMessage
		if err := json.Unmarshal(expressions, &nodes); err != nil {
			return err
		}

		for _, node := range nodes {
			var tempNode map[string]json.RawMessage
			if err := json.Unmarshal(node, &tempNode); err != nil {
				return err
			}

			var tempNodeType ast_pb.NodeType
			if err := json.Unmarshal(tempNode["node_type"], &tempNodeType); err != nil {
				return err
			}

			parsedNode, err := unmarshalNode(node, tempNodeType)
			if err != nil {
				return err
			}

			f.Expressions = append(f.Expressions, parsedNode)
		}

	}

	return nil
}

func (f *InlineArray) ToProto() NodeType {
	proto := ast_pb.InlineArray{
		Id:              f.GetId(),
		NodeType:        f.GetType(),
		Src:             f.GetSrc().ToProto(),
		IsEmpty:         f.IsEmpty(),
		TypeDescription: f.TypeDescription.ToProto(),
	}

	for _, expr := range f.GetExpressions() {
		proto.Expressions = append(proto.Expressions, expr.ToProto().(*v3.TypedStruct))
	}

	for _, typeDesc := range f.GetTypeDescriptions() {
		proto.TypeDescriptions = append(proto.TypeDescriptions, typeDesc.ToProto())
	}

	return NewTypedStruct(&proto, "InlineArray")
}

func (f *InlineArray) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.InlineArrayContext,
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
	f.Empty = ctx.IsEmpty()

	expression := NewExpression(f.ASTBuilder)
	for _, expr := range ctx.InlineArrayExpression().AllExpression() {
		parsedExp := expression.Parse(unit, contractNode, fnNode, bodyNode, vDeclar, f, f.GetId(), expr)
		f.Expressions = append(
			f.Expressions,
			parsedExp,
		)
		f.TypeDescriptions = append(f.TypeDescriptions, parsedExp.GetTypeDescription())
	}

	f.TypeDescription = f.buildTypeDescription()

	return f
}

// buildTypeDescription constructs the type description of the Function node.
func (f *InlineArray) buildTypeDescription() *TypeDescription {
	typeString := "["
	typeIdentifier := "t_inline_array_"
	typeStrings := make([]string, 0)
	typeIdentifiers := make([]string, 0)

	for _, paramType := range f.GetTypeDescriptions() {
		if paramType == nil {
			typeStrings = append(typeStrings, "unknown")
			typeIdentifiers = append(typeIdentifiers, "unknown")
			continue
		}

		typeStrings = append(typeStrings, paramType.TypeString)
		typeIdentifiers = append(typeIdentifiers, "$_"+paramType.TypeIdentifier)
	}

	typeString += strings.Join(typeStrings, ",") + "]"
	typeIdentifier += strings.Join(typeIdentifiers, "$")

	if !strings.HasSuffix(typeIdentifier, "$") {
		typeIdentifier += "$"
	}

	re := regexp.MustCompile(`\${2,}`)
	typeIdentifier = re.ReplaceAllString(typeIdentifier, "$")

	return &TypeDescription{
		TypeString:     typeString,
		TypeIdentifier: typeIdentifier,
	}
}
