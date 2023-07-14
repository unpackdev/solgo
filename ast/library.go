package ast

type LibraryName struct {
	ID                    int    `json:"id"`
	Name                  string `json:"name"`
	NodeType              string `json:"node_type"`
	ReferencedDeclaration int    `json:"referenced_declaration"`
	Src                   string `json:"src"`
}
