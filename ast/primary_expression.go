package ast

import (
	"encoding/hex"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type PrimaryExpression struct {
	*ASTBuilder

	Id                     int64             `json:"id"`
	NodeType               ast_pb.NodeType   `json:"node_type"`
	Kind                   ast_pb.NodeType   `json:"kind,omitempty"`
	Value                  string            `json:"value,omitempty"`
	HexValue               string            `json:"hex_value,omitempty"`
	Src                    SrcNode           `json:"src"`
	Name                   string            `json:"name,omitempty"`
	TypeDescription        TypeDescription   `json:"type_descriptions,omitempty"`
	OverloadedDeclarations []int64           `json:"overloaded_declarations"`
	ReferencedDeclaration  int64             `json:"referenced_declaration"`
	IsPure                 bool              `json:"is_pure"`
	ArgumentTypes          []TypeDescription `json:"argument_types,omitempty"`
}

func NewPrimaryExpression(b *ASTBuilder) *PrimaryExpression {
	return &PrimaryExpression{
		ASTBuilder:             b,
		OverloadedDeclarations: make([]int64, 0),
		NodeType:               ast_pb.NodeType_IDENTIFIER,
	}
}

func (p *PrimaryExpression) GetId() int64 {
	return p.Id
}

func (p *PrimaryExpression) GetType() ast_pb.NodeType {
	return p.NodeType
}

func (p *PrimaryExpression) GetSrc() SrcNode {
	return p.Src
}

func (p *PrimaryExpression) GetName() string {
	return p.Name
}

func (p *PrimaryExpression) GetTypeDescription() TypeDescription {
	return p.TypeDescription
}

func (p *PrimaryExpression) ToProto() NodeType {
	return ast_pb.PrimaryExpression{}
}

func (p *PrimaryExpression) Parse(
	unit *SourceUnit[Node[ast_pb.SourceUnit]],
	contractNode Node[NodeType],
	fnNode Node[NodeType],
	bodyNode *BodyNode,
	vDeclar *VariableDeclaration,
	expNode Node[NodeType],
	ctx *parser.PrimaryExpressionContext,
) {
	p.Id = p.GetNextID()
	p.Src = SrcNode{
		Id:     p.GetNextID(),
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

			return bodyNode.GetId()
		}(),
	}

	if expNode != nil {
		switch expNodeCtx := expNode.(type) {
		case *FunctionCall:
			for _, argument := range expNodeCtx.GetArguments() {
				p.ArgumentTypes = append(p.ArgumentTypes, argument.GetTypeDescription())
			}
		}
	}

	if ctx.Identifier() != nil {
		p.Name = ctx.Identifier().GetText()
	}

	referenceFound := false

	// Search for argument reference in state variable declarations.
	for _, node := range p.currentStateVariables {
		if node.GetName() == ctx.GetText() {
			referenceFound = true
			p.ReferencedDeclaration = node.Id
			p.TypeDescription = TypeDescription{
				TypeIdentifier: node.TypeName.GetTypeDescriptions().TypeIdentifier,
				TypeString:     node.TypeName.GetTypeDescriptions().TypeString,
			}
		}
	}

	// Search for argument reference in statement declarations.
	if !referenceFound {
		for _, statement := range bodyNode.Statements {
			switch statementCtx := statement.(type) {
			case *VariableDeclaration:
				for _, declaration := range statementCtx.Declarations {
					if declaration.Name == ctx.GetText() {
						referenceFound = true
						p.ReferencedDeclaration = declaration.Id
						p.TypeDescription = declaration.TypeName.TypeDescription
					}
				}
			default:
				fmt.Println("Statement type: ", reflect.TypeOf(statement))
			}

			/*
				for _, argument := range statement.GetArguments() {
					if argument.GetName() == expressionCtx.GetText() {
						referenceFound = true
						toReturn.ReferencedDeclaration = argument.Id
						toReturn.TypeDescriptions = argument.GetTypeDescriptions()
					}
				} */
		}
	}

	// If search for reference in statement declarations failed,
	// search for reference in function parameters.
	if !referenceFound && fnNode != nil {
		fnNode := fnNode.(FunctionNode[ast_pb.Function])
		if fnNode.GetParameters() != nil {
			for _, parameter := range fnNode.GetParameters().Parameters {
				if parameter.Name == ctx.GetText() {
					referenceFound = true
					p.ReferencedDeclaration = parameter.Id
					p.TypeDescription = parameter.TypeName.TypeDescription
				}
			}
		}
	}

	if !referenceFound {
		/* 		for _, errorNode := range b.currentErrors {
			if errorNode.GetName() == expressionCtx.GetText() {
				referenceFound = true
				toReturn.ReferencedDeclaration = errorNode.Id
				toReturn.TypeDescriptions = errorNode.GetTypeName().GetTypeDescriptions()
			}
		} */
	}

	literalCtx := ctx.Literal()
	if literalCtx != nil {
		p.NodeType = ast_pb.NodeType_LITERAL
		p.IsPure = true

		if literalCtx.BooleanLiteral() != nil {
			if p.Name == "true" || p.Name == "false" {
				p.Name = ""
			}

			p.Kind = ast_pb.NodeType_BOOLEAN
			p.Value = strings.TrimSpace(
				// There can be hex 22 at beginning and end of literal.
				// We should drop it as that's ASCII for double quote.
				strings.ReplaceAll(literalCtx.BooleanLiteral().GetText(), "\"", ""),
			)
			p.HexValue = hex.EncodeToString([]byte(p.Value))

			p.TypeDescription = TypeDescription{
				TypeIdentifier: "t_bool",
				TypeString:     "bool",
			}
		} else if literalCtx.StringLiteral() != nil {
			p.Name = ""
			p.Kind = ast_pb.NodeType_STRING

			p.Value = strings.TrimSpace(
				// There can be hex 22 at beginning and end of literal.
				// We should drop it as that's ASCII for double quote.
				strings.ReplaceAll(literalCtx.StringLiteral().GetText(), "\"", ""),
			)
			p.HexValue = hex.EncodeToString([]byte(p.Value))

			p.TypeDescription = TypeDescription{
				TypeIdentifier: "t_string_literal",
				TypeString: fmt.Sprintf(
					"literal_string %s",
					literalCtx.StringLiteral().GetText(),
				),
			}
		} else if literalCtx.NumberLiteral() != nil {
			p.Kind = ast_pb.NodeType_NUMBER

			p.Value = strings.TrimSpace(
				// There can be hex 22 at beginning and end of literal.
				// We should drop it as that's ASCII for double quote.
				strings.ReplaceAll(literalCtx.NumberLiteral().GetText(), "\"", ""),
			)
			p.HexValue = hex.EncodeToString([]byte(p.Value))

			// Check if the number is a floating-point number
			if strings.Contains(p.Value, ".") {
				parts := strings.Split(p.Value, ".")

				// The numerator is the number without the decimal point
				numerator, _ := strconv.Atoi(parts[0] + parts[1])

				// The denominator is a power of 10 equal to the number of digits in the fractional part
				denominator := int(math.Pow(10, float64(len(parts[1]))))

				p.TypeDescription = TypeDescription{
					TypeIdentifier: fmt.Sprintf("t_rational_%d_by_%d", numerator, denominator),
					TypeString: fmt.Sprintf(
						"fixed_const %s",
						literalCtx.NumberLiteral().GetText(),
					),
				}
			} else {
				numerator, _ := strconv.Atoi(p.Value)

				// The denominator for an integer is 1
				denominator := 1

				p.TypeDescription = TypeDescription{
					TypeIdentifier: fmt.Sprintf("t_rational_%d_by_%d", numerator, denominator),
					TypeString: fmt.Sprintf(
						"int_const %s",
						literalCtx.NumberLiteral().GetText(),
					),
				}
			}
		} else if literalCtx.HexStringLiteral() != nil {
			p.Kind = ast_pb.NodeType_HEX_STRING

			p.Value = strings.TrimSpace(
				// There can be hex 22 at beginning and end of literal.
				// We should drop it as that's ASCII for double quote.
				strings.ReplaceAll(literalCtx.StringLiteral().GetText(), "\"", ""),
			)
			p.HexValue = hex.EncodeToString([]byte(p.Value))

			p.TypeDescription = TypeDescription{
				TypeIdentifier: "t_string_hex_literal",
				TypeString: fmt.Sprintf(
					"literal_hex_string %s",
					literalCtx.StringLiteral().GetText(),
				),
			}
		} else {
			if ctx.GetText() == "msg" {
				p.TypeDescription = TypeDescription{
					TypeIdentifier: "t_magic_message",
					TypeString:     "msg",
				}
			}

			if p.TypeDescription.TypeString == "" {
				if expNode != nil {
					p.TypeDescription = expNode.GetTypeDescription()
				}
			}
		}
	}
}
