package util

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config stores all configuration of the application.
type Config struct {
	Environment          string        `mapstructure:"ENVIRONMENT"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	MigrationURL         string        `mapstructure:"MIGRATION_URL"`
	RedisAddress         string        `mapstructure:"REDIS_ADDRESS"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EmailSenderName      string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

// LoadConfig loads configuration from environment variables.
// It does not rely on an external .env file.
func LoadConfig() (Config, error) {
	var config Config

	// Attempt to load the app.env file from the current directory if it exists
	err := godotenv.Load("app.env")
	if err != nil {
		log.Println("No app.env file found. Relying solely on environment variables.")
	}

	// Assign environment variables to config
	config.Environment = os.Getenv("ENVIRONMENT")
	config.DBSource = os.Getenv("DB_SOURCE")
	config.MigrationURL = os.Getenv("MIGRATION_URL")
	config.RedisAddress = os.Getenv("REDIS_ADDRESS")
	config.HTTPServerAddress = os.Getenv("HTTP_SERVER_ADDRESS")
	config.GRPCServerAddress = os.Getenv("GRPC_SERVER_ADDRESS")
	config.TokenSymmetricKey = os.Getenv("TOKEN_SYMMETRIC_KEY")

	// Parse duration strings into time.Duration types
	accessDurationStr := os.Getenv("ACCESS_TOKEN_DURATION")
	config.AccessTokenDuration, err = time.ParseDuration(accessDurationStr)
	if err != nil {
		return config, fmt.Errorf("invalid ACCESS_TOKEN_DURATION: %w", err)
	}

	refreshDurationStr := os.Getenv("REFRESH_TOKEN_DURATION")
	config.RefreshTokenDuration, err = time.ParseDuration(refreshDurationStr)
	if err != nil {
		return config, fmt.Errorf("invalid REFRESH_TOKEN_DURATION: %w", err)
	}

	config.EmailSenderName = os.Getenv("EMAIL_SENDER_NAME")
	config.EmailSenderAddress = os.Getenv("EMAIL_SENDER_ADDRESS")
	config.EmailSenderPassword = os.Getenv("EMAIL_SENDER_PASSWORD")

	// Validate required fields
	missingVars := []string{}
	if config.Environment == "" {
		missingVars = append(missingVars, "ENVIRONMENT")
	}
	if config.DBSource == "" {
		missingVars = append(missingVars, "DB_SOURCE")
	}
	if config.AccessTokenDuration == 0 {
		missingVars = append(missingVars, "ACCESS_TOKEN_DURATION")
	}
	if config.RefreshTokenDuration == 0 {
		missingVars = append(missingVars, "REFRESH_TOKEN_DURATION")
	}
	if config.TokenSymmetricKey == "" {
		missingVars = append(missingVars, "TOKEN_SYMMETRIC_KEY")
	}
	if config.EmailSenderName == "" {
		missingVars = append(missingVars, "EMAIL_SENDER_NAME")
	}
	if config.EmailSenderAddress == "" {
		missingVars = append(missingVars, "EMAIL_SENDER_ADDRESS")
	}
	if config.EmailSenderPassword == "" {
		missingVars = append(missingVars, "EMAIL_SENDER_PASSWORD")
	}

	if len(missingVars) > 0 {
		return config, fmt.Errorf("missing required environment variables: %v", missingVars)
	}

	// Debug: Print loaded configuration (exclude sensitive data)
	if config.Environment != "production" {
		fmt.Printf("Loaded Config:\n")
		fmt.Printf("Environment: %s\n", config.Environment)
		fmt.Printf("DBSource: %s\n", config.DBSource)
		fmt.Printf("MigrationURL: %s\n", config.MigrationURL)
		fmt.Printf("RedisAddress: %s\n", config.RedisAddress)
		fmt.Printf("HTTPServerAddress: %s\n", config.HTTPServerAddress)
		fmt.Printf("GRPCServerAddress: %s\n", config.GRPCServerAddress)
		fmt.Printf("TokenSymmetricKey: [REDACTED]\n")
		fmt.Printf("AccessTokenDuration: %s\n", config.AccessTokenDuration)
		fmt.Printf("RefreshTokenDuration: %s\n", config.RefreshTokenDuration)
		fmt.Printf("EmailSenderName: %s\n", config.EmailSenderName)
		fmt.Printf("EmailSenderAddress: %s\n", config.EmailSenderAddress)
	}

	return config, nil
}
