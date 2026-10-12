package query

import (
	"github.com/mace-codes/go-keelson/domain/customer"
	"github.com/mace-codes/go-keelson/internal/customer/interface/ports"
)

// Ensure that customerQueryHandler implements the ForGettingCustomerProfileByID interface.
var _ ports.ForGettingCustomerProfileByID = (*customerQueryHandler)(nil).GetCustomerProfileByID

// GetCustomerProfileByID retrieves a customer profile by its ID.
func (h *customerQueryHandler) GetCustomerProfileByID(id string) (customer.Profile, error) {

	valiidatedID, err := customer.ParseID(id)
	if err != nil {
		return customer.Profile{}, err
	}

	return h.getCustomerProfileByID(valiidatedID)
}
