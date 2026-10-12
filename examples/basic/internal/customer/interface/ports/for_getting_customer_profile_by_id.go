package ports

import (
	"github.com/mace-codes/go-keelson/examples/basic/domain/customer"
)

// ForGettingCustomerProfileByID defines the driving port for getting a customer profile by ID.
type ForGettingCustomerProfileByID func(id string) (customer.Profile, error)
