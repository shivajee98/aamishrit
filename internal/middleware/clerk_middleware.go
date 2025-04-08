package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/gofiber/fiber/v2"
)

type contextKey string

const userIDKey contextKey = "userID"

func ClerkMiddleware(secretKey string) fiber.Handler {
	// Set Clerk key only once — safe to call repeatedly, will just override internally
	clerk.SetKey(secretKey)
	log.Println("Clerk key initialized from middleware")

	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Missing Authorization header",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Malformed Authorization header",
			})
		}

		token := parts[1]
		if token == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Empty token",
			})
		}

		claims, err := jwt.Verify(c.Context(), &jwt.VerifyParams{
			Token: token,
		})
		if err != nil {
			log.Printf("JWT verification failed: %v", err)
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Invalid token",
				"error":   err.Error(),
			})
		}

		log.Printf("JWT verified. User ID: %s", claims.Subject)
		c.Locals(string(userIDKey), claims.Subject)

		return c.Next()
	}
}
