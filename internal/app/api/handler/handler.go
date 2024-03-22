package handler

import (
	"log/slog"
)

type ServiceInterface interface {
	Validator
}

type Handler struct {
	HealthCheck
	CheckCard
}

func New(services ServiceInterface, log *slog.Logger) *Handler {
	return &Handler{
		HealthCheck{log: log},
		CheckCard{log: log, validator: services},
	}
}
