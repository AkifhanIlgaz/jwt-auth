package handler

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
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

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": t,
		})
	}

	return c.SendStatus(fiber.StatusUnauthorized)
}

func (h *Handler) Private(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	fmt.Println(user)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}
