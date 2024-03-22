package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/markraiter/cardcheck/internal/model"
)

type Validator interface {
	Validate(card *model.Card) (*model.ResponseW, error)
}

type CheckCard struct {
	log       *slog.Logger
	validator Validator
}

// @Summary Validate card
// @Description Validate card - check if card number is valid and expiration date is not in the past
// @Tags check
// @Accept json
// @Produce json
// @Param card body model.Card true "Card to validate"
// @Success 200 {object} model.ResponseW
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /check [post]
func (cc *CheckCard) Validate(c *fiber.Ctx) error {
	const operation = "handler.CheckCard.Validate"

	log := cc.log.With(slog.String("operation", operation))

	var card model.Card

	if err := c.BodyParser(&card); err != nil {
		log.Error("error parsing request body", model.Err(err))

		return c.Status(fiber.StatusBadRequest).JSON(model.ResponseMessage{Message: err.Error()})
	}

	result, err := cc.validator.Validate(&card)
	if err != nil {
		log.Error("error validating card", model.Err(err))

		return c.Status(fiber.StatusInternalServerError).JSON(model.ResponseMessage{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
