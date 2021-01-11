package main

import (
	"dmb-auth-service/global"
	"dmb-auth-service/handlers"
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
	port    = global.Config("APP_PORT")
	prod, _ = strconv.ParseBool(global.Config("APP_PROD"))
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

	api.Post("/signup", handlers.Signup)
	api.Post("/login", handlers.Login)
	api.Post("/logout", middleware.Protected(), handlers.Logout)
	api.Post("/refreshToken", middleware.Protected(), handlers.RefreshToken)
	api.Post("/resetPassword", handlers.ResetPassword)

	app.Use(handlers.NotFound)

	log.Fatal(app.Listen(":" + port))
}
