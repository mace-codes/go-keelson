package config

// AppEnv represents the application environment configuration.
type AppEnv string

// Valid AppEnv values.
const (
	EnvLocal       AppEnv = "local"
	EnvDevelopment AppEnv = "development"
	EnvStaging     AppEnv = "staging"
	EnvTest        AppEnv = "test"
	EnvProduction  AppEnv = "production"
)

// IsValid checks if the AppEnv value is one of the valid predefined environments.
func (e AppEnv) IsValid() bool {
	switch e {
	case EnvLocal, EnvDevelopment, EnvStaging, EnvTest, EnvProduction:
		return true
	default:
		return false
	}
}

// String returns the string representation of the AppEnv value.
func (e AppEnv) String() string {
	return string(e)
}
