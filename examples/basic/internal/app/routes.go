package app

import (
	transport "github.com/mace-codes/go-keelson-transport"
	cstRest "github.com/mace-codes/go-keelson/internal/customer/interface/rest"
)

// serveRoutes takes in route groups exported by the rest interface of bounded contexts
func serveRoutes(deps *dependencies) ([]transport.RoutesConfig, error) {
	routes := []transport.RoutesConfig{} // Initialize an empty slice to hold the route configurations

	// Health check route is built-in and always available through the transport package. No need to add them here. They are automatically registered when you use the transport package.

	// Note: You can add more route groups here as needed. For example, if you have other bounded contexts with their own REST handlers, you can append their routes to the 'routes' slice.

	if deps.customerRestHandler != nil {
		routes = append(routes, cstRest.Routes(deps.customerRestHandler)...)
	}

	return routes, nil
}
