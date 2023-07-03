package ast

type EventParameterNode struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Indexed bool   `json:"indexed"`
}

func (e *EventParameterNode) Children() []Node {
	return nil
}

type EventNode struct {
	Name       string                `json:"name"`
	Anonymous  bool                  `json:"anonymous"`
	Parameters []*EventParameterNode `json:"parameters"`
}

func (e *EventNode) Children() []Node {
	nodes := make([]Node, len(e.Parameters))
	for i, parameter := range e.Parameters {
		nodes[i] = parameter
	}
	return nodes
}
