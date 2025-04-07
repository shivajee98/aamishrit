package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shivajee98/aamishrit/internal/config"
)

func FetchClerkUser(userID string) (map[string]any, error) {
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

	var userData map[string] any
	if err := json.NewDecoder(resp.Body).Decode(&userData); err != nil {
		return nil, err
	}

	return userData, nil
}
