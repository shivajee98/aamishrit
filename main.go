// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strings"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/cors"
// )

// func main() {
// 	app := fiber.New()

// 	// ✅ CORS Middleware
// 	app.Use(cors.New(cors.Config{
// 		AllowOrigins: "http://localhost:5173",
// 		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
// 		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
// 	}))

// 	clerkSecret := "pk_test_Y3J1Y2lhbC1qYWd1YXItNDUuY2xlcmsuYWNjb3VudHMuZGV2JA"

// 	// ✅ Protected route middleware
// 	app.Use("/api", ClerkMiddleware(clerkSecret))

// 	app.Get("/api/protected", func(c *fiber.Ctx) error {
// 		userID := c.Locals("user_id").(string)
// 		return c.JSON(fiber.Map{
// 			"message": "You are authenticated",
// 			"user_id": userID,
// 		})
// 	})

// 	log.Fatal(app.Listen(":3000"))
// }

// func ClerkMiddleware(secretKey string) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		authHeader := c.Get("Authorization")
// 		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid auth header"})
// 		}

// 		token := strings.TrimPrefix(authHeader, "Bearer ")

// 		req, _ := http.NewRequest("GET", "https://api.clerk.dev/v1/sessions/"+token, nil)
// 		req.Header.Set("Authorization", "Bearer "+secretKey)

// 		client := &http.Client{}
// 		resp, err := client.Do(req)
// 		if err != nil || resp.StatusCode != http.StatusOK {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid session"})
// 		}
// 		defer resp.Body.Close()

// 		var session struct {
// 			UserID string `json:"user_id"`
// 		}
// 		if err := json.NewDecoder(resp.Body).Decode(&session); err != nil {
// 			fmt.Println("Failed to decode session:", err)
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid session response"})
// 		}

// 		c.Locals("user_id", session.UserID)
// 		return c.Next()
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"github.com/clerk/clerk-sdk-go/v2"
// 	"github.com/clerk/clerk-sdk-go/v2/jwt"
// 	"github.com/clerk/clerk-sdk-go/v2/user"
// )

// // Middleware to handle CORS
// func corsMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Allow CORS for requests coming from port 5173
// 		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

// 		// Handle OPTIONS request (preflight)
// 		if r.Method == http.MethodOptions {
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}

// 		// Call the next handler
// 		next.ServeHTTP(w, r)
// 	})
// }

// func main() {
// 	clerk.SetKey("sk_test_E4byMmlHEmRqxrksTBYabveYG1yZiZNlzj7CEV46mh")

// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", publicRoute)
// 	mux.HandleFunc("/protected", protectedRoute)

// 	// Wrap the mux with CORS middleware
// 	http.ListenAndServe(":3000", corsMiddleware(mux))
// }

// func publicRoute(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(`{"access": "public"}`))
// }

// func protectedRoute(w http.ResponseWriter, r *http.Request) {
// 	// Get the session JWT from the Authorization header
// 	sessionToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

// 	// Verify the session
// 	claims, err := jwt.Verify(r.Context(), &jwt.VerifyParams{
// 		Token: sessionToken,
// 	})
// 	if err != nil {
// 		// handle the error
// 		w.WriteHeader(http.StatusUnauthorized)
// 		w.Write([]byte(`{"access": "unauthorized"}`))
// 		return
// 	}

// 	usr, err := user.Get(r.Context(), claims.Subject)
// 	if err != nil {
// 		// handle the error
// 	}
// 	fmt.Fprintf(w, `{"user_id": "%s", "user_banned": "%t"}`, usr.ID, usr.Banned)
// }

package main

import (
	"log"
	"net/http" // For status codes
	"os"
	"strings"

	// Clerk v2 SDK - Ensure you have a recent version (go get github.com/clerk/clerk-sdk-go/v2)
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/jwt" // Use the v2 jwt package

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger" // Added logger from example
)

// Custom type for context key to avoid collisions
type contextKey string

const userIDKey contextKey = "userID"

