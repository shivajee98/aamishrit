package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string
}

type AppConfig struct {
	ClerkSecretKey string
}

func LoadEnv() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	return &Config{
		DatabaseURL: os.Getenv("SUPABASE_URL"),
	}
}

func LoadConfig() *AppConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error Reading Config file: %v", err)
	}

	return &AppConfig{
		ClerkSecretKey: viper.GetString("clerk.secret_key"),
	}
}
