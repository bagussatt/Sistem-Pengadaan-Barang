package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"purchasing-backend/config"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"error": "Missing Authorization header"})
	}

	// Format harus: Bearer <token>
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"error": "Invalid Authorization format"})
	}

	tokenString := parts[1]

	// Parse & validate token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Pastikan signing method HMAC
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return config.JWTSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	// Ambil claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"error": "Invalid token claims"})
	}

	// Ambil user_id dari JWT
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"error": "user_id not found in token"})
	}

	// Simpan ke context
	c.Locals("user_id", uint(userIDFloat))

	return c.Next()
}
