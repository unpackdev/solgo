package ast

import (
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	ast_pb "github.com/txpull/protos/dist/go/ast"
	"github.com/txpull/solgo/parser"
)

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
				comment := &ast_pb.Comment{
					Id: int64(len(b.comments)),
					Src: &ast_pb.Src{
						Line:        int64(token.GetLine()),
						Column:      int64(token.GetColumn()),
						Start:       int64(token.GetStart()),
						End:         int64(token.GetStop()),
						Length:      int64(token.GetStop() - token.GetStart() + 1),
						ParentIndex: b.nextID,
					},
					NodeType: ast_pb.NodeType_NODE_TYPE_COMMENT,
					Value:    strings.TrimSpace(token.GetText()),
				}

				if strings.Contains(token.GetText(), "SPDX-License-Identifier") {
					comment.NodeType = ast_pb.NodeType_NODE_TYPE_LICENSE
				}

				b.comments = append(b.comments, comment)
			}
			if token.GetTokenType() == parser.SolidityLexerCOMMENT {
				comment := &ast_pb.Comment{
					Id: int64(len(b.comments)),
					Src: &ast_pb.Src{
						Line:        int64(token.GetLine()),
						Column:      int64(token.GetColumn()),
						Start:       int64(token.GetStart()),
						End:         int64(token.GetStop()),
						Length:      int64(token.GetStop() - token.GetStart() + 1),
						ParentIndex: b.nextID,
					},
					NodeType: ast_pb.NodeType_NODE_TYPE_COMMENT_MULTILINE,
					Value:    strings.TrimSpace(token.GetText()),
				}

				if strings.Contains(token.GetText(), "SPDX-License-Identifier") {
					comment.NodeType = ast_pb.NodeType_NODE_TYPE_LICENSE
				}

				b.comments = append(b.comments, comment)
			}
		}

		// We should not iterate over comments again.
		b.commentsParsed = true
	}
}

func (b *ASTBuilder) GetLicense() string {
	licenseRegex := regexp.MustCompile(`SPDX-License-Identifier:\s*(.+)`)
	for _, comment := range b.comments {
		if comment.NodeType == ast_pb.NodeType_NODE_TYPE_LICENSE {
			matches := licenseRegex.FindStringSubmatch(comment.Value)
			if len(matches) > 1 {
				return matches[1]
			}
		}
	}

	return ""
}
