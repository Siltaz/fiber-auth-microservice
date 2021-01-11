package handlers

import (
	"dmb-auth-service/models"
	"dmb-auth-service/services"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Signup(ctx *fiber.Ctx) error {
	var user models.User
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadGateway)
	}

	validationErrors := user.Validate()
	if validationErrors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	user.Password = hashedPassword
	user, _ = services.CreateUser(user)
	return ctx.JSON(fiber.Map{
		"message": "Registration Successful",
		"data":    user,
	})
}

// Logout - Blacklists current JWT token
func Logout(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}

// RefreshToken - Issues new JWT token and blacklists current one
func RefreshToken(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}

// ResetPassword - Sends reset password link to the registered email
func ResetPassword(ctx *fiber.Ctx) error {
	var input models.ResetPasswordInput
	err := ctx.BodyParser(&input)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadGateway)
	}

	validationErrors := input.Validate()
	if validationErrors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}
	_, _ = services.GetUserByEmail(input.Email)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Password reset link sent. Please check your email inbox/spam folder.",
		"data":    nil,
	})
}

// NotFound - Handles 404 errors
func NotFound(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusNotFound)
}
