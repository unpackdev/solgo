package ast

func discoverReferenceByCtxName(b *ASTBuilder, name string) (Node[NodeType], *TypeDescription) {

	for _, node := range b.sourceUnits {
		if node.GetName() == name {
			return node, node.GetTypeDescription()
		}
	}

	for _, node := range b.currentStateVariables {
		if node.GetName() == name {
			return node, node.TypeDescription
		}
	}

	for _, node := range b.currentEvents {
		eventNode := node.(*EventDefinition)
		if eventNode.GetName() == name {
			return node, node.GetTypeDescription()
		}
	}

	return nil, nil
}
