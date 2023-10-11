package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type Handler struct {
	x bool
}

func (h *Handler) Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// TODO: Check if the user exists
	if username == "jon" && password == "password" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name":  "Jon Doe",
			"admin": true,
			"exp":   time.Now().Add(72 * time.Hour).Unix(),
		})

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(map[string]string{
			"token": t,
		})
	}

	return c.Status(fiber.ErrUnauthorized.Code).JSON(map[string]string{
		"message": fiber.ErrUnauthorized.Message,
	})
}

func main() {
	router := fiber.New()
	handler := Handler{}

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("Hello World")
	})

	router.Post("/login", handler.Login)

	router.Server().Logger.Printf("%v", router.Listen(":1323"))

}