func main() {
	// --- Configuration ---
	clerkSecretKey := os.Getenv("CLERK_SECRET_KEY")
	if clerkSecretKey == "" {
		log.Println("WARN: CLERK_SECRET_KEY environment variable not set. Using placeholder (NOT recommended).")
		// Replace with your actual key during development if needed, but env var is best.
		clerkSecretKey = "sk_test_E4byMmlHEmRqxrksTBYabveYG1yZiZNlzj7CEV46mh"
	}
	if !strings.HasPrefix(clerkSecretKey, "sk_") {
		log.Fatal("FATAL: Clerk Secret Key must start with 'sk_'")
	}

	listenAddr := os.Getenv("HOST")
	if listenAddr == "" {
		listenAddr = ":3000" // Default port
	}

	// --- Clerk Initialization (Call ONCE at startup) ---
	// Setting the key globally enables jwt.Verify to automatically handle JWKS fetching.
	// This requires a recent Clerk v2 SDK version (like v2.0.2-beta.7 or later stable releases).
	clerk.SetKey(clerkSecretKey)
	log.Println("Clerk Secret Key set.")

	// --- Fiber App Initialization ---
	app := fiber.New()

	// --- Global Middleware ---
	app.Use(
		// Basic logging
		logger.New(),
		// CORS configuration
		cors.New(cors.Config{
			AllowOrigins: "http://localhost:5173", // Adjust to your frontend URL
			AllowHeaders: "Origin, Content-Type, Accept, Authorization",
			AllowMethods: "GET, POST, OPTIONS, PUT, DELETE",
		}),
	)

	// --- Public Route ---
	app.Get("/", publicRoute)

	// --- Protected Routes Group ---
	// Apply the authorization middleware only to routes within this group
	protected := app.Group("/protected", ClerkAuthMiddleware()) // Apply middleware here

	protected.Get("/", protectedRoute)
	// Add more protected routes here:
	// protected.Post("/items", createItemHandler)
	// protected.Get("/items/:id", getItemHandler)

	// --- Start Server ---
	log.Printf("Starting server on %s\n", listenAddr)
	log.Fatal(app.Listen(listenAddr))
}

// ClerkAuthMiddleware creates a Fiber middleware handler for checking Clerk JWT tokens.
// It relies on clerk.SetKey having been called during initialization.
func ClerkAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("Running Clerk Auth Middleware...") // Log entry

		authHeader := c.Get("Authorization") // Use c.Get for simplicity
		if authHeader == "" {
			log.Println("Middleware: Missing Authorization header")
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Missing Authorization header",
			})
		}

		// Expecting "Bearer YOUR_TOKEN"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			log.Println("Middleware: Malformed Authorization header")
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Malformed Authorization header",
			})
		}
		sessionToken := parts[1]

		if sessionToken == "" {
			log.Println("Middleware: Missing token after Bearer")
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Missing token",
			})
		}

		// Verify the token using Clerk v2 jwt package
		// clerk.SetKey() ensures JWKS are fetched automatically
		claims, err := jwt.Verify(c.Context(), &jwt.VerifyParams{
			Token: sessionToken,
			// You can add Leeway here if needed:
			// Leeway: 60 * time.Second,
		})

		if err != nil {
			// Log the specific verification error
			log.Printf("Middleware: Clerk token verification failed: %v\n", err)
			// Check for specific errors like clock skew if needed
			if strings.Contains(err.Error(), "token issued in the future") {
				log.Println("Middleware: Possible clock skew detected.")
				// Potentially return a more specific error message or handle differently
			}
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized: Invalid token",
				"error":   err.Error(), // Optionally include error in dev mode
			})
		}

		// --- Token is valid ---
		log.Printf("Middleware: Token verified successfully for UserID: %s\n", claims.Subject)

		// Store the UserID in context locals for downstream handlers
		c.Locals(string(userIDKey), claims.Subject) // Use contextKey type

		// Optional: You could fetch user details here if needed globally,
		// but often it's better to do it in the specific handler that needs it.
		// client, _ := clerk.NewClient(clerkSecretKey) // Create client if needed
		// usr, err := client.Users().Read(c.Context(), claims.Subject) ...

		// Proceed to the next middleware or route handler
		return c.Next()
	}
}

// publicRoute handler
func publicRoute(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"access": "public",
	})
}

// protectedRoute handler - runs *after* ClerkAuthMiddleware
func protectedRoute(c *fiber.Ctx) error {
	// Retrieve the UserID stored by the middleware
	userIDValue := c.Locals(string(userIDKey)) // Use contextKey type

	// Type assert the value retrieved from locals
	userID, ok := userIDValue.(string)
	if !ok || userID == "" {
		// This should theoretically not happen if middleware ran correctly,
		// but it's good practice to check.
		log.Println("Protected Route: UserID not found in context or is not a string.")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error: User context not found",
		})
	}

	log.Printf("Protected Route: Access granted for UserID: %s\n", userID)

	// Now you can use the userID
	// Example: Fetch user-specific data from DB, etc.

	// You could fetch full user details here if needed for this specific route
	// client, _ := clerk.NewClient(os.Getenv("CLERK_SECRET_KEY")) // Maybe create client once and pass around?
	// usr, err := client.Users().Read(c.Context(), userID)
	// if err != nil { ... handle error ... }

	return c.JSON(fiber.Map{
		"access":  "protected",
		"user_id": userID,
		// "user_details": usr, // If fetched
	})
}
