package ast

type NodeWithParameters interface {
	GetName() string
	GetNodes() []Node[NodeType]
	GetTypeDescription() *TypeDescription
	GetParameters() *ParameterList
}

func checkParameters(node NodeWithParameters, name string) bool {
	for _, parameter := range node.GetParameters().Parameters {
		if parameter.Name == name {
			return true
		}
	}
	return false
}

func discoverReferenceByCtxName(b *ASTBuilder, name string) (Node[NodeType], *TypeDescription) {
	for _, node := range b.sourceUnits {
		if node.GetName() == name {
			return node, node.GetTypeDescription()
		}

		for _, subNode := range node.GetNodes() {
			switch nodeCtx := subNode.(type) {
			case NodeWithParameters:
				if nodeCtx.GetName() == name || checkParameters(nodeCtx, name) {
					return subNode, subNode.GetTypeDescription()
				}
			}
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

	for _, modifier := range b.currentModifiers {
		modifierNode := modifier.(*ModifierDefinition)
		if modifierNode.GetName() == name {
			return modifier, modifier.GetTypeDescription()
		}

		for _, parameter := range modifierNode.Parameters.Parameters {
			if parameter.GetName() == name {
				return modifier, modifier.GetTypeDescription()
			}
		}
	}

	return nil, nil
}
