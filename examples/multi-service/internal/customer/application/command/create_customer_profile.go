package command

import (
	"context"

	"github.com/mace-codes/go-keelson/examples/multi-service/domain/customer"
	"github.com/mace-codes/go-keelson/examples/multi-service/internal/customer/interface/ports"
)

// Ensure that customerCommandHandler implements the ForCreatingNewCustomerProfile interface.
var _ ports.ForCreatingNewCustomerProfile = (*customerCommandHandler)(nil).CreateCustomerProfile

// CreateCustomerProfile creates a new customer profile.
func (h *customerCommandHandler) CreateCustomerProfile(ctx context.Context, profile customer.UnvalidatedProfile) error {

	newProfile, err := customer.NewProfile(profile)
	if err != nil {
		return err
	}

	return h.insertCustomerProfile(ctx, newProfile)
}
