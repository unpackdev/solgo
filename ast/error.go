package ast

type ErrorNode struct {
	Name   string            `json:"name"`
	Values []*ErrorValueNode `json:"values"`
}

func (e *ErrorNode) Children() []Node {
	nodes := make([]Node, len(e.Values))
	for i, value := range e.Values {
		nodes[i] = value
	}
	return nodes
}

type ErrorValueNode struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Code int    `json:"code"`
}

func (ev *ErrorValueNode) Children() []Node {
	// Error value nodes don't have any child nodes
	return nil
}
