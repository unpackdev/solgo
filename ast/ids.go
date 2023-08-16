package ast

import "sync/atomic"

// GetNextID generates the next unique identifier for nodes in the abstract syntax tree (AST).
// It uses an atomic operation to ensure thread safety.
func (b *ASTBuilder) GetNextID() int64 {
	// Increment the value of b.nextID atomically and then subtract 1 to get the next unique ID.
	return atomic.AddInt64(&b.nextID, 1) - 1
}
