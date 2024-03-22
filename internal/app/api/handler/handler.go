package handler

import (
	"log/slog"

	"github.com/go-playground/validator"
)

type ServiceInterface interface {
	Validator
}

type Handler struct {
	HealthCheck
	CheckCard
}

func New(services ServiceInterface, reqValidator *validator.Validate, log *slog.Logger) *Handler {
	return &Handler{
		HealthCheck{log: log},
		CheckCard{log: log, validator: services, reqValidator: reqValidator},
	}
}
