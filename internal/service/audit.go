package service

import (
	"context"
	"cruda-audit-log/internal/domain"
	"cruda-audit-log/pkg/models/audit"
)

type Repository interface {
	Insert(ctx context.Context, item domain.LogItem) error
}

type Audit struct {
	repo Repository
}

func NewAudit(repo Repository) *Audit {
	return &Audit{
		repo: repo,
	}
}

func (s *Audit) Insert(ctx context.Context, req *audit.LogRequest) error {
	item := domain.LogItem{
		Action:    req.GetEntity().String(),
		Entity:    req.GetEntity().String(),
		EntityID:  req.GetEntityId(),
		Timestamp: req.GetTimestamp().AsTime(),
	}

	return s.repo.Insert(ctx, item)
}
