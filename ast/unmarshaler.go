package ast

import (
	"errors"
	"github.com/goccy/go-json"

	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"go.uber.org/zap"
)

func unmarshalNode(data []byte, nodeType ast_pb.NodeType) (Node[NodeType], error) {
	switch nodeType {
	case ast_pb.NodeType_ENUM_DEFINITION:
		var toReturn *EnumDefinition
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_ENUM_VALUE:
		var toReturn *Parameter
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_EVENT_DEFINITION:
		var toReturn *EventDefinition
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_EMIT_STATEMENT:
		var toReturn *Emit
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_ERROR_DEFINITION:
		var toReturn *ErrorDefinition
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_VARIABLE_DECLARATION:
		var toReturn *VariableDeclaration
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_EXPRESSION_OPERATION:
		var toReturn *ExprOperation
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
	case ast_pb.NodeType_WHILE_STATEMENT:
		var toReturn *WhileStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
	case ast_pb.NodeType_AND_OPERATION:
		var toReturn *AndOperation
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
	case ast_pb.NodeType_FOR_STATEMENT:
		var toReturn *ForStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
	case ast_pb.NodeType_TRY_STATEMENT:
		var toReturn *TryStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
	case ast_pb.NodeType_DO_WHILE_STATEMENT:
		var toReturn *DoWhileStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
	case ast_pb.NodeType_TRY_CATCH_CLAUSE:
		var toReturn *CatchStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
	case ast_pb.NodeType_BIT_XOR_OPERATION:
		var toReturn *BitXorOperation
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
	case ast_pb.NodeType_BIT_AND_OPERATION:
		var toReturn *BitAndOperation
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
	case ast_pb.NodeType_BIT_OR_OPERATION:
		var toReturn *BitOrOperation
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
	case ast_pb.NodeType_TUPLE_EXPRESSION:
		var toReturn *TupleExpression
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_CONDITIONAL_EXPRESSION:
		var toReturn *Conditional
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_PRAGMA_DIRECTIVE:
		var toReturn *Pragma
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_IMPORT_DIRECTIVE:
		var toReturn *Import
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_USING_FOR_DIRECTIVE:
		var toReturn *UsingDirective
		if err := json.Unmarshal(data, &toReturn); err != nil {
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
					return nil, err
				}
				return toReturn, nil
			case ast_pb.NodeType_KIND_LIBRARY:
				var toReturn *Library
				if err := json.Unmarshal(data, &toReturn); err != nil {
					return nil, err
				}
				return toReturn, nil
			case ast_pb.NodeType_KIND_INTERFACE:
				var toReturn *Interface
				if err := json.Unmarshal(data, &toReturn); err != nil {
					return nil, err
				}
				return toReturn, nil
			default:
				zap.L().Error(
					"unknown contract kind while importing JSON",
					zap.String("kind", fKind.String()),
				)
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
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_PAYABLE_CONVERSION:
		var toReturn *PayableConversion
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_NEW_EXPRESSION:
		var toReturn *NewExpr
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_IDENTIFIER:
		var toReturn *PrimaryExpression
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_LITERAL:
		var toReturn *PrimaryExpression
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_PLACEHOLDER_STATEMENT:
		var toReturn *PrimaryExpression
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_BINARY_OPERATION:
		var toReturn *BinaryOperation
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_MEMBER_ACCESS:
		var toReturn *MemberAccessExpression
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_RETURN_STATEMENT:
		var toReturn *ReturnStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_ASSIGNMENT:
		var toReturn *Assignment
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_REVERT_STATEMENT:
		var toReturn *RevertStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_BLOCK:
		var toReturn *BodyNode
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_UNCHECKED_BLOCK:
		var toReturn *BodyNode
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_IF_STATEMENT:
		var toReturn *IfStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_BREAK:
		var toReturn *BreakStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_CONTINUE:
		var toReturn *ContinueStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_INDEX_ACCESS:
		var toReturn *IndexAccess
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_SHIFT_OPERATION:
		var toReturn *ShiftOperation
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_USER_DEFINED_VALUE_TYPE:
		var toReturn *UserDefinedValueTypeDefinition
		if err := json.Unmarshal(data, &toReturn); err != nil {
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
					return nil, err
				}
				return toReturn, nil
			case ast_pb.NodeType_KIND_UNARY_SUFFIX:
				var toReturn *UnarySuffix
				if err := json.Unmarshal(data, &toReturn); err != nil {
					return nil, err
				}
				return toReturn, nil
			}
		}

		return nil, errors.New("unknown unary operation kind while importing JSON")
	case ast_pb.NodeType_INLINE_ARRAY:
		var toReturn *InlineArray
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil

	//
	// ASSEMBLY / YUL NODES
	//

	case ast_pb.NodeType_ASSEMBLY_STATEMENT:
		var toReturn *Yul
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_YUL_STATEMENT:
		var toReturn *YulStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_YUL_VARIABLE_DECLARATION:
		var toReturn *YulAssignment
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_YUL_ASSIGNMENT:
		var toReturn *YulAssignment
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_YUL_BLOCK:
		var toReturn *YulBlockStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_YUL_FOR:
		var toReturn *YulForStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_YUL_IF:
		var toReturn *YulIfStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_YUL_SWITCH:
		var toReturn *YulSwitchStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_YUL_SWITCH_CASE:
		var toReturn *YulSwitchCaseStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_YUL_LITERAL:
		var toReturn *YulLiteralStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_YUL_FUNCTION_CALL:
		var toReturn *YulFunctionCallStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_YUL_EXPRESSION:
		var toReturn *YulExpressionStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_YUL_FUNCTION_DEFINITION:
		var toReturn *YulFunctionDefinition
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_YUL_BREAK:
		var toReturn *YulBreakStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_YUL_CONTINUE:
		var toReturn *YulContinueStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_YUL_LEAVE:
		var toReturn *YulLeaveStatement
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	case ast_pb.NodeType_YUL_IDENTIFIER:
		var toReturn *YulIdentifier
		if err := json.Unmarshal(data, &toReturn); err != nil {
			return nil, err
		}
		return toReturn, nil
	default:
		zap.L().Error(
			"unknown node type while importing JSON",
			zap.String("kind", nodeType.String()),
		)
	}

	return nil, nil
}
