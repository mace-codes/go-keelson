package ports

import "github.com/mace-codes/go-keelson/examples/multi-service/domain/customer"

// ForGettingCustomerProfileByID defines the driven port for getting a customer profile by ID.
type ForGettingCustomerProfileByID func(customer.ID) (customer.Profile, error)
