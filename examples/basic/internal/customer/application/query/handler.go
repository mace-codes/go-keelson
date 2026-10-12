package query

import "github.com/mace-codes/go-keelson/examples/basic/internal/customer/application/query/ports"

type customerQueryHandler struct {
	getCustomerProfileByID ports.ForGettingCustomerProfileByID
}

// NewCustomerQueryHandler creates a new instance of customerQueryHandler.
func NewCustomerQueryHandler(
	getCustomerProfileByID ports.ForGettingCustomerProfileByID,
) *customerQueryHandler {
	return &customerQueryHandler{
		getCustomerProfileByID: getCustomerProfileByID,
	}
}
