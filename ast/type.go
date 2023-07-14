package ast

type TypeName struct {
	ID               int               `json:"id"`
	Name             string            `json:"name"`
	NodeType         string            `json:"node_type"`
	Src              string            `json:"src"`
	TypeDescriptions *TypeDescriptions `json:"type_descriptions"`
}

type TypeDescriptions struct {
	TypeIdentifier string `json:"type_identifier"`
	TypeString     string `json:"type_string"`
}
