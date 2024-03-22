package service

import "log/slog"

type StorageInterface interface{}

type Services struct {
	CardCheck
}

func New(log *slog.Logger) *Services {
	return &Services{
		CardCheck{log: log},
	}
}
