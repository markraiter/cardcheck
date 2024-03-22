package model

type ResponseMessage struct {
	Message string `json:"message" example:"response message"`
}

type ResponseW struct {
	Valid bool  `json:"valid" example:"true"`
	Error Error `json:"error"`
}

type Error struct {
	Code    string `json:"code" example:"001"`
	Message string `json:"message" example:"error message"`
}
