package ast

import (
	"fmt"
	"reflect"

	ast_pb "github.com/txpull/protos/dist/go/ast"
)

// Resolver is a structure that helps in resolving the nodes of an Abstract Syntax Tree (AST).
// It contains a reference to an ASTBuilder and a map of UnprocessedNodes.
type Resolver struct {
	*ASTBuilder

	// Nodes that could not be processed while parsing AST.
	// This will resolve issues with forward referencing...
	UnprocessedNodes map[int64]UnprocessedNode

	// Temporary storage for already discovered targets to be able resolve forward references
	// instead of looping again through the whole AST and searching for the target that we already know.
	discoveredTargets map[string]Node[NodeType]
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
		ASTBuilder:        builder,
		UnprocessedNodes:  make(map[int64]UnprocessedNode, 0),
		discoveredTargets: make(map[string]Node[NodeType], 0),
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
func (r *Resolver) ResolveByNode(node Node[NodeType], name string) (int64, *TypeDescription) {
	rNode, rNodeType := r.resolveByNode(name, node)

	// Node could not be found in this moment, we are going to see if we can discover it in the
	// future at the end of whole parsing process.
	if rNodeType == nil {
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
func (r *Resolver) resolveByNode(name string, baseNode Node[NodeType]) (int64, *TypeDescription) {
	if node, nodeType := r.bySourceUnit(name); nodeType != nil {
		return node, nodeType
	}

	if node, nodeType := r.byFunction(name); nodeType != nil {
		return node, nodeType
	}

	if node, nodeType := r.byStateVariables(name); nodeType != nil {
		return node, nodeType
	}

	if node, nodeType := r.byVariables(name); nodeType != nil {
		return node, nodeType
	}

	if node, nodeType := r.byEvents(name); nodeType != nil {
		return node, nodeType
	}

	if node, nodeType := r.byEnums(name); nodeType != nil {
		return node, nodeType
	}

	if node, nodeType := r.byStructs(name); nodeType != nil {
		return node, nodeType
	}

	if node, nodeType := r.byErrors(name); nodeType != nil {
		return node, nodeType
	}

	if node, nodeType := r.byModifiers(name); nodeType != nil {
		return node, nodeType
	}

	return 0, nil
}

// Resolve attempts to resolve all UnprocessedNodes in the Resolver and sets the entry source unit for the AST.
// It updates the node references in the AST and removes the nodes from the UnprocessedNodes map once they are resolved.
// If a node cannot be resolved, it is left in the UnprocessedNodes map for future resolution.
func (r *Resolver) Resolve() []error {
	// Resolve all source unit symbols that are not resolved yet.
	r.resolveExportedSymbols()

	// In case that entry source unit is not set, we are going to set it now.
	// Do some kumbaya to figure it out...
	r.resolveEntrySourceUnit()

	// There can be multiple errors happening from attempt to resolve unprocessed events.
	// Instead of returning just one, we're going to ensure all are returned at the same time.
	var errors []error

	// Now this is a hack, but working hack so we're going to go with it...
	// Reference update can come in any direction really. For example B that uses A can come first
	// and because of it, it B will never discover A. In order to ensure that is not the case here,
	// we are going to iterate few times through unprocessed references...
	for i := 0; i <= 1; i++ {
		for nodeId, node := range r.UnprocessedNodes {
			if rNodeId, rNodeType := r.resolveByNode(node.Name, node.Node); rNodeType != nil {
				if updated := r.tree.UpdateNodeReferenceById(nodeId, rNodeId, rNodeType); updated {
					delete(r.UnprocessedNodes, nodeId)
				} else {
					errors = append(
						errors,
						fmt.Errorf(
							"unable to update node reference by id %d - name: %s - type: %v",
							nodeId, node.Name, rNodeType,
						),
					)
				}
			}
		}

		// No need to go through the process again if we already have all of the nodes resolved...
		if len(r.UnprocessedNodes) == 0 {
			break
		}
	}

	for nodeId, node := range r.UnprocessedNodes {
		errors = append(
			errors,
			fmt.Errorf(
				"unable to resolve node by id %d - name: %s - type: %v",
				nodeId, node.Name, reflect.TypeOf(node.Node),
			),
		)
	}

	return errors
}

func (r *Resolver) resolveExportedSymbols() {
	for _, sourceNode := range r.sourceUnits {

		// In case any imports are available and they are not exported
		// we are going to append them to the exported symbols.
		for _, node := range sourceNode.GetNodes() {
			if node.GetType() == ast_pb.NodeType_IMPORT_DIRECTIVE {
				importNode := node.(*Import)
				if !r.symbolExists(importNode.GetName(), sourceNode.GetExportedSymbols()) {
					sourceNode.ExportedSymbols = append(
						sourceNode.ExportedSymbols,
						Symbol{
							Id:           importNode.GetId(),
							Name:         importNode.GetName(),
							AbsolutePath: importNode.GetAbsolutePath(),
						},
					)
				}
			}
		}

		// Base contracts will be available if contract or interface inherits any of the
		// clases or interfaces.
		for _, baseContract := range sourceNode.GetBaseContracts() {
			if !r.symbolExists(baseContract.GetBaseName().GetName(), sourceNode.GetExportedSymbols()) {
				sourceNode.ExportedSymbols = append(
					sourceNode.ExportedSymbols,
					Symbol{
						Id:   baseContract.GetId(),
						Name: baseContract.GetBaseName().GetName(),
					},
				)
			}
		}
	}
}

func (r *Resolver) symbolExists(name string, symbols []Symbol) bool {
	for _, symbol := range symbols {
		if symbol.GetName() == name {
			return true
		}
	}

	return false
}

func (r *Resolver) resolveEntrySourceUnit() {
	// Entry source unit is already calculated, we are going to skip the check.
	// Note if you are reading this, it's the best practice to always set it as
	// margin for potential errors is minimized.
	if r.tree.astRoot.EntrySourceUnit != 0 {
		return
	}

	var entrySourceUnit int64
	for _, node := range r.sourceUnits {
		for _, entry := range node.GetExportedSymbols() {
			if len(r.sources.EntrySourceUnitName) > 0 &&
				r.sources.EntrySourceUnitName == entry.GetName() {
				r.tree.astRoot.SetEntrySourceUnit(entry.GetId())
				return
			}

			// We should look for the highest amount of exported symbols and then
			// take that source unit as entry source unit.
			// This is not the best solution, that is for sure...
			// Or go with the highest source unit id, that should work as well?
			// Well none of these should work...
			// Probably the best one would be the one that is not imported by any other
			// source unit and has the highest amount of exported symbols and is contract type.
			if entry.GetId() > entrySourceUnit {
				entrySourceUnit = entry.GetId()
			}
		}
	}

	r.tree.astRoot.SetEntrySourceUnit(entrySourceUnit)
}

func (r *Resolver) bySourceUnit(name string) (int64, *TypeDescription) {
	for _, node := range r.sourceUnits {
		if node.GetName() == name {
			return node.GetId(), node.GetTypeDescription()
		}
	}

	return 0, nil
}

func (r *Resolver) byStateVariables(name string) (int64, *TypeDescription) {
	for _, node := range r.currentStateVariables {
		if node.GetName() == name {
			// It could be that node got updated in the mean time so we're going to
			// look up for the node and it's type description...
			if node.GetTypeDescription() == nil {
				if target := r.tree.GetById(node.GetId()); target != nil {
					return target.GetId(), target.GetTypeDescription()
				}
			}
			return node.GetId(), node.GetTypeDescription()
		}
	}

	return 0, nil
}

func (r *Resolver) byVariables(name string) (int64, *TypeDescription) {
	for _, node := range r.currentVariables {
		variable := node.(*VariableDeclaration)

		for _, declaration := range variable.Declarations {
			if declaration.GetName() == name {
				return node.GetId(), declaration.GetTypeDescription()
			}

		}
	}

	return 0, nil
}

func (r *Resolver) byEvents(name string) (int64, *TypeDescription) {
	for _, node := range r.currentEvents {
		eventNode := node.(*EventDefinition)
		if eventNode.GetName() == name {
			return node.GetId(), node.GetTypeDescription()
		}
	}

	return 0, nil
}

func (r *Resolver) byEnums(name string) (int64, *TypeDescription) {
	for _, node := range r.currentEnums {
		enumNode := node.(*EnumDefinition)
		if enumNode.GetName() == name {
			return node.GetId(), node.GetTypeDescription()
		}

		for _, member := range enumNode.GetMembers() {
			if member.GetName() == name {
				return node.GetId(), node.GetTypeDescription()
			}
		}
	}

	return 0, nil
}

func (r *Resolver) byStructs(name string) (int64, *TypeDescription) {
	for _, node := range r.currentStructs {
		structNode := node.(*StructDefinition)
		if structNode.GetName() == name {
			return node.GetId(), node.GetTypeDescription()
		}

		for _, member := range structNode.Members {
			if member.GetName() == name {
				return node.GetId(), node.GetTypeDescription()
			}
		}
	}

	return 0, nil
}

func (r *Resolver) byErrors(name string) (int64, *TypeDescription) {
	for _, node := range r.currentErrors {
		errorNode := node.(*ErrorDefinition)
		if errorNode.GetName() == name {
			return node.GetId(), node.GetTypeDescription()
		}

		for _, member := range errorNode.Parameters.Parameters {
			if member.GetName() == name {
				return node.GetId(), node.GetTypeDescription()
			}
		}
	}

	return 0, nil
}

func (r *Resolver) byModifiers(name string) (int64, *TypeDescription) {
	for _, modifier := range r.currentModifiers {
		modifierNode := modifier.(*ModifierDefinition)
		if modifierNode.GetName() == name {
			return modifier.GetId(), modifier.GetTypeDescription()
		}

		for _, parameter := range modifierNode.GetParameters().GetParameters() {
			if parameter.GetName() == name {
				return parameter.GetId(), parameter.GetTypeDescription()
			}
		}
	}

	return 0, nil
}

func (r *Resolver) byFunction(name string) (int64, *TypeDescription) {
	for _, node := range r.currentFunctions {
		switch nodeCtx := node.(type) {
		case *Constructor:
			if nodeCtx.GetNodes() != nil {
				for _, member := range nodeCtx.GetNodes() {
					if recNode, recNodeType := r.byRecursiveSearch(member, name); recNode != nil {
						return recNode.GetId(), recNodeType
					}
				}
			}

		case *Function:
			if nodeCtx.GetName() == name {
				return node.GetId(), node.GetTypeDescription()
			}

			if nodeCtx.GetNodes() != nil {
				for _, member := range nodeCtx.GetNodes() {
					if recNode, recNodeType := r.byRecursiveSearch(member, name); recNode != nil {
						return recNode.GetId(), recNodeType
					}
				}
			}
		}
	}

	return 0, nil
}

// byRecursiveSearch is a helper function that attempts to resolve a node by its name by recursively searching the node's children.
// It returns the resolved Node and its TypeDescription, or nil if the node cannot be found.
func (r *Resolver) byRecursiveSearch(node Node[NodeType], name string) (Node[NodeType], *TypeDescription) {
	if node == nil || node.GetNodes() == nil {
		return nil, nil
	}

	if node.GetType() == ast_pb.NodeType_DO_WHILE_STATEMENT {
		nodeCtx := node.(*DoWhileStatement)
		for _, condition := range nodeCtx.GetCondition().GetNodes() {
			if primary, ok := condition.(*PrimaryExpression); ok {
				if primary.GetName() == name {
					return primary, primary.GetTypeDescription()
				}
			}
		}
	}

	if node.GetType() == ast_pb.NodeType_FOR_STATEMENT {
		nodeCtx := node.(*ForStatement)
		for _, condition := range nodeCtx.GetCondition().GetNodes() {
			if primary, ok := condition.(*PrimaryExpression); ok {
				if primary.GetName() == name {
					return primary, primary.GetTypeDescription()
				}
			}
		}
	}

	if node.GetType() == ast_pb.NodeType_IDENTIFIER {
		nodeCtx := node.(*PrimaryExpression)
		if nodeCtx.GetName() == name {
			return nodeCtx, nodeCtx.GetTypeDescription()
		}
	}

	for _, n := range node.GetNodes() {
		if n == nil {
			continue
		}

		// Needs to be here as there are no parent nodes available...
		if n.GetType() == ast_pb.NodeType_IDENTIFIER {
			nodeCtx := n.(*PrimaryExpression)
			if nodeCtx.GetName() == name {
				return nodeCtx, nodeCtx.GetTypeDescription()
			}
		}

		if node, nodeType := r.byRecursiveSearch(n, name); node != nil && nodeType != nil {
			return node, nodeType
		}
	}

	return nil, nil
}

/* func (r *Resolver) hasMember(node interface{}, member string) bool {
	val := reflect.ValueOf(node)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return false
	}

	typeOfInput := val.Type()
	for i := 0; i < val.NumField(); i++ {
		if typeOfInput.Field(i).Name == member {
			return true
		}
	}

	field := val.FieldByName(member)
	return field.IsValid()
}
*/
