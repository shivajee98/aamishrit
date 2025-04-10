package middleware

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gofiber/fiber/v2"
)

type contextKey string

const UserIDKey contextKey = "clerk_id"

func ClerkMiddleware(secretKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Missing Authorization header",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Malformed Authorization header",
			})
		}

		token := parts[1]
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Empty token",
			})
		}

		claims, err := jwt.Verify(c.Context(), &jwt.VerifyParams{
			Token: token,
		})
		if err != nil {
			log.Printf("JWT verification failed: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Invalid token",
				"error":   err.Error(),
			})
		}

		ctx := c.Context()
		userDetails, err := user.Get(ctx, claims.Subject)
		jsonData, _ := json.Marshal(userDetails)
		c.Locals(string(UserIDKey), claims.Subject)
		log.Printf("%s : %s", UserIDKey, string(jsonData))

		return c.Next()
	}
}

func AdminMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// I will resove the error in next commit
		user := c.Locals("user").(YourUserModel)

		if user.Role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Admin access only",
			})
		}

		return c.Next()
	}
}
