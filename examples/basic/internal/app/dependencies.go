package app

import (
	"context"

	"github.com/mace-codes/go-keelson/domain/customer"
	"github.com/mace-codes/go-keelson/internal/app/config"
	"go.uber.org/zap"

	transport "github.com/mace-codes/go-keelson-transport"
	cstCmd "github.com/mace-codes/go-keelson/internal/customer/application/command"
	cstQry "github.com/mace-codes/go-keelson/internal/customer/application/query"
	cstRest "github.com/mace-codes/go-keelson/internal/customer/interface/rest"
)

// NoOps used for now, but you can replace it with actual implementations as needed.
func insertNoop(ctx context.Context, prof customer.Profile) error {
	return nil
}

func getCustomerNoop(id customer.ID) (customer.Profile, error) {
	return customer.Profile{}, nil
}

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
	cfg    *config.Config // Configuration for the application always available in the dependencies struct
	logger *zap.Logger

	// Add fields for your dependencies here, e.g., database connections, external services, etc.

	// REST Handlers
	customerRestHandler *cstRest.CustomerRestHandler // Example REST handler for the customer bounded context

	startables []startable
	stopables  []stoppable
}

// Critical returns a list of service critical health reporters
func (d *dependencies) Critical() []transport.Reporter {
	return []transport.Reporter{}
}

// Optional returns a list of service optional health reporters
func (d *dependencies) Optional() []transport.Reporter {
	return []transport.Reporter{}
}

// initializeDependencies sets up the necessary dependencies for the application based on the provided configuration.
func initializeDependencies(cfg *config.Config) (*dependencies, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	// Build and initialize your dependencies here based on the configuration.

	customerCmdHandler := cstCmd.NewCustomerCommandHandler(
		insertNoop,
	)

	customerQryHandler := cstQry.NewCustomerQueryHandler(
		getCustomerNoop,
	)

	cstRestHandler := cstRest.NewCustomerRestHandler(
		customerCmdHandler.CreateCustomerProfile,
		customerQryHandler.GetCustomerProfileByID,
	)

	return &dependencies{
		cfg:    cfg,
		logger: logger,

		// Initialize your dependencies here based on the configuration.

		customerRestHandler: cstRestHandler,

		startables: []startable{},
		stopables:  []stoppable{},
	}, nil
}
