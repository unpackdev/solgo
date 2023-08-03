package ast

import (
	"regexp"

	ast_pb "github.com/txpull/protos/dist/go/ast"
)

// getLicense extracts the license from the provided comments.
// It uses a regular expression to match the SPDX-License-Identifier pattern in the comment text.
// If a license is found, it is returned as a string.
// If no license is found, it returns the string "unknown".
//
// The function takes a slice of CommentNode pointers as input, where each CommentNode represents a comment in the source code.
// Each CommentNode has a NodeType and Text. The NodeType indicates the type of the node, and the Text contains the text of the comment.
//
// The function iterates over the provided comments. For each comment, it checks if the NodeType is a license.
// If it is, the function uses a regular expression to find the SPDX-License-Identifier in the comment text.
// If the SPDX-License-Identifier is found, the function returns the license as a string.
// If the SPDX-License-Identifier is not found in any of the comments, the function returns the string "unknown".
func getLicense(comments []*Comment) string {
	// Define the regular expression for the SPDX-License-Identifier.
	licenseRegex := regexp.MustCompile(`SPDX-License-Identifier:\s*(.+)`)

	// Iterate over the provided comments.
	for _, comment := range comments {
		// Check if the NodeType of the comment is a license.
		if comment.NodeType == ast_pb.NodeType_LICENSE {
			// Find the SPDX-License-Identifier in the comment text.
			matches := licenseRegex.FindStringSubmatch(comment.Text)

			// If the SPDX-License-Identifier is found, return the license as a string.
			if len(matches) > 1 {
				return matches[1]
			}
		}
	}

	// If the SPDX-License-Identifier is not found in any of the comments, return the string "unknown".
	return "unknown"
}
