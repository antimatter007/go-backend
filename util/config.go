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

// LoadConfig loads configuration from the specified path's app.env file and environment variables.
// If the app.env file does not exist at the specified path, it falls back to environment variables.
func LoadConfig(path string) (Config, error) {
	var config Config

	// Attempt to load the app.env file from the specified path
	envFilePath := fmt.Sprintf("%s/app.env", path)
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Printf("No app.env file found at %s. Relying solely on environment variables.\n", envFilePath)
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
	config.AccessTokenDuration, err = time.ParseDuration(os.Getenv("ACCESS_TOKEN_DURATION"))
	if err != nil {
		return config, fmt.Errorf("invalid ACCESS_TOKEN_DURATION: %w", err)
	}

	config.RefreshTokenDuration, err = time.ParseDuration(os.Getenv("REFRESH_TOKEN_DURATION"))
	if err != nil {
		return config, fmt.Errorf("invalid REFRESH_TOKEN_DURATION: %w", err)
	}

	config.EmailSenderName = os.Getenv("EMAIL_SENDER_NAME")
	config.EmailSenderAddress = os.Getenv("EMAIL_SENDER_ADDRESS")
	config.EmailSenderPassword = os.Getenv("EMAIL_SENDER_PASSWORD")

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
