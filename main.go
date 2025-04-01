package main

import (
	"fmt"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/rs/cors"
)

func main() {
	clerk.SetKey("sk_test_E4byMmlHEmRqxrksTBYabveYG1yZiZNlzj7CEV46mh")

	mux := http.NewServeMux()
	mux.HandleFunc("/", publicRoute)

	protectedHandler := http.HandlerFunc(protectedRoute)
	mux.Handle(
		"/protected",
		clerkhttp.WithHeaderAuthorization()(protectedHandler),
	)

	// CORS Configuration
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Allow frontend origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true, // Allow cookies/auth headers
	})

	// Wrap mux with CORS middleware
	handler := corsHandler.Handler(mux)

	fmt.Println("Server running on port 3000...")
	http.ListenAndServe(":3000", handler)
}

func publicRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"access": "public"}`))
}

func protectedRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Protected route hit")

	authHeader := r.Header.Get("Authorization")
	fmt.Println("Authorization Header:", authHeader) // Print received token

	_, ok := clerk.SessionClaimsFromContext(r.Context())
	if !ok {
		fmt.Println("No claims found, unauthorized request")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"access": "unauthorized"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"access": "protected"}`))
}
