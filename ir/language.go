package ir

type Language string

func (l Language) String() string {
	return string(l)
}

const (
	LanguageSolidity Language = "solidity"
	LanguageVyper    Language = "vyper"
)
