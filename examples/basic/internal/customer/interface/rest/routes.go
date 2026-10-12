package rest

import (
	"net/http"

	transport "github.com/mace-codes/go-keelson-transport"
)

// Routes returns the route configurations for the customer bounded context.
func Routes(h *CustomerRestHandler) []transport.RoutesConfig {
	return []transport.RoutesConfig{
		{
			Methods:  []string{http.MethodPost},
			Path:     "/customer",
			Handler:  h.CreateCustomerProfile,
			Adapters: nil, // You can add any necessary adapters here if needed
		},
		{
			Methods:  []string{http.MethodGet},
			Path:     "/customers/{id}",
			Handler:  h.GetCustomerProfileByID,
			Adapters: nil, // You can add any necessary adapters here if needed
		},
	}
}
