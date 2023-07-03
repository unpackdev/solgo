package ast

type EnumMemberNode struct {
	Name string `json:"name"`
}

func (e *EnumMemberNode) Children() []Node {
	return nil
}

type EnumNode struct {
	Name         string            `json:"name"`         // Name of the enum
	MemberValues []*EnumMemberNode `json:"memberValues"` // Values of the enum members
}

func (e *EnumNode) Children() []Node {
	nodes := make([]Node, len(e.MemberValues))
	for i, member := range e.MemberValues {
		nodes[i] = member
	}
	return nodes
}
