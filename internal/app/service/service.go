package service

import "log/slog"

type Services struct {
	CardCheck
}

func New(log *slog.Logger) *Services {
	return &Services{
		CardCheck{log: log},
	}
}
