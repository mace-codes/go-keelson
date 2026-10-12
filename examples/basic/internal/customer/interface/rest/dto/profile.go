package dto

import (
	"github.com/mace-codes/go-keelson/domain/customer"
)

// Profile represents the data transfer object for a customer profile.
type Profile struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	State     string `json:"state"`
}

// DTOToUnvalidatedProfile converts a Profile DTO to a domain customer.UnvalidatedProfile.
func DTOToUnvalidatedProfile(dto Profile) customer.UnvalidatedProfile {
	return customer.UnvalidatedProfile{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Phone:     dto.Phone,
		State:     dto.State,
	}
}

// DomainProfileToDTO converts a domain customer.Profile to a Profile DTO.
func DomainProfileToDTO(domainProfile customer.Profile) Profile {
	return Profile{
		ID:        domainProfile.ID.String(),
		FirstName: domainProfile.FirstName.String(),
		LastName:  domainProfile.LastName.String(),
		Phone:     domainProfile.Phone.String(),
		State:     domainProfile.State.String(),
	}
}
