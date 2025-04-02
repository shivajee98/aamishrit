// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/clerk/clerk-sdk-go/v2"
// 	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
// 	"github.com/clerk/clerk-sdk-go/v2/user"
// 	"github.com/rs/cors"
// )

// func main() {
// 	clerk.SetKey("sk_test_E4byMmlHEmRqxrksTBYabveYG1yZiZNlzj7CEV46mh")

// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", publicRoute)

// 	protectedHandler := http.HandlerFunc(protectedRoute)
// 	mux.Handle(
// 		"/protected",
// 		clerkhttp.WithHeaderAuthorization()(protectedHandler),
// 	)

// 	// CORS Configuration
// 	corsHandler := cors.New(cors.Options{
// 		AllowedOrigins:   []string{"http://localhost:5173"},
// 		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 		AllowedHeaders:   []string{"Authorization", "Content-Type"},
// 		AllowCredentials: true,
// 	})

// 	// Wrap mux with CORS middleware
// 	handler := corsHandler.Handler(mux)

// 	fmt.Println("Server running on port 3000...")
// 	http.ListenAndServe(":3000", handler)
// }

// func publicRoute(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(`{"access": "public"}`))
// }

// func protectedRoute(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Protected route hit")

// 	claims, ok := clerk.SessionClaimsFromContext(r.Context())
// 	if !ok {
// 		fmt.Println("No claims found, unauthorized request")
// 		w.WriteHeader(http.StatusUnauthorized)
// 		w.Write([]byte(`{"access": "unauthorized"}`))
// 		return
// 	}

// 	userID := claims.Subject // Extract user ID from Clerk session
// 	fmt.Println("User ID:", userID)

// 	// Fetch user details
// 	ctx := context.Background()
// 	userDetails, err := user.Get(ctx, userID)
// 	if err != nil {
// 		fmt.Println("Error fetching user details:", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(`{"error": "failed to retrieve user"}`))
// 		return
// 	}

// 	// Convert user details to JSON
// 	userJSON, _ := json.Marshal(userDetails.PhoneNumbers)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(userJSON)
// }

// package main

// import (
// 	"context"
// 	"log"
// 	"strings"
// 	"time"

// 	"github.com/clerk/clerk-sdk-go/v2"
// 	"github.com/clerk/clerk-sdk-go/v2/jwks"
// 	"github.com/clerk/clerk-sdk-go/v2/jwt"
// 	"github.com/clerk/clerk-sdk-go/v2/user"
// 	"github.com/gofiber/fiber/v2"
// )

// // JWKStore to cache JSON Web Key
// type JWKStore struct {
// 	key         *clerk.JSONWebKey
// 	lastFetched time.Time
// }

// func (s *JWKStore) GetJWK() *clerk.JSONWebKey {
// 	// Cache expires after 5 minutes (adjust as needed)
// 	if s.key != nil && time.Since(s.lastFetched) < 5*time.Minute {
// 		return s.key
// 	}
// 	return nil
// }

// func (s *JWKStore) SetJWK(jwk *clerk.JSONWebKey) {
// 	s.key = jwk
// 	s.lastFetched = time.Now()
// }

// func NewJWKStore() *JWKStore {
// 	return &JWKStore{}
// }

// func main() {
// 	app := fiber.New()
// 	jwkStore := NewJWKStore()
// 	app.Use(func(c *fiber.Ctx) error {
// 		c.Set("Access-Control-Allow-Origin", "http://localhost:5173")
// 		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
// 		c.Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
// 		c.Set("Access-Control-Allow-Credentials", "true")

// 		if c.Method() == "OPTIONS" {
// 			return c.SendStatus(fiber.StatusNoContent)
// 		}

// 		return c.Next()
// 	})

// 	config := &clerk.ClientConfig{}
// 	config.Key = clerk.String("sk_test_E4byMmlHEmRqxrksTBYabveYG1yZiZNlzj7CEV46mh") // Replace with actual secret key
// 	jwksClient := jwks.NewClient(config)

// 	// Public route (no authentication required)
// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.JSON(fiber.Map{"access": "public"})
// 	})

// 	// Protected route with JWT verification
// 	app.Get("/protected", protectedRoute(jwksClient, jwkStore))

// 	log.Println("Server running on port 3000...")
// 	log.Fatal(app.Listen(":3000"))
// }

