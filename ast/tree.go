package ast

import (
	"go.uber.org/zap"
)

type Tree struct {
	*ASTBuilder
}

func NewTree(b *ASTBuilder) *Tree {
	return &Tree{
		ASTBuilder: b,
	}
}

func (t *Tree) GetById(id int64) Node[NodeType] {
	for _, node := range t.astRoot.GetNodes() {
		if node.GetId() == id {
			return node
		}

		if n := t.byRecursiveId(node, id); n != nil {
			return n
		}
	}
	return nil
}

func (t *Tree) UpdateNodeReferenceById(nodeId int64, nodeRefId int64, typeRef *TypeDescription) bool {
	if nodeId == 0 || nodeRefId == 0 {
		zap.L().Warn(
			"Invalid arguments provided to UpdateNodeReferenceId",
			zap.Int64("nodeId", nodeId),
			zap.Int64("node_ref_id", nodeRefId),
			zap.Any("type_ref", typeRef),
		)
		return false
	}

	for _, child := range t.astRoot.GetNodes() {
		if n := t.byRecursiveReferenceUpdate(child, nodeId, nodeRefId, typeRef); n {
			return n
		}
	}

	return false
}

func (t *Tree) byRecursiveReferenceUpdate(child Node[NodeType], nodeId int64, nodeRefId int64, typeRef *TypeDescription) bool {
	if child.GetId() == nodeId {
		child.SetReferenceDescriptor(nodeRefId, typeRef)
		return true
	}

	for _, c := range child.GetNodes() {
		if n := t.byRecursiveReferenceUpdate(c, nodeId, nodeRefId, typeRef); n {
			return n
		}
	}

	return false
}

func (t *Tree) byRecursiveId(node Node[NodeType], id int64) Node[NodeType] {
	if node.GetId() == id {
		return node
	}

	for _, child := range node.GetNodes() {
		if n := t.byRecursiveId(child, id); n != nil {
			return n
		}
	}

	return nil
}
