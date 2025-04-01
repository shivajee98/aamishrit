package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL    string
	ClerkSecretKey string
}

func LoadEnv() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	return &Config{
		DatabaseURL: os.Getenv("SUPABASE_URL"),
		ClerkSecretKey: os.Getenv("CLERK_SECRET_KEY"),
	}
}
