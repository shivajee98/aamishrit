package middleware

import (
	"context"
	"strings"
	"sync"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwks"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gofiber/fiber/v2"
)

// JWKStore is an in-memory cache for JSON Web Keys
type JWKStore struct {
	mu  sync.RWMutex
	jwk *clerk.JSONWebKey
}

// GetJWK retrieves the cached JWK
func (s *JWKStore) GetJWK() *clerk.JSONWebKey {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.jwk
}

// SetJWK caches the JWK for later use
func (s *JWKStore) SetJWK(jwk *clerk.JSONWebKey) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jwk = jwk
}

func NewJWKStore() *JWKStore {
	return &JWKStore{}
}

// ClerkAuthMiddleware is the authentication middleware for Fiber
func ClerkAuthMiddleware(jwksClient *jwks.Client, store *JWKStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
		}

		// Extract the Bearer token
		sessionToken := strings.TrimPrefix(authHeader, "Bearer ")
		if sessionToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
		}

		// Try to retrieve cached JWK
		jwk := store.GetJWK()
		if jwk == nil {
			// Decode token to get Key ID
			unsafeClaims, err := jwt.Decode(context.Background(), &jwt.DecodeParams{
				Token: sessionToken,
			})
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to decode token"})
			}

			// Fetch the JSON Web Key
			jwk, err = jwt.GetJSONWebKey(context.Background(), &jwt.GetJSONWebKeyParams{
				KeyID:      unsafeClaims.KeyID,
				JWKSClient: jwksClient,
			})
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to fetch JWK"})
			}

			// Store the JWK for future requests
			store.SetJWK(jwk)
		}

		// Verify the session
		claims, err := jwt.Verify(context.Background(), &jwt.VerifyParams{
			Token: sessionToken,
			JWK:   jwk,
		})
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		// Fetch user details
		usr, err := user.Get(context.Background(), claims.Subject)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
		}

		// Attach user ID to the context for later use
		c.Locals("user_id", usr.ID)

		return c.Next()
	}
}
