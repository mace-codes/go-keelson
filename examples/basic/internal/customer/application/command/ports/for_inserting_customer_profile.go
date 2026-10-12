package ports

import (
	"context"

	"github.com/mace-codes/go-keelson/examples/basic/domain/customer"
)

// ForInsertingCustomerProfile defines the driven port for inserting a customer profile.
type ForInsertingCustomerProfile func(context.Context, customer.Profile) error
