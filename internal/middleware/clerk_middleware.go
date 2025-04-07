package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ClerkMiddleware(secretKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid auth header"})
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate session and extract user_id
		req, _ := http.NewRequest("GET", "https://api.clerk.dev/v1/sessions/"+token, nil)
		req.Header.Set("Authorization", "Bearer "+secretKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != http.StatusOK {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid session"})
		}
		defer resp.Body.Close()

		var session struct {
			UserID string `json:"user_id"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&session); err != nil {
			fmt.Println("Failed to decode session:", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid session response"})
		}

		// Inject into context
		c.Locals("user_id", session.UserID)

		return c.Next()
	}
}
