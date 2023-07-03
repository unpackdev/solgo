package ast

// UsingDirectiveNode represents a using directive in the AST.
type UsingDirectiveNode struct {
	Alias      string `json:"alias"`
	Type       string `json:"type"`
	IsWildcard bool   `json:"is_wildcard"`
	IsGlobal   bool   `json:"is_global"`
	IsUserDef  bool   `json:"is_user_defined"`
}

func (n *UsingDirectiveNode) Children() []Node {
	return nil
}
