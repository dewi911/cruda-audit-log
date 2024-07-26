package repository

import (
	"context"
	audit "cruda-audit-log/pkg/domain"
	"fmt"
)

type Audit struct {
}

func (r *Audit) Insert(ctx context.Context, item audit.LogItem) error {
	fmt.Printf("INSERTED %+v", item)
	return nil
}
