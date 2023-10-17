package ir

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unpackdev/solgo/ast"
)

func TestLinks(t *testing.T) {
	tests := []struct {
		name          string
		comments      []string
		expectedLinks []*Link
	}{
		{
			name:          "test no links",
			comments:      []string{"This is a test comment."},
			expectedLinks: []*Link{},
		},
		{
			name:          "test social links without punctuation",
			comments:      []string{"Follow us on https://twitter.com/ and https://t.me/unpackdev"},
			expectedLinks: []*Link{{Location: "https://twitter.com/", Social: true, Network: "twitter_x"}, {Location: "https://t.me/unpackdev", Social: true, Network: "telegram"}},
		},
		{
			name:          "test social links with punctuation",
			comments:      []string{"Follow us on https://twitter.com/. and https://t.me/unpackdev."},
			expectedLinks: []*Link{{Location: "https://twitter.com/", Social: true, Network: "twitter_x"}, {Location: "https://t.me/unpackdev", Social: true, Network: "telegram"}},
		},
		{
			name:          "test social links that are pointing to base site",
			comments:      []string{"Follow us on https://twitter.com. and https://t.me."},
			expectedLinks: []*Link{{Location: "https://twitter.com", Social: false}, {Location: "https://t.me", Social: false}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &RootSourceUnit{
				Unit: &ast.RootNode{Comments: makeComments(tt.comments)},
			}

			b := &Builder{}
			b.processLinks(root)
			assert.Equal(t, tt.expectedLinks, root.Links)
		})
	}
}

// Helper function to generate comments for test cases
func makeComments(texts []string) []*ast.Comment {
	var comments []*ast.Comment
	for _, text := range texts {
		comments = append(comments, &ast.Comment{Text: text})
	}
	return comments
}
