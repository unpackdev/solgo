package ast

type Body struct {
	ID         int         `json:"id"`
	NodeType   string      `json:"node_type"`
	Src        string      `json:"src"`
	Statements []Statement `json:"statements"`
}
