package util

import (
	"fmt"
	"strings"
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
	viper.AutomaticEnv()

	// Set a key replacer to match environment variable names
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("unable to decode into struct: %w", err)
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
