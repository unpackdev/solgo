package ast

// Symbol represents a symbol in the Solidity abstract syntax tree (AST).
type Symbol struct {
	Id           int64  `json:"id"`            // Unique identifier for the symbol
	Name         string `json:"name"`          // Name of the symbol
	AbsolutePath string `json:"absolute_path"` // Absolute path to the symbol
}

// NewSymbol creates a new Symbol instance with the provided attributes.
func NewSymbol(id int64, name string, absolutePath string) Symbol {
	return Symbol{
		Id:           id,
		Name:         name,
		AbsolutePath: absolutePath,
	}
}

// GetId returns the unique identifier of the symbol.
func (s Symbol) GetId() int64 {
	return s.Id
}

// GetName returns the name of the symbol.
func (s Symbol) GetName() string {
	return s.Name
}

// GetAbsolutePath returns the absolute path to the symbol.
func (s Symbol) GetAbsolutePath() string {
	return s.AbsolutePath
}
