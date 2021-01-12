package handler

import (
	"dmb-auth-service/service"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofiber/fiber/v2"
)

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

func (req ResetPasswordRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, is.Email),
	)
}

func ResetPassword(ctx *fiber.Ctx) error {
	var req ResetPasswordRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(fiber.StatusBadGateway)
	}
	if errors := req.Validate(); errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	_, _ = service.GetUserByEmail(req.Email)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Password reset link sent. Please check your email inbox/spam folder.",
		"data":    nil,
	})
}
