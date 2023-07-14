package ast

import (
	"regexp"
	"strings"
)

func getLiterals(literal string) []string {
	// This regular expression matches sequences of word characters (letters, digits, underscores)
	// and sequences of non-word characters. It treats each match as a separate word.
	re := regexp.MustCompile(`\w+|\W+`)
	allLiterals := re.FindAllString(literal, -1)
	var literals []string
	for _, field := range allLiterals {
		field = strings.Trim(field, " ")
		if field != "" {
			// If the field is not empty after trimming spaces, add it to the literals
			literals = append(literals, field)
		}
	}
	return literals
}