// func protectedRoute(jwksClient *jwks.Client, store *JWKStore) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		// Get session token from the Authorization header
// 		sessionToken := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

// 		if sessionToken == "" {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"access": "unauthorized"})
// 		}

// 		// Try fetching cached JWK
// 		jwk := store.GetJWK()
// 		if jwk == nil {
// 			// Decode token to extract Key ID
// 			unsafeClaims, err := jwt.Decode(context.Background(), &jwt.DecodeParams{
// 				Token: sessionToken,
// 			})
// 			if err != nil {
// 				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
// 			}

// 			// Fetch new JWK from Clerk
// 			jwk, err = jwt.GetJSONWebKey(context.Background(), &jwt.GetJSONWebKeyParams{
// 				KeyID:      unsafeClaims.KeyID,
// 				JWKSClient: jwksClient,
// 			})
// 			if err != nil {
// 				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "failed to fetch JWK"})
// 			}

// 			// Cache the JWK for future use
// 			store.SetJWK(jwk)
// 		}

// 		// Verify token using cached JWK
// 		claims, err := jwt.Verify(context.Background(), &jwt.VerifyParams{
// 			Token: sessionToken,
// 			JWK:   jwk,
// 		})
// 		if err != nil {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
// 		}

// 		// Fetch user details from Clerk
// 		usr, err := user.Get(context.Background(), claims.Subject)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to retrieve user"})
// 		}

// 		return c.JSON(fiber.Map{
// 			"user_id":       usr.ID,
// 			"user_banned":   usr.Banned,
// 			"phone_numbers": usr.PhoneNumbers,
// 		})
// 	}
// }

package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwks"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/clerk/clerk-sdk-go/v2/user"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", publicRoute)

	// Initialize storage for JSON Web Keys. You can cache/store
	// the key for as long as it's valid, and pass it to jwt.Verify.
	// This way jwt.Verify won't make requests to the Clerk
	// Backend API to refetch the JSON Web Key.
	// Make sure you refetch the JSON Web Key whenever your
	// Clerk Secret Key changes.
	jwkStore := NewJWKStore()

	config := &clerk.ClientConfig{}
	config.Key = clerk.String("sk_test_E4byMmlHEmRqxrksTBYabveYG1yZiZNlzj7CEV46mh")
	jwksClient := jwks.NewClient(config)
	mux.HandleFunc("/protected", protectedRoute(jwksClient, jwkStore))

	http.ListenAndServe(":3000", mux)
}

func publicRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"access": "public"}`))
}

func protectedRoute(jwksClient *jwks.Client, store JWKStore) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the session JWT from the Authorization header
		sessionToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

		// Attempt to get the JSON Web Key from your store.
		jwk := store.GetJWK()
		if jwk == nil {
			// Decode the session JWT so that we can find the key ID.
			unsafeClaims, err := jwt.Decode(r.Context(), &jwt.DecodeParams{
				Token: sessionToken,
			})
			if err != nil {
				// handle the error
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"access": "unauthorized"}`))
				return
			}

			// Fetch the JSON Web Key
			jwk, err = jwt.GetJSONWebKey(r.Context(), &jwt.GetJSONWebKeyParams{
				KeyID:      unsafeClaims.KeyID,
				JWKSClient: jwksClient,
			})
			if err != nil {
				// handle the error
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"access": "unauthorized"}`))
				return
			}
		}
		// Write the JSON Web Key to your store, so that next time
		// you can use the cached value.
		store.SetJWK(jwk)

		// Verify the session
		claims, err := jwt.Verify(r.Context(), &jwt.VerifyParams{
			Token: sessionToken,
			JWK:   jwk,
		})
		if err != nil {
			// handle the error
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"access": "unauthorized"}`))
			return
		}

		usr, err := user.Get(r.Context(), claims.Subject)
		if err != nil {
			// handle the error
		}
		fmt.Fprintf(w, `{"user_id": "%s", "user_banned": "%t"}`, usr.ID, usr.Banned)
	}
}

// Sample interface for JSON Web Key storage.
// Implementation may vary.
type JWKStore interface {
	GetJWK() *clerk.JSONWebKey
	SetJWK(*clerk.JSONWebKey)
}

func NewJWKStore() JWKStore {
	// Implementation may vary. This can be an
	// in-memory store, database, caching layer,...
	return nil
}
