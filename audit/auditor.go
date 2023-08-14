package audit

import "context"

type Auditor struct {
	ctx context.Context
}

func NewAuditor(ctx context.Context) *Auditor {
	return &Auditor{
		ctx: ctx,
	}
}
