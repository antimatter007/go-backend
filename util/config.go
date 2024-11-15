package util

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from environment variables.
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
func LoadConfig() (config Config, err error) {
	viper.AutomaticEnv() // Automatically read environment variables

	// Optionally, set a prefix if your env variables have one
	// viper.SetEnvPrefix("APP")

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("unable to decode into struct: %w", err)
	}

	// Debug: print the loaded configuration (excluding sensitive data)
	fmt.Printf("Loaded Config: %+v\n", config)

	return config, nil
}
