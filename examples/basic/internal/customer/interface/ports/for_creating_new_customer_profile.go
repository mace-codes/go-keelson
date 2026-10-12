package ports

import (
	"context"

	"github.com/mace-codes/go-keelson/domain/customer"
)

// ForCreatingNewCustomerProfile defines the driving port for creating a new customer profile.
type ForCreatingNewCustomerProfile func(ctx context.Context, profile customer.UnvalidatedProfile) error
