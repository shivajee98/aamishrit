package middleware

import (
	"net/http"
	"strings"
	"sync"

	"github.com/clerk/clerk-sdk-go/v2/jwks"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gofiber/fiber/v2"
)

// JWKCache is a simple in-memory store for JSON Web Keys
type JWKCache struct {
	mu  sync.RWMutex
	key *jwks.JSONWebKey
}

var jwkCache = &JWKCache{}

func (c *JWKCache) Get() *jwks.JSONWebKey {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.key
}

func (c *JWKCache) Set(key *jwks.JSONWebKey) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.key = key
}

func ClerkAuthMiddleware(c *fiber.Ctx) error {
	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
	if tokenString == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
	}

	// Try to use cached JWK
	jwk := jwkCache.Get()
	if jwk == nil {
		// Decode the token to extract Key ID
		unsafeClaims, err := jwt.Decode(c.Context(), &jwt.DecodeParams{
			Token: tokenString,
		})
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		// Fetch the JSON Web Key from Clerk
		jwk, err = jwt.GetJSONWebKey(c.Context(), &jwt.GetJSONWebKeyParams{
			KeyID: unsafeClaims.KeyID,
		})
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Failed to fetch JWK"})
		}

		// Cache the JWK
		jwkCache.Set(jwk)
	}

	// Verify JWT using the cached JWK
	claims, err := jwt.Verify(c.Context(), &jwt.VerifyParams{
		Token: tokenString,
		JWK:   jwk,
	})
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorised"})
	}

	// Fetch user details
	_, err = user.Get(c.Context(), claims.Subject)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User Not Found"})
	}

	return c.Next()
}
