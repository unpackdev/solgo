package ast

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

type CommentNode struct {
	// Id is the unique identifier of the comment node.
	Id int64 `json:"id"`
	// Src is the source node locations of the comment node.
	Src SrcNode `json:"src"`
	// NodeType is the type of the AST node.
	NodeType ast_pb.NodeType `json:"node_type"`
	// Value is the value of the comment node.
	Text string `json:"text"`
}

func (c *CommentNode) GetId() int64 {
	return c.Id
}

func (c *CommentNode) GetType() ast_pb.NodeType {
	return c.NodeType
}

func (c *CommentNode) GetSrc() SrcNode {
	return c.Src
}

func (c *CommentNode) GetText() string {
	return c.Text
}

func (c *CommentNode) ToProto() *ast_pb.Comment {
	return &ast_pb.Comment{
		Id:       c.Id,
		NodeType: c.NodeType,
		Src:      c.Src.ToProto(),
		Text:     c.Text,
	}
}

// EnterEveryRule is called when the parser enters any rule in the grammar.
// It is used to search for license and any comments that code has.
// ANTLR parser by default have comments disabled to be parsed as tokens, so we need to
// search for them manually using the CommonTokenStream.
func (b *ASTBuilder) EnterEveryRule(ctx antlr.ParserRuleContext) {
	if !b.commentsParsed {
		stream := b.parser.GetTokenStream().(*antlr.CommonTokenStream)
		tokens := stream.GetAllTokens()

		for _, token := range tokens {
			if token.GetTokenType() == parser.SolidityLexerLINE_COMMENT {
				comment := &CommentNode{
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
				comment := &CommentNode{
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
