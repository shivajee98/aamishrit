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

		// ctx := c.Context()
		// userDetails, err := user.Get(ctx, claims.Subject)
		// jsonData, _ := json.Marshal(userDetails.ID)
		c.Locals(string(UserIDKey), claims.Subject)
		// log.Printf("%s : %s", UserIDKey, string(jsonData))

		return c.Next()
	}
}

func ClerkAdminMiddleware(secretKey string) fiber.Handler {
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

		// Get full user details
		ctx := c.Context()
		userDetails, err := user.Get(ctx, claims.Subject)
		if err != nil {
			log.Printf("Error fetching user from Clerk: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error: Failed to get user from Clerk",
				"error":   err.Error(),
			})
		}

		// Unmarshal public metadata into map
		var metadata map[string]interface{}
		if err := json.Unmarshal(userDetails.PublicMetadata, &metadata); err != nil {
			log.Printf("Failed to unmarshal public metadata for user %s: %v", claims.Subject, err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error: Failed to parse public metadata",
			})
		}

		// Extract role
		roleVal, ok := metadata["role"]
		if !ok {
			log.Printf("Missing role in public metadata for user %s", claims.Subject)
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden: Admin access required (no role found)",
			})
		}

		roleStr, valid := roleVal.(string)
		if !valid || roleStr != "admin" {
			log.Printf("Access denied: user %s is not admin", claims.Subject)
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden: Admin access required",
			})
		}

		// Log full user object if needed
		userJson, _ := json.Marshal(userDetails)
		log.Printf("Authorized admin user: %s", string(userJson))

		c.Locals(string(UserIDKey), claims.Subject)
		return c.Next()
	}
}
