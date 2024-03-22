package service

import (
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/markraiter/cardcheck/internal/model"
)

type CardCheck struct {
	log *slog.Logger
}

func (cc *CardCheck) Validate(card *model.Card) (*model.ResponseW, error) {
	const operation = "service.CardCheck.Validate"

	log := cc.log.With(slog.String("operation", operation))

	log.Info("attempting to validate card")

	response := model.ResponseW{
		Valid: true,
		Error: model.Error{
			Code:    "001",
			Message: "card is valid",
		},
	}

	expMonth, err := strconv.Atoi(card.ExpirationMonth)
	if err != nil {
		log.Error("invalid expiration month")

		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	expYear, err := strconv.Atoi(card.ExpirationYear)
	if err != nil {
		log.Error("invalid expiration year")

		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	if !isValidCardNumber(card.CardNumber) {
		log.Error("invalid card number")

		response.Valid = false
		response.Error = model.Error{Code: "002", Message: "invalid card number"}

		return &response, nil
	}

	if !isValidExpirationDate(expMonth, expYear) {
		log.Error("invalid expiration date")

		response.Valid = false
		response.Error = model.Error{Code: "003", Message: "invalid expiration date"}

		return &response, nil
	}

	log.Info("card has been validated")

	return &response, nil
}

func isValidCardNumber(cardNumber string) bool {
	if _, err := strconv.Atoi(cardNumber); err != nil || len(cardNumber) < 12 || len(cardNumber) > 19 {
		return false
	}

	sum := 0
	isSecondDigit := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')

		if isSecondDigit {
			digit *= 2

			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		isSecondDigit = !isSecondDigit
	}

	return sum%10 == 0
}

func isValidExpirationDate(expirationMonth, expirationYear int) bool {
	currentYear, currentMonth, _ := time.Now().Date()

	if expirationYear < int(currentYear) || expirationYear == int(currentYear) && expirationMonth < int(currentMonth) {
		return false
	}

	return true
}
