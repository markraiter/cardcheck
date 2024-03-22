// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Mark Raiter",
            "email": "raitermark@proton.me"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/check": {
            "post": {
                "description": "Validate card - check if card number is valid and expiration date is not in the past",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "check"
                ],
                "summary": "Validate card",
                "parameters": [
                    {
                        "description": "Card to validate",
                        "name": "card",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Card"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseW"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Ping health of API for Docker.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Shows the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Card": {
            "type": "object",
            "required": [
                "card_number",
                "expiration_month",
                "expiration_year"
            ],
            "properties": {
                "card_number": {
                    "type": "string",
                    "example": "1234567890123456"
                },
                "expiration_month": {
                    "type": "string",
                    "example": "12"
                },
                "expiration_year": {
                    "type": "string",
                    "example": "2028"
                },
                "id": {
                    "type": "string",
                    "example": ""
                }
            }
        },
        "model.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "001"
                },
                "message": {
                    "type": "string",
                    "example": "error message"
                }
            }
        },
        "model.ResponseMessage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "response message"
                }
            }
        },
        "model.ResponseW": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/model.Error"
                },
                "valid": {
                    "type": "boolean",
                    "example": true
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Cardcheck API",
	Description:      "This is an API for validating credit cards.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
