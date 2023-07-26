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

	for _, node := range b.currentEnums {
		enumNode := node.(*EnumDefinition)
		if enumNode.GetName() == name {
			return node, node.GetTypeDescription()
		}
	}

	for _, node := range b.currentStructs {
		structNode := node.(*StructDefinition)
		if structNode.GetName() == name {
			return node, node.GetTypeDescription()
		}

		for _, member := range structNode.Members {
			if member.GetName() == name {
				return node, node.GetTypeDescription()
			}
		}
	}

	for _, node := range b.currentErrors {
		errorNode := node.(*ErrorDefinition)
		if errorNode.GetName() == name {
			return node, node.GetTypeDescription()
		}

		for _, member := range errorNode.Parameters.Parameters {
			if member.GetName() == name {
				return node, node.GetTypeDescription()
			}
		}
	}

	return nil, nil
}
