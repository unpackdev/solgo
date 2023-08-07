package ir

// Language represents a programming language.
type Language string

// String returns the string representation of the Language.
func (l Language) String() string {
	return string(l)
}

// Constants representing different programming languages.
const (
	LanguageSolidity Language = "solidity"
	LanguageVyper    Language = "vyper"
)
