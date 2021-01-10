package handlers

import (
	"dmb-auth-service/models"
	"dmb-auth-service/services"
	"github.com/gofiber/fiber/v2"
)

func Signup(ctx *fiber.Ctx) error {
	var user models.User
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	user, _ = services.CreateUser(user)
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func Login(ctx *fiber.Ctx) error {
	var (
		input models.LoginInput
		user  models.User
	)

	err := ctx.BodyParser(&input)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	user, _ = services.GetUserByEmail(input.Email)

	if input.Password != user.Password {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}

func Logout(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}

func RefreshToken(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}

func ResetPassword(ctx *fiber.Ctx) error {
	var (
		input models.ResetPasswordInput
	)

	err := ctx.BodyParser(&input)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	_, _ = services.GetUserByEmail(input.Email)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Password reset link sent to your email. Please check your inbox/spam folder.",
	})
}

func NotFound(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusNotFound)
}
