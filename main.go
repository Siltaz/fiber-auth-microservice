package main

import (
	"dmb-auth-service/config"
	"dmb-auth-service/handler"
	"dmb-auth-service/middleware"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port    = config.Config("APP_PORT")
	prod, _ = strconv.ParseBool(config.Config("APP_PROD"))
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		ServerHeader:  "AuthServer",
		Prefork:       prod,
	})

	app.Use(logger.New())
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 20 * time.Second,
	}))
	app.Use(recover.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "POST",
	}))

	api := app.Group("/api/v1/auth")

	api.Post("/signup", handler.Signup)
	api.Post("/login", handler.Login)
	api.Post("/logout", middleware.ProtectedRoute(), handler.Logout)
	api.Post("/renew-token", handler.RenewToken)
	api.Post("/reset-password", handler.ResetPassword)

	app.Use(handler.NotFound)

	log.Fatal(app.Listen(":" + port))
}
