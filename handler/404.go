package handler

import (
	"github.com/gofiber/fiber/v2"
)

func NotFound(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusNotFound)
}
