package handlers

import (
	"dmb-auth-service/global"
	"dmb-auth-service/models"
	"dmb-auth-service/services"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(ctx *fiber.Ctx) error {
	var (
		input models.LoginInput
		user  models.User
	)
	err := ctx.BodyParser(&input)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadGateway)
	}

	validationErrors := input.Validate()
	if validationErrors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}

	user, _ = services.GetUserByEmail(input.Email)

	if !CheckPasswordHash(input.Password, user.Password) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Bad Credentials",
			"data":    nil,
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	jwtToken, err := token.SignedString([]byte(global.Config("JWT_SECRET")))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.JSON(fiber.Map{
		"message": "Login Successful",
		"data":    jwtToken,
	})
}
