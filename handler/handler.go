package handler

import (
	"time"

	"github.com/AkifhanIlgaz/jwt-auth/config"
	"github.com/AkifhanIlgaz/jwt-auth/models"
	"github.com/AkifhanIlgaz/jwt-auth/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Login route
func Login(c *fiber.Ctx) error {
	// Extract the credentials from the request body
	var loginRequest models.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Find the user by credentials
	user, err := repository.FindByCredentials(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	day := 24 * time.Hour
	claims := jwt.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"fav":   user.FavoritePhrase,
		"exp":   time.Now().Add(day * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(models.LoginResponse{
		Token: t,
	})

}

func Protected(c *fiber.Ctx) error {
	// Get the user from the context and return it
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	favPhrase := claims["fav"].(string)
	return c.SendString("Welcome 👋" + email + " " + favPhrase)
}
