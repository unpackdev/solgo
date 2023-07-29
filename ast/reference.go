package ast

// Resolver is a structure that helps in resolving the nodes of an Abstract Syntax Tree (AST).
// It contains a reference to an ASTBuilder and a map of UnprocessedNodes.
type Resolver struct {
	*ASTBuilder

	// Nodes that could not be processed while parsing AST.
	// This will resolve issues with forward referencing...
	UnprocessedNodes map[int64]UnprocessedNode
}

// UnprocessedNode is a structure that represents a node that could not be processed during the parsing of the AST.
type UnprocessedNode struct {
	Id   int64          `json:"id"`
	Name string         `json:"name"`
	Node Node[NodeType] `json:"ref"`
}

// NewResolver creates a new Resolver with the provided ASTBuilder and initializes the UnprocessedNodes map.
func NewResolver(builder *ASTBuilder) *Resolver {
	return &Resolver{
		ASTBuilder:       builder,
		UnprocessedNodes: make(map[int64]UnprocessedNode, 0),
	}
}

// GetUnprocessedNodes returns the map of UnprocessedNodes in the Resolver.
func (r *Resolver) GetUnprocessedNodes() map[int64]UnprocessedNode {
	return r.UnprocessedNodes
}

// GetUnprocessedCount returns the number of UnprocessedNodes in the Resolver.
func (r *Resolver) GetUnprocessedCount() int {
	return len(r.UnprocessedNodes)
}

// ResolveByNode attempts to resolve a node by its name and returns the resolved Node and its TypeDescription.
// If the node cannot be found, it is added to the UnprocessedNodes map for future resolution.
func (r *Resolver) ResolveByNode(node Node[NodeType], name string) (Node[NodeType], *TypeDescription) {
	rNode, rNodeType := r.resolveByNode(name, node)

	// Node could not be found in this moment, we are going to see if we can discover it in the
	// future at the end of whole parsing process.
	if rNode == nil && rNodeType == nil {
		r.UnprocessedNodes[node.GetId()] = UnprocessedNode{
			Id:   node.GetId(),
			Name: name,
			Node: node,
		}
	}

	return rNode, rNodeType
}

// resolveByNode is a helper function that attempts to resolve a node by its name by checking various node types.
// It returns the resolved Node and its TypeDescription, or nil if the node cannot be found.
func (r *Resolver) resolveByNode(name string, node Node[NodeType]) (Node[NodeType], *TypeDescription) {
	if node, nodeType := r.bySourceUnit(name); node != nil && nodeType != nil {
		return node, nodeType
	}

	if node, nodeType := r.byFunction(name); node != nil {
		return node, nodeType
	}

	if node, nodeType := r.byStateVariables(name); node != nil && nodeType != nil {
		return node, nodeType
	}

	if node, nodeType := r.byVariables(name); node != nil && nodeType != nil {
		return node, nodeType
	}

	if node, nodeType := r.byEvents(name); node != nil && nodeType != nil {
		return node, nodeType
	}

	if node, nodeType := r.byEnums(name); node != nil && nodeType != nil {
		return node, nodeType
	}

	if node, nodeType := r.byStructs(name); node != nil && nodeType != nil {
		return node, nodeType
	}

	if node, nodeType := r.byErrors(name); node != nil && nodeType != nil {
		return node, nodeType
	}

	if node, nodeType := r.byModifiers(name); node != nil && nodeType != nil {
		return node, nodeType
	}

	return nil, nil
}

// Resolve attempts to resolve all UnprocessedNodes in the Resolver and sets the entry source unit for the AST.
// It updates the node references in the AST and removes the nodes from the UnprocessedNodes map once they are resolved.
// If a node cannot be resolved, it is left in the UnprocessedNodes map for future resolution.
func (r *Resolver) Resolve() error {

	var entrySourceUnit int64
	for _, node := range r.sourceUnits {
		for _, entry := range node.GetExportedSymbols() {
			if entry.GetId() > entrySourceUnit {
				entrySourceUnit = entry.GetId()
			}
		}
	}
	r.tree.astRoot.SetEntrySourceUnit(entrySourceUnit)

	for nodeId, node := range r.UnprocessedNodes {
		if rNode, rNodeType := r.resolveByNode(node.Name, node.Node); rNode != nil {
			if updated := r.tree.UpdateNodeReferenceById(nodeId, rNode.GetId(), rNodeType); updated {
				delete(r.UnprocessedNodes, nodeId)
			}
		}
	}

	return nil
}

func (r *Resolver) bySourceUnit(name string) (Node[NodeType], *TypeDescription) {
	for _, node := range r.sourceUnits {
		if node.GetName() == name {
			return node, node.GetTypeDescription()
		}
	}

	return nil, nil
}

func (r *Resolver) byStateVariables(name string) (Node[NodeType], *TypeDescription) {
	for _, node := range r.currentStateVariables {
		if node.GetName() == name {
			return node, node.TypeDescription
		}
	}

	return nil, nil
}

func (r *Resolver) byVariables(name string) (Node[NodeType], *TypeDescription) {
	for _, node := range r.currentVariables {
		variable := node.(*VariableDeclaration)

		for _, declaration := range variable.Declarations {
			if declaration.GetName() == name {
				return node, declaration.GetTypeDescription()
			}
		}
	}

	return nil, nil
}

func (r *Resolver) byEvents(name string) (Node[NodeType], *TypeDescription) {
	for _, node := range r.currentEvents {
		eventNode := node.(*EventDefinition)
		if eventNode.GetName() == name {
			return node, node.GetTypeDescription()
		}
	}

	return nil, nil
}

func (r *Resolver) byEnums(name string) (Node[NodeType], *TypeDescription) {
	for _, node := range r.currentEnums {
		enumNode := node.(*EnumDefinition)
		if enumNode.GetName() == name {
			return node, node.GetTypeDescription()
		}
	}

	return nil, nil
}

func (r *Resolver) byStructs(name string) (Node[NodeType], *TypeDescription) {
	for _, node := range r.currentStructs {
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

	return nil, nil
}

func (r *Resolver) byErrors(name string) (Node[NodeType], *TypeDescription) {
	for _, node := range r.currentErrors {
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

func (r *Resolver) byModifiers(name string) (Node[NodeType], *TypeDescription) {
	for _, modifier := range r.currentModifiers {
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

func (r *Resolver) byFunction(name string) (Node[NodeType], *TypeDescription) {
	for _, node := range r.currentFunctions {
		switch nodeCtx := node.(type) {
		case *Constructor:
			if nodeCtx.GetNodes() != nil {
				for _, member := range nodeCtx.GetNodes() {
					if node, nodeType := r.byRecursiveSearch(member, name); node != nil && nodeType != nil {
						return node, nodeType
					}
				}
			}

		case Function:

			if nodeCtx.GetName() == name {
				return node, node.GetTypeDescription()
			}

			if nodeCtx.GetNodes() != nil {
				for _, member := range nodeCtx.GetNodes() {
					if node, nodeType := r.byRecursiveSearch(member, name); node != nil && nodeType != nil {
						return node, nodeType
					}
				}
			}
		}
	}

	return nil, nil
}

// byRecursiveSearch is a helper function that attempts to resolve a node by its name by recursively searching the node's children.
// It returns the resolved Node and its TypeDescription, or nil if the node cannot be found.
func (r *Resolver) byRecursiveSearch(node Node[NodeType], name string) (Node[NodeType], *TypeDescription) {
	return nil, nil
}
