package app

import (
	"github.com/mace-codes/go-keelson/internal/app/config"
)

// startable defines an interface for components that can be started. It requires a Start method that returns an error if the component fails to start.
type startable interface {
	Start() error
}

// stoppable defines an interface for components that can be stopped. It requires a Stop method that returns an error if the component fails to stop.
type stoppable interface {
	Stop() error
}

// dependencies holds the application's dependencies, such as database connections, external services, etc.
type dependencies struct {
	cfg *config.Config // Configuration for the application always available in the dependencies struct
	// Add fields for your dependencies here, e.g., database connections, external services, etc.

	startables []startable
	stopables  []stoppable
}

// initializeDependencies sets up the necessary dependencies for the application based on the provided configuration.
func initializeDependencies(cfg *config.Config) (*dependencies, error) {

	// Build and initialize your dependencies here based on the configuration.

	return &dependencies{
		cfg: cfg,

		// Initialize your dependencies here based on the configuration.

		startables: []startable{},
		stopables:  []stoppable{},
	}, nil
}
