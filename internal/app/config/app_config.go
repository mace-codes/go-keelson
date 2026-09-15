package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Config represents the application configuration.
type Config struct {
	ServiceName string `env:"SERVICE_NAME"`
	Env         AppEnv `env:"ENV"`
	LogLevel    string `env:"LOG_LEVEL,default=info"`
	Server      serverConfig
}

// Validate checks if c is a valid Config instance. It returns an error if any required fields are missing or invalid.
func (c *Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ServiceName, validation.Required),
		validation.Field(&c.Env, validation.Required, validation.In(EnvLocal, EnvDevelopment, EnvStaging, EnvTest, EnvProduction)),
		validation.Field(&c.LogLevel, validation.Required),
		validation.Field(&c.Server, validation.Required),
	)
}

// Load loads the application configuration from environment variables and validates it. It returns a Config instance or an error if the configuration is invalid.
func Load() (*Config, error) {
	var cfg Config
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
