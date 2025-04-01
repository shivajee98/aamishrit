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

// JWKStore for caching JSON Web Keys
type JWKStore struct {
	mu  sync.RWMutex
	jwk *clerk.JSONWebKey
}

func NewJWKStore() *JWKStore {
	return &JWKStore{}
}

func (s *JWKStore) GetJWK() *clerk.JSONWebKey {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.jwk
}

func (s *JWKStore) SetJWK(jwk *clerk.JSONWebKey) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jwk = jwk
}

// ClerkAuthMiddleware authenticates requests using Clerk
func ClerkAuthMiddleware(jwksClient *jwks.Client, store *JWKStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
		}

		sessionToken := strings.TrimPrefix(authHeader, "Bearer ")
		if sessionToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
		}

		jwk := store.GetJWK()
		if jwk == nil {
			unsafeClaims, err := jwt.Decode(context.Background(), &jwt.DecodeParams{Token: sessionToken})
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to decode token"})
			}

			jwk, err = jwt.GetJSONWebKey(context.Background(), &jwt.GetJSONWebKeyParams{
				KeyID:      unsafeClaims.KeyID,
				JWKSClient: jwksClient,
			})
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to fetch JWK"})
			}

			store.SetJWK(jwk)
		}

		claims, err := jwt.Verify(context.Background(), &jwt.VerifyParams{Token: sessionToken, JWK: jwk})
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		usr, err := user.Get(context.Background(), claims.Subject)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
		}

		c.Locals("user_id", usr.ID)
		return c.Next()
	}
}
