// package main

// import (
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt/v4"
// )

// type Handler struct {
// 	x bool
// }

// func (h *Handler) Login(c *fiber.Ctx) error {
// 	username := c.FormValue("username")
// 	password := c.FormValue("password")

// 	// TODO: Check if the user exists
// 	if username == "name" && password == "password" {
// 		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 			"name":  "Jon Doe",
// 			"admin": true,
// 			"exp":   time.Now().Add(72 * time.Hour).Unix(),
// 		})

// 		t, err := token.SignedString([]byte("secret"))
// 		if err != nil {
// 			return err
// 		}

// 		return c.Status(fiber.StatusOK).JSON(map[string]string{
// 			"token": t,
// 		})
// 	}

//		return c.SendStatus(fiber.StatusUnauthorized)
//	}
package main
