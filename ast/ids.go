package ast

import "sync/atomic"

func (b *ASTBuilder) GetNextID() int64 {
	return atomic.AddInt64(&b.nextID, 1) - 1
}
