package config

import (
	"github.com/go-ozzo/ozzo-validation/v4/is"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type serverConfig struct {
	Port        int    `env:"PORT,default=8080"`
	Host        string `env:"HOST,default=0.0.0.0"`
	TLSCertPath string `env:"TLS_CERT_PATH,omitempty"`
	TLSKeyPath  string `env:"TLS_KEY_PATH,omitempty"`
}

// Validate checks if s is a valid serverConfig instance. It returns an error if any required fields are missing or invalid.
func (s *serverConfig) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Port, validation.Required, validation.Min(1), validation.Max(65535)),
		validation.Field(&s.Host, validation.Required, is.Host),
		validation.Field(&s.TLSCertPath, validation.When(s.TLSKeyPath != "", validation.Required)),
		validation.Field(&s.TLSKeyPath, validation.When(s.TLSCertPath != "", validation.Required)),
	)
}
