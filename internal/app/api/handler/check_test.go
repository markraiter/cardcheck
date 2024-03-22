package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/markraiter/cardcheck/internal/model"
)

func Test_CheckCard_Validate(t *testing.T) {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	reqValidator := validator.New()

	tests := []struct {
		name           string
		card           *model.Card
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "valid request",
			card: &model.Card{
				CardNumber:      "5167803252097675",
				ExpirationMonth: "12",
				ExpirationYear:  "2024",
			},
			expectedStatus: fiber.StatusOK,
			expectedBody:   `{"valid":true,"error":{"code":"","message":""}}`,
		},
		{
			name: "no card number",
			card: &model.Card{
				CardNumber:      "",
				ExpirationMonth: "12",
				ExpirationYear:  "2024",
			},
			expectedStatus: fiber.StatusBadRequest,
			expectedBody:   `{"message":"Key: 'Card.CardNumber' Error:Field validation for 'CardNumber' failed on the 'required' tag"}`,
		},
		{
			name: "no expiration month",
			card: &model.Card{
				CardNumber:      "5167803252097675",
				ExpirationMonth: "",
				ExpirationYear:  "2024",
			},
			expectedStatus: fiber.StatusBadRequest,
			expectedBody:   `{"message":"Key: 'Card.ExpirationMonth' Error:Field validation for 'ExpirationMonth' failed on the 'required' tag"}`,
		},
		{
			name: "no expiration year",
			card: &model.Card{
				CardNumber:      "5167803252097675",
				ExpirationMonth: "12",
				ExpirationYear:  "",
			},
			expectedStatus: fiber.StatusBadRequest,
			expectedBody:   `{"message":"Key: 'Card.ExpirationYear' Error:Field validation for 'ExpirationYear' failed on the 'required' tag"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()

			cc := &CheckCard{
				log: log,
				validator: mockValidator{
					card: tt.card,
				},
				reqValidator: reqValidator,
			}

			app.Post("/check", cc.Validate)

			body, _ := json.Marshal(tt.card)

			req := httptest.NewRequest("POST", "/check", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req)

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			responseBody, _ := io.ReadAll(resp.Body)

			if string(responseBody) != tt.expectedBody {
				t.Errorf("expected body %s, got %s", tt.expectedBody, string(responseBody))
			}
		})
	}
}

type mockValidator struct {
	card *model.Card
}

func (v mockValidator) Validate(card *model.Card) (*model.ResponseW, error) {
	if card.CardNumber == v.card.CardNumber {
		return &model.ResponseW{Valid: true}, nil
	}

	return nil, errors.New("Invalid card number")
}
