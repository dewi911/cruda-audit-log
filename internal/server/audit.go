package server

import (
	"context"
	"cruda-audit-log/pkg/models/audit"
)

type AuditService interface {
	Insert(ctx context.Context, req *audit.LogRequest) error
}

type AuditServer struct {
	audit.UnimplementedAuditServiceServer
	service AuditService
}

func NewAuditServer(service AuditService) *AuditServer {
	return &AuditServer{
		service: service,
	}
}

func (h *AuditServer) Log(ctx context.Context, req *audit.LogRequest) (*audit.Empty, error) {
	err := h.service.Insert(ctx, req)

	return &audit.Empty{}, err
}
