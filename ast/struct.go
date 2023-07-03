package ast

// StructMemberNode represents a member of a struct.
type StructMemberNode struct {
	Name string
	Type string
}

// Children returns an empty slice of nodes.
func (s *StructMemberNode) Children() []Node {
	return nil
}

// StructNode represents a struct definition in Solidity.
type StructNode struct {
	Name    string
	Members []*StructMemberNode
}

// AddMember adds a member to the struct.
func (s *StructNode) AddMember(member *StructMemberNode) {
	s.Members = append(s.Members, member)
}

// Children returns an empty slice of nodes.
func (s *StructNode) Children() []Node {
	children := make([]Node, 0)
	for _, member := range s.Members {
		children = append(children, member)
	}
	return children
}
