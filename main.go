package main

import (
	"github.com/AkifhanIlgaz/jwt-auth/handler"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var h handler.Handler

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("Hello World")
	})

	app.Post("/login", h.Login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	app.Get("/private", h.Private)

	app.Server().Logger.Printf("%v", app.Listen(":1323"))

}
