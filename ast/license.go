package ast

import (
	"regexp"

	ast_pb "github.com/txpull/protos/dist/go/ast"
)

func GetLicense(comments []*CommentNode) string {
	licenseRegex := regexp.MustCompile(`SPDX-License-Identifier:\s*(.+)`)
	for _, comment := range comments {
		if comment.NodeType == ast_pb.NodeType_LICENSE {
			matches := licenseRegex.FindStringSubmatch(comment.Text)
			if len(matches) > 1 {
				return matches[1]
			}
		}
	}
	return "unknown"
}
