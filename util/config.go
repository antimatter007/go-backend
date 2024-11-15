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

// LoadConfig loads configuration from the .env file and environment variables.
func LoadConfig() (Config, error) {
	var config Config

	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Relying solely on environment variables.")
	}

	// Assign environment variables to config
	config.Environment = os.Getenv("ENVIRONMENT")
	config.DBSource = os.Getenv("DB_SOURCE")
	config.MigrationURL = os.Getenv("MIGRATION_URL")
	config.RedisAddress = os.Getenv("REDIS_ADDRESS")
	config.HTTPServerAddress = os.Getenv("HTTP_SERVER_ADDRESS")
	config.GRPCServerAddress = os.Getenv("GRPC_SERVER_ADDRESS")
	config.TokenSymmetricKey = os.Getenv("TOKEN_SYMMETRIC_KEY")
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
	fmt.Printf("Loaded Config:\n")
	fmt.Printf("Environment: %s\n", config.Environment)
	fmt.Printf("DBSource: %s\n", config.DBSource)
	fmt.Printf("MigrationURL: %s\n", config.MigrationURL)
	fmt.Printf("RedisAddress: %s\n", config.RedisAddress)
	fmt.Printf("HTTPServerAddress: %s\n", config.HTTPServerAddress)
	fmt.Printf("GRPCServerAddress: %s\n", config.GRPCServerAddress)
	fmt.Printf("TokenSymmetricKey: %s\n", config.TokenSymmetricKey)
	fmt.Printf("AccessTokenDuration: %s\n", config.AccessTokenDuration)
	fmt.Printf("RefreshTokenDuration: %s\n", config.RefreshTokenDuration)
	fmt.Printf("EmailSenderName: %s\n", config.EmailSenderName)
	fmt.Printf("EmailSenderAddress: %s\n", config.EmailSenderAddress)
	// Avoid printing sensitive information
	// fmt.Printf("EmailSenderPassword: %s\n", config.EmailSenderPassword)

	return config, nil
}
