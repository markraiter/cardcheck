package service

import (
	"log/slog"
	"os"
	"testing"

	"github.com/markraiter/cardcheck/internal/model"
)

func Test_Validate(t *testing.T) {
	cc := &CardCheck{
		log: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
	}

	tests := []struct {
		name      string
		card      *model.Card
		errorCode string
	}{
		{
			name: "valid card",
			card: &model.Card{
				CardNumber:      "5167803252097675",
				ExpirationMonth: "12",
				ExpirationYear:  "2024",
			},
			errorCode: "001",
		},
		{
			name: "invalid card number",
			card: &model.Card{
				CardNumber:      "1234",
				ExpirationMonth: "12",
				ExpirationYear:  "2024",
			},
			errorCode: "002",
		},
		{
			name: "invalid expiration month",
			card: &model.Card{
				CardNumber:      "5167803252097675",
				ExpirationMonth: "1",
				ExpirationYear:  "2024",
			},
			errorCode: "003",
		},
		{
			name: "invalid expiration month > 12",
			card: &model.Card{
				CardNumber:      "5167803252097675",
				ExpirationMonth: "13",
				ExpirationYear:  "2024",
			},
			errorCode: "003",
		},
		{
			name: "invalid expiration year",
			card: &model.Card{
				CardNumber:      "5167803252097675",
				ExpirationMonth: "12",
				ExpirationYear:  "2020",
			},
			errorCode: "003",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := cc.Validate(tt.card)

			if result.Error.Code != tt.errorCode {
				t.Errorf("expected error code %s, got %s", tt.errorCode, result.Error.Code)
			}
		})
	}
}
