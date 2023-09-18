package ast

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
	ast_pb "github.com/unpackdev/protos/dist/go/ast"
	"github.com/unpackdev/solgo/parser"
)

// Comment represents a comment in an abstract syntax tree.
type Comment struct {
	Id       int64           `json:"id"`
	Src      SrcNode         `json:"src"`
	NodeType ast_pb.NodeType `json:"node_type"`
	Text     string          `json:"text"`
}

// GetId returns the ID of the Comment.
func (c *Comment) GetId() int64 {
	return c.Id
}

// GetType returns the NodeType of the Comment.
func (c *Comment) GetType() ast_pb.NodeType {
	return c.NodeType
}

// GetSrc returns the source information of the Comment.
func (c *Comment) GetSrc() SrcNode {
	return c.Src
}

// GetText returns the text value of the Comment.
func (c *Comment) GetText() string {
	return c.Text
}

// ToProto converts the Comment to its corresponding protocol buffer representation.
func (c *Comment) ToProto() *ast_pb.Comment {
	return &ast_pb.Comment{
		Id:       c.GetId(),
		NodeType: c.GetType(),
		Src:      c.GetSrc().ToProto(),
		Text:     c.GetText(),
	}
}

// EnterEveryRule is called when the parser enters any rule in the grammar.
// It is used to search for licenses and comments in the code.
// ANTLR parser, by default, has comments disabled to be parsed as tokens.
// Therefore, we manually search for them using the CommonTokenStream.
func (b *ASTBuilder) EnterEveryRule(ctx antlr.ParserRuleContext) {
	if !b.commentsParsed {
		stream := b.parser.GetTokenStream().(*antlr.CommonTokenStream)
		tokens := stream.GetAllTokens()

		for _, token := range tokens {
			if token.GetTokenType() == parser.SolidityLexerLINE_COMMENT {
				comment := &Comment{
					Id: b.GetNextID(),
					Src: SrcNode{
						Line:        int64(token.GetLine()),
						Column:      int64(token.GetColumn()),
						Start:       int64(token.GetStart()),
						End:         int64(token.GetStop()),
						Length:      int64(token.GetStop() - token.GetStart() + 1),
						ParentIndex: b.nextID,
					},
					NodeType: ast_pb.NodeType_COMMENT,
					Text:     strings.TrimSpace(token.GetText()),
				}

				if strings.Contains(token.GetText(), "SPDX-License-Identifier") {
					comment.NodeType = ast_pb.NodeType_LICENSE
				}

				b.comments = append(b.comments, comment)
			}
			if token.GetTokenType() == parser.SolidityLexerCOMMENT {
				comment := &Comment{
					Id: b.GetNextID(),
					Src: SrcNode{
						Line:        int64(token.GetLine()),
						Column:      int64(token.GetColumn()),
						Start:       int64(token.GetStart()),
						End:         int64(token.GetStop()),
						Length:      int64(token.GetStop() - token.GetStart() + 1),
						ParentIndex: b.nextID,
					},
					NodeType: ast_pb.NodeType_COMMENT_MULTILINE,
					Text:     strings.TrimSpace(token.GetText()),
				}

				if strings.Contains(token.GetText(), "SPDX-License-Identifier") {
					comment.NodeType = ast_pb.NodeType_LICENSE
				}

				b.comments = append(b.comments, comment)
			}
		}

		// We should not iterate over comments again.
		b.commentsParsed = true
	}
}
