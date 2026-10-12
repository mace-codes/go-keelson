package customer

import (
	"github.com/mace-codes/go-keelson/domain/common"
)

// UnvalidatedProfile represents a customer profile that has not been validated yet.
type UnvalidatedProfile struct {
	FirstName string
	LastName  string
	Phone     string
	State     string
}

// Profile represents a validated customer profile.
type Profile struct {
	ID        ID
	FirstName common.Name
	LastName  common.Name
	Phone     common.PhoneNumber
	State     common.State
}

// NewProfile creates a new validated Profile from an UnvalidatedProfile.
func NewProfile(unvalidated UnvalidatedProfile) (Profile, error) {

	newID := GenerateID()

	// Validation logic would go here

	return Profile{
		ID:        newID,
		FirstName: common.Name(unvalidated.FirstName),
		LastName:  common.Name(unvalidated.LastName),
		Phone:     common.PhoneNumber(unvalidated.Phone),
		State:     common.State(unvalidated.State),
	}, nil
}
