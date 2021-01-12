package handler

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
)

type RenewTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (req RenewTokenRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.RefreshToken, validation.Required),
	)
}

func RenewToken(ctx *fiber.Ctx) error {
	var req RenewTokenRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(fiber.StatusBadGateway)
	}
	if errors := req.Validate(); errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return ctx.SendStatus(fiber.StatusOK)
}