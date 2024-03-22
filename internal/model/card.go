package model

type Card struct {
	ID              string `bson:"_id,omitempty" json:"id" validate:"omitempty" example:""`
	CardNumber      string `bson:"card_number" json:"card_number" validate:"required" example:"1234567890123456"`
	ExpirationMonth string `bson:"expiration_month" json:"expiration_month" validate:"required" example:"12"`
	ExpirationYear  string `bson:"expiration_year" json:"expiration_year" validate:"required" example:"2028"`
}
