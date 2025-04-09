package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shivajee98/aamishrit/internal/config"
)

// Struct to hold Clerk user data response
type ClerkUser struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phone_number"` // This is assuming "phone_number" is the correct field name in Clerk API.
}

func FetchClerkUser(userID string) (*ClerkUser, error) {
	url := fmt.Sprintf("https://api.clerk.dev/v1/users/%s", userID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	cfg := config.LoadEnv()

	ClerkSecretKey := cfg.ClerkSecretKey

	req.Header.Set("Authorization", "Bearer "+ClerkSecretKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Clerk returned status %d", resp.StatusCode)
	}

	var userData ClerkUser
	if err := json.NewDecoder(resp.Body).Decode(&userData); err != nil {
		return nil, err
	}
	fmt.Println("User Data from Clerk:", userData)

	return &userData, nil
}
