package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from environment variables.
type Config struct {
	Environment          string        `mapstructure:"ENVIRONMENT" validate:"required"`
	DBSource             string        `mapstructure:"DB_SOURCE" validate:"required,uri"`
	MigrationURL         string        `mapstructure:"MIGRATION_URL" validate:"required,uri"`
	RedisAddress         string        `mapstructure:"REDIS_ADDRESS" validate:"required,uri"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS" validate:"required,hostname_port"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS" validate:"required,hostname_port"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY" validate:"required,min=32,max=32"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION" validate:"required"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION" validate:"required"`
	EmailSenderName      string        `mapstructure:"EMAIL_SENDER_NAME" validate:"required"`
	EmailSenderAddress   string        `mapstructure:"EMAIL_SENDER_ADDRESS" validate:"required,email"`
	EmailSenderPassword  string        `mapstructure:"EMAIL_SENDER_PASSWORD" validate:"required"`
}

// LoadConfig loads configuration from environment variables.

func LoadConfig() (config Config, err error) {
	viper.AutomaticEnv() // Automatically read environment variables

	// Replace dots with underscores in environment variables
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("unable to decode into struct: %w", err)
	}

	// Validate the configuration
	validate := validator.New()
	err = validate.Struct(config)
	if err != nil {
		return config, fmt.Errorf("configuration validation failed: %w", err)
	}

	// Debug: print each configuration field individually
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
