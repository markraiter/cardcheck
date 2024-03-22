package model

type Card struct {
	CardNumber      string `json:"card_number" validate:"required" example:"1234567890123456"`
	ExpirationMonth string `json:"expiration_month" validate:"required" example:"12"`
	ExpirationYear  string `json:"expiration_year" validate:"required" example:"2028"`
}
