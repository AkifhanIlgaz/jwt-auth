package main

import (
	"github.com/AkifhanIlgaz/jwt-auth/config"
	"github.com/AkifhanIlgaz/jwt-auth/handler"
	"github.com/AkifhanIlgaz/jwt-auth/middlewares"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	jwt := middlewares.NewAuthMiddleware(config.Secret)

	app.Post("/login", handler.Login)
	app.Get("/protected", jwt, handler.Protected)

	app.Listen(":3000")
}
