package ast

import (
	"encoding/json"
	"errors"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"go.uber.org/zap"
)

func unmarshalNode(data []byte, nodeType ast_pb.NodeType) (Node[NodeType], error) {
	switch nodeType {
	case ast_pb.NodeType_ENUM_DEFINITION:
		var toReturn *EnumDefinition
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("enum definition")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_ENUM_VALUE:
		var toReturn *Parameter
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("enum value")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_EVENT_DEFINITION:
		var toReturn *EventDefinition
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("event definition")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_EMIT_STATEMENT:
		var toReturn *Emit
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("emit statement")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_ERROR_DEFINITION:
		var toReturn *ErrorDefinition
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("error definition")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_VARIABLE_DECLARATION:
		var toReturn *VariableDeclaration
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("variable declaration")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_EXPRESSION_OPERATION:
		var toReturn *ExprOperation
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("expression operation")
			return nil, err
		}
	case ast_pb.NodeType_WHILE_STATEMENT:
		var toReturn *WhileStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("while statement")
			return nil, err
		}
	case ast_pb.NodeType_AND_OPERATION:
		var toReturn *AndOperation
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("and operation")
			return nil, err
		}
	case ast_pb.NodeType_FOR_STATEMENT:
		var toReturn *ForStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("for declaration")
			return nil, err
		}
	case ast_pb.NodeType_TRY_STATEMENT:
		var toReturn *TryStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("try declaration")
			return nil, err
		}
	case ast_pb.NodeType_DO_WHILE_STATEMENT:
		var toReturn *DoWhileStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("do-while declaration")
			return nil, err
		}
	case ast_pb.NodeType_TRY_CATCH_CLAUSE:
		var toReturn *CatchStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("catch declaration")
			return nil, err
		}
	case ast_pb.NodeType_BIT_XOR_OPERATION:
		var toReturn *BitXorOperation
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("bit xor operation")
			return nil, err
		}
	case ast_pb.NodeType_BIT_AND_OPERATION:
		var toReturn *BitAndOperation
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("bit and operation")
			return nil, err
		}
	case ast_pb.NodeType_BIT_OR_OPERATION:
		var toReturn *BitOrOperation
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("bit or operation")
			return nil, err
		}
	case ast_pb.NodeType_TUPLE_EXPRESSION:
		var toReturn *TupleExpression
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("tuple expression")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_CONDITIONAL_EXPRESSION:
		var toReturn *Conditional
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("conditional expression")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_PRAGMA_DIRECTIVE:
		var toReturn *Pragma
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("pragma directive")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_IMPORT_DIRECTIVE:
		var toReturn *Import
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("import directive")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_USING_FOR_DIRECTIVE:
		var toReturn *UsingDirective
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("using for directive")
			return nil, err
		}
		return toReturn, nil

	case ast_pb.NodeType_CONTRACT_DEFINITION:

		var tempMap map[string]json.RawMessage
		if err := json.Unmarshal(data, &tempMap); err != nil {
			return nil, err
		}

		if kind, ok := tempMap["kind"]; ok {
			var fKind ast_pb.NodeType
			if err := json.Unmarshal(kind, &fKind); err != nil {
				return nil, err
			}

			switch fKind {
			case ast_pb.NodeType_KIND_CONTRACT:
				var toReturn *Contract
				if err := json.Unmarshal(data, &toReturn); err != nil {
					panic("contract definition")
					return nil, err
				}
				return toReturn, nil
			case ast_pb.NodeType_KIND_LIBRARY:
				var toReturn *Library
				if err := json.Unmarshal(data, &toReturn); err != nil {
					panic("library definition")
					return nil, err
				}
				return toReturn, nil
			case ast_pb.NodeType_KIND_INTERFACE:
				var toReturn *Interface
				if err := json.Unmarshal(data, &toReturn); err != nil {
					panic("interface definition")
					return nil, err
				}
				return toReturn, nil
			}
		}

		var toReturn *Contract
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_FUNCTION_DEFINITION:
		var tempMap map[string]json.RawMessage
		if err := json.Unmarshal(data, &tempMap); err != nil {
			return nil, err
		}

		if kind, ok := tempMap["kind"]; ok {
			var fKind ast_pb.NodeType
			if err := json.Unmarshal(kind, &fKind); err != nil {
				return nil, err
			}

			switch fKind {
			case ast_pb.NodeType_CONSTRUCTOR:
				var toReturn *Constructor
				if err := json.Unmarshal(data, &toReturn); err != nil {
					return nil, err
				}
			case ast_pb.NodeType_KIND_FUNCTION:
				var toReturn *Function
				if err := json.Unmarshal(data, &toReturn); err != nil {
					return nil, err
				}
				return toReturn, nil
			case ast_pb.NodeType_FALLBACK:
				var toReturn *Fallback
				if err := json.Unmarshal(data, &toReturn); err != nil {
					return nil, err
				}
			case ast_pb.NodeType_RECEIVE:
				var toReturn *Receive
				if err := json.Unmarshal(data, &toReturn); err != nil {
					return nil, err
				}
			default:
				zap.L().Error(
					"unknown function kind while importing JSON",
					zap.String("kind", fKind.String()),
				)
			}
		}

		var toReturn *Function
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_MODIFIER_DEFINITION:
		var toReturn *ModifierDefinition
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil

	case ast_pb.NodeType_FALLBACK:
		var toReturn *Fallback
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil

	case ast_pb.NodeType_STRUCT_DEFINITION:
		var toReturn *StructDefinition
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_FUNCTION_CALL:
		var toReturn *FunctionCall
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_FUNCTION_CALL_OPTION:
		var toReturn *FunctionCallOption
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("function call option")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_PAYABLE_CONVERSION:
		var toReturn *PayableConversion
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("payable conversion")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_NEW_EXPRESSION:
		var toReturn *NewExpr
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("new expr")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_IDENTIFIER:
		var toReturn *PrimaryExpression
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("identifier")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_LITERAL:
		var toReturn *PrimaryExpression
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("literal")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_BINARY_OPERATION:
		var toReturn *BinaryOperation
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("binary operation")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_MEMBER_ACCESS:
		var toReturn *MemberAccessExpression
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("member access")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_RETURN_STATEMENT:
		var toReturn *ReturnStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("return statement")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_ASSIGNMENT:
		var toReturn *Assignment
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("assignment")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_REVERT_STATEMENT:
		var toReturn *RevertStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("revert statement")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_BLOCK:
		var toReturn *BodyNode
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("block")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_UNCHECKED_BLOCK:
		var toReturn *BodyNode
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("unchecked block")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_IF_STATEMENT:
		var toReturn *IfStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("if statement")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_BREAK:
		var toReturn *BreakStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("break statement")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_CONTINUE:
		var toReturn *ContinueStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("continue statement")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_INDEX_ACCESS:
		var toReturn *IndexAccess
		if err := json.Unmarshal(data, &toReturn); err != nil {
			panic("index access")
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_UNARY_OPERATION:
		var tempMap map[string]json.RawMessage
		if err := json.Unmarshal(data, &tempMap); err != nil {
			return nil, err
		}

		if kind, ok := tempMap["kind"]; ok {
			var fKind ast_pb.NodeType
			if err := json.Unmarshal(kind, &fKind); err != nil {
				return nil, err
			}

			switch fKind {
			case ast_pb.NodeType_KIND_UNARY_PREFIX:
				var toReturn *UnaryPrefix
				if err := json.Unmarshal(data, &toReturn); err != nil {
					panic("unary prefix operation")
					return nil, err
				}
				return toReturn, nil
			case ast_pb.NodeType_KIND_UNARY_SUFFIX:
				var toReturn *UnarySuffix
				if err := json.Unmarshal(data, &toReturn); err != nil {
					panic("unary suffix operation")
					return nil, err
				}
				return toReturn, nil
			}
		}

		return nil, errors.New("unknown unary operation kind while importing JSON")

		//
		// FUTURE IMPLEMENTATIONS....
		//

	case ast_pb.NodeType_ASSEMBLY_STATEMENT:
		var toReturn *AssemblyStatement
		/* 		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		} */
		return toReturn, nil
	case ast_pb.NodeType_YUL_STATEMENT:
		var toReturn *YulStatement
		/* 		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		} */
		return toReturn, nil
	default:
		panic(nodeType)
	}

	return nil, nil
}
