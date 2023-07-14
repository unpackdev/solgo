package ast

import (
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/txpull/solgo/parser"
)

type CommentNode struct {
	ID       int64    `json:"id"`
	NodeType NodeType `json:"node_type"`
	Src      Src      `json:"src"`
	Value    string   `json:"value"`
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
					ID: int64(len(b.comments)),
					Src: Src{
						Line:   token.GetLine(),
						Start:  token.GetStart(),
						End:    token.GetStop(),
						Length: token.GetStop() - token.GetStart() + 1,
						Index:  b.nextID,
					},
					NodeType: NodeTypeCommentLine,
					Value:    strings.TrimSpace(token.GetText()),
				}

				if strings.Contains(token.GetText(), "SPDX-License-Identifier") {
					comment.NodeType = NodeTypeLicense
				}

				b.comments = append(b.comments, comment)
			}
			if token.GetTokenType() == parser.SolidityLexerCOMMENT {
				comment := &CommentNode{
					ID: int64(len(b.comments)),
					Src: Src{
						Line:   token.GetLine(),
						Start:  token.GetStart(),
						End:    token.GetStop(),
						Length: token.GetStop() - token.GetStart() + 1,
						Index:  b.nextID,
					},
					NodeType: NodeTypeCommentMultiLine,
					Value:    strings.TrimSpace(token.GetText()),
				}

				if strings.Contains(token.GetText(), "SPDX-License-Identifier") {
					comment.NodeType = NodeTypeLicense
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
		if comment.NodeType == NodeTypeLicense {
			matches := licenseRegex.FindStringSubmatch(comment.Value)
			if len(matches) > 1 {
				return matches[1]
			}
		}
	}

	return ""
}
