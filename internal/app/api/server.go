package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/markraiter/cardcheck/config"
)

const (
	bodyLimit = 100 * 1024 * 1024 // 100 MB RESTRICION
)

type Server struct {
	HTTPServer *fiber.App
}

// New returns new instance of the Server.
func New(cfg *config.Config, handler *handler.Handler) *Server {
	server := new(Server)

	fconfig := fiber.Config{
		ReadTimeout:  cfg.Server.AppReadTimeout,
		WriteTimeout: cfg.Server.AppWriteTimeout,
		IdleTimeout:  cfg.Server.AppIdleTimeout,
		BodyLimit:    bodyLimit,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var localError *fiber.Error
			if errors.As(err, &localError) {
				code = localError.Code
			}

			c.Status(code)

			if err := c.JSON(model.Response{Message: localError.Message}); err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, err.Error())
			}

			return nil
		},
	}
	server.HTTPServer = fiber.New(fconfig)

	server.HTTPServer.Use(recover.New())

	server.HTTPServer.Use(logger.New())

	server.initRoutes(server.HTTPServer, handler, cfg)

	return server
}

func (s *Server) Shutdown(ctx context.Context) error {
	const op = "api.Server.Shutdown"

	return fmt.Errorf("%s: %w", op, s.HTTPServer.ShutdownWithContext(ctx))
}
