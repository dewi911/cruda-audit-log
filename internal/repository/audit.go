package repository

import (
	"context"
	"cruda-audit-log/internal/domain"
	"fmt"
)

type Audit struct {
}

func (r *Audit) Insert(ctx context.Context, item domain.LogItem) error {
	fmt.Printf("INSERTED %+v", item)
	return nil
}
