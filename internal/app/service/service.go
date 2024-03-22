package service

import "log/slog"

type StorageInterface interface{}

type Services struct{}

func New(storage StorageInterface, log *slog.Logger) *Services {
	return &Services{}
}
