package ast

type StructMemberNode struct {
	Name string
	Type string
}

func (s *StructMemberNode) Children() []Node {
	return nil
}

type StructNode struct {
	Name    string
	Members []*StructMemberNode
}

func (s *StructNode) AddMember(member *StructMemberNode) {
	s.Members = append(s.Members, member)
}

func (s *StructNode) Children() []Node {
	children := make([]Node, 0)
	for _, member := range s.Members {
		children = append(children, member)
	}
	return children
}
