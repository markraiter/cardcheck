package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/swagger"
	"github.com/markraiter/cardcheck/internal/app/api/handler"
	"github.com/markraiter/cardcheck/internal/config"
)

const (
	apiPrefix = "/api"
	auth      = "/auth"
	report    = "/report"
)

func (s Server) initRoutes(app *fiber.App, handler *handler.Handler, cfg *config.Config) {

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get(apiPrefix+"/health", timeout.NewWithContext(handler.APIHealth, cfg.Server.AppReadTimeout))

	// api := app.Group(apiPrefix)
	// {

	// }
}
