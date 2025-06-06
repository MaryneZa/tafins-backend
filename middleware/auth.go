package middleware

import (
	"os"
	"fmt"
	"strings"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type TokenCustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}


func AuthMiddleware(c fiber.Ctx) error {
	secretKey := os.Getenv("JWT_SECRETKEY")

	authHeader := c.Get("Authorization")
	var tokenString string

	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString = authHeader[7:]
	} else {
		tokenString = c.Cookies("access_token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
	}

	token, err := jwt.ParseWithClaims(tokenString, &TokenCustomClaims{}, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error, unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	} 
	if claims, ok := token.Claims.(*TokenCustomClaims); ok && token.Valid {
		c.Locals("user_id", claims.UserID)
        return c.Next()
    }
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token claims"})
}