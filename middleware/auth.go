package middleware

import (
	"fmt"
	"todo_api/models"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(c *fiber.Ctx) error {

	if strings.HasPrefix(c.Path(), "/auth") {
		return c.Next()
	}

	header := c.Get("Authorization")

	if header == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	if !strings.HasPrefix(header, "Bearer") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	
	secretKey := os.Getenv("JWT_SECRET")
	
	acces_token := strings.Split(header, " ")[1]
	
	claims := &models.Claim{}
	
	_, err := jwt.ParseWithClaims(acces_token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secretKey), nil
		}
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.Locals("user", claims)

	return c.Next()
}
