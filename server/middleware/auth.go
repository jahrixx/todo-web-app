package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")

		if auth == "" || !strings.HasPrefix(auth, "Bearer") {
			return c.Status(401).JSON(fiber.Map{"error": "Missing or Malformed Token"})
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid token", "details": err.Error()})
		}

		claims := token.Claims.(jwt.MapClaims)
		UserID, ok := claims["user_id"].(string)
		if !ok {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid claims"})
		}

		c.Locals("user_id", UserID)
		fmt.Println("VERIFYING WITH:", os.Getenv("JWT_SECRET"))
		return c.Next()
	}
}
