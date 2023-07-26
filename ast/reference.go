package ast

func discoverReferenceByCtxName(b *ASTBuilder, name string) (Node[NodeType], *TypeDescription) {
	for _, node := range b.currentStateVariables {
		if node.GetName() == name {
			return node, node.TypeDescription
		}
	}

	return nil, nil
}
