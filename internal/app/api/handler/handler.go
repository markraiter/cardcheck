package handler

import (
	"log/slog"
)

type ServiceInterface interface {
}

type Handler struct {
	HealthCheck
}

func New(services ServiceInterface, log *slog.Logger) *Handler {
	return &Handler{
		HealthCheck{log: log},
	}
}
