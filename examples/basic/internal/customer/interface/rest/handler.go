package rest

import (
	"net/http"

	"github.com/mace-codes/go-keelson/examples/basic/internal/customer/interface/ports"
)

// CustomerRestHandler is a REST handler for the customer bounded context.
type CustomerRestHandler struct {
	CreateCustomerProfile  http.HandlerFunc
	GetCustomerProfileByID http.HandlerFunc
}

// NewCustomerRestHandler creates a new instance of CustomerRestHandler.
func NewCustomerRestHandler(
	createCustomerProfile ports.ForCreatingNewCustomerProfile,
	getCustomerProfile ports.ForGettingCustomerProfileByID,
) *CustomerRestHandler {
	return &CustomerRestHandler{
		CreateCustomerProfile:  CreateCustomerProfile(createCustomerProfile),
		GetCustomerProfileByID: GetCustomerProfileByID(getCustomerProfile),
	}
}
