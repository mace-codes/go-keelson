package command

import (
	"github.com/mace-codes/go-keelson/internal/customer/application/command/ports"
)

type customerCommandHandler struct {
	insertCustomerProfile ports.ForInsertingCustomerProfile
}

// NewCustomerCommandHandler creates a new instance of customerCommandHandler.
func NewCustomerCommandHandler(
	insertCustomerProfile ports.ForInsertingCustomerProfile,
) *customerCommandHandler {
	return &customerCommandHandler{
		insertCustomerProfile: insertCustomerProfile,
	}
}
