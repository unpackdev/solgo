package syntaxerrors

import "strings"

// ReplaceErrorMessage replaces a specific error message with a new message.
// It takes the original message, the old text to replace, and the new text, and returns the modified message.
func ReplaceErrorMessage(originalMsg, oldText, newText string) string {
	return strings.ReplaceAll(originalMsg, oldText, newText)
}
