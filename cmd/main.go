package main

import (
	"cruda-audit-log/internal/repository"
	"cruda-audit-log/internal/server"
	"cruda-audit-log/internal/service"
	"fmt"
	"log"
	"time"
)

func main() {

	auditRepo := &repository.Audit{}
	auditService := service.NewAudit(auditRepo)
	auditSrv := server.NewAuditServer(auditService)
	srv := server.New(auditSrv)

	fmt.Println("starting server...", time.Now())

	if err := srv.ListenAndServe("9000"); err != nil {
		log.Fatal(err)
	}
}
