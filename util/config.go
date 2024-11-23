package util

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config stores all configuration of the application.
type Config struct {
	Environment          string        // Application environment (development, production, etc.)
	DBSource             string        // Database connection string (hardcoded)
	MigrationURL         string        // URL for database migrations
	RedisURL             string        // Redis connection URL
	RedisAddress         string        // Redis server address
	RedisPassword        string        // Redis password
	HTTPServerAddress    string        // HTTP server address
	GRPCServerAddress    string        // gRPC server address
	TokenSymmetricKey    string        // Symmetric key for token signing
	AccessTokenDuration  time.Duration // Duration for access tokens
	RefreshTokenDuration time.Duration // Duration for refresh tokens
	EmailSenderName      string        // Name of the email sender
	EmailSenderAddress   string        // Email address of the sender
	EmailSenderPassword  string        // Password for the sender's email account
}

// LoadConfig loads configuration from environment variables.
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

	// Hardcoded DBSource (PostgreSQL connection string)
	config.DBSource = "postgresql://postgres:lkjukPxvDEXTlgnxPvqHtorWdRNPjejG@autorack.proxy.rlwy.net:12999/railway?sslmode=disable"

	config.MigrationURL = os.Getenv("MIGRATION_URL")
	config.RedisURL = os.Getenv("REDIS_URL")
	config.HTTPServerAddress = os.Getenv("HTTP_SERVER_ADDRESS")
	config.GRPCServerAddress = os.Getenv("GRPC_SERVER_ADDRESS")
	config.TokenSymmetricKey = os.Getenv("TOKEN_SYMMETRIC_KEY")
	config.EmailSenderName = os.Getenv("EMAIL_SENDER_NAME")
	config.EmailSenderAddress = os.Getenv("EMAIL_SENDER_ADDRESS")
	config.EmailSenderPassword = os.Getenv("EMAIL_SENDER_PASSWORD")

	// Parse duration strings into time.Duration types
	accessTokenDurationStr := os.Getenv("ACCESS_TOKEN_DURATION")
	if accessTokenDurationStr == "" {
		return config, fmt.Errorf("ACCESS_TOKEN_DURATION is not set")
	}
	config.AccessTokenDuration, err = time.ParseDuration(accessTokenDurationStr)
	if err != nil {
		return config, fmt.Errorf("invalid ACCESS_TOKEN_DURATION: %w", err)
	}

	refreshTokenDurationStr := os.Getenv("REFRESH_TOKEN_DURATION")
	if refreshTokenDurationStr == "" {
		return config, fmt.Errorf("REFRESH_TOKEN_DURATION is not set")
	}
	config.RefreshTokenDuration, err = time.ParseDuration(refreshTokenDurationStr)
	if err != nil {
		return config, fmt.Errorf("invalid REFRESH_TOKEN_DURATION: %w", err)
	}

	// Parse Redis URL
	if config.RedisURL == "" {
		return config, fmt.Errorf("REDIS_URL is not set")
	}
	err = parseRedisURL(config.RedisURL, &config)
	if err != nil {
		return config, fmt.Errorf("failed to parse REDIS_URL: %w", err)
	}

	// Validate required fields
	missingFields := []string{}

	if config.Environment == "" {
		missingFields = append(missingFields, "ENVIRONMENT")
	}
	if config.DBSource == "" {
		missingFields = append(missingFields, "DB_SOURCE")
	}
	if config.MigrationURL == "" {
		missingFields = append(missingFields, "MIGRATION_URL")
	}
	if config.RedisAddress == "" {
		missingFields = append(missingFields, "REDIS_ADDRESS")
	}
	if config.RedisPassword == "" {
		missingFields = append(missingFields, "REDIS_PASSWORD")
	}
	if config.HTTPServerAddress == "" {
		missingFields = append(missingFields, "HTTP_SERVER_ADDRESS")
	}
	if config.GRPCServerAddress == "" {
		missingFields = append(missingFields, "GRPC_SERVER_ADDRESS")
	}
	if config.TokenSymmetricKey == "" {
		missingFields = append(missingFields, "TOKEN_SYMMETRIC_KEY")
	}
	if config.EmailSenderName == "" {
		missingFields = append(missingFields, "EMAIL_SENDER_NAME")
	}
	if config.EmailSenderAddress == "" {
		missingFields = append(missingFields, "EMAIL_SENDER_ADDRESS")
	}
	if config.EmailSenderPassword == "" {
		missingFields = append(missingFields, "EMAIL_SENDER_PASSWORD")
	}

	if len(missingFields) > 0 {
		return config, fmt.Errorf("missing required environment variables: %v", missingFields)
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
		// Do not print EmailSenderPassword or RedisPassword
	}

	return config, nil
}

// parseRedisURL parses the Redis URL and updates the config with address and password.
func parseRedisURL(redisURL string, config *Config) error {
	parsedURL, err := url.Parse(redisURL)
	if err != nil {
		return fmt.Errorf("invalid Redis URL: %w", err)
	}

	// Extract host and port
	config.RedisAddress = parsedURL.Host

	// Extract password
	if parsedURL.User != nil {
		config.RedisPassword, _ = parsedURL.User.Password()
	}

	return nil
}
