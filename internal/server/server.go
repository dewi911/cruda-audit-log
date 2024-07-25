package server

import (
	"cruda-audit-log/internal/domain"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	grpcSrv     *grpc.Server
	auditServer domain.AuditServiceServer
}

func New(auditServer domain.AuditServiceServer) *Server {
	return &Server{
		grpcSrv:     grpc.NewServer(),
		auditServer: auditServer,
	}
}

func (s *Server) ListenAndServe(port string) error {
	addr := fmt.Sprintf(":%s", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	domain.RegisterAuditServiceServer(s.grpcSrv, s.auditServer)

	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return nil
}
