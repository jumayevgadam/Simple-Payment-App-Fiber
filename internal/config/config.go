package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config struct keeps,
// All needed configurations for this project
// Must be keep them in .env file.
type Config struct {
	Server struct {
		HTTPPort    string `envconfig:"HTTP_PORT" validate:"required"`
		MetricsPort string `envconfig:"METRICS_PORT" validate:"required"`
	}
	Postgres PostgresDB
	JWT      JWTOps
	Redis    RedisDB
}

// PostgresDB options for this project.
type PostgresDB struct {
	Host     string `envconfig:"DB_HOST" validate:"required"`
	Port     string `envconfig:"DB_PORT" validate:"required"`
	Name     string `envconfig:"DB_NAME" validate:"required"`
	User     string `envconfig:"DB_USER" validate:"required"`
	Password string `envconfig:"DB_PASSWORD" validate:"required"`
	SslMode  string `envconfig:"DB_SSLMODE" validate:"required"`
}

// RedisDB options for this project.
type RedisDB struct {
	Address  string `envconfig:"REDIS_ADDRESS" validate:"required"`
	Password string `envconfig:"REDIS_PASSWORD"`
}

// LoadConfig read fields of Config  struct and return it.
func LoadConfig() (*Config, error) {
	// Read .env file with this method.
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("config.LoadConfig.Load: %w", err)
	}

	// Get instance of config file.
	var c Config
	// Populate the specified struct based on environment variables.
	if err := envconfig.Process("", &c); err != nil {
		return nil, fmt.Errorf("envconfig.Process: %w", err)
	}

	// Validate the Config.
	if err := validator.New().Struct(c); err != nil {
		return nil, fmt.Errorf("error in validating Config struct: %w", err)
	}

	return &c, nil
}
