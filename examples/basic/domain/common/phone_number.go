package common

import (
	"fmt"

	"github.com/nyaruka/phonenumbers"
)

// PhoneNumber represents a validated phone number.
type PhoneNumber string

// ErrInvalidPhoneNumber is returned when a phone number is invalid.
var ErrInvalidPhoneNumber = fmt.Errorf("invalid phone number")

// ParsePhoneNumber validates and parses a phone number string into a PhoneNumber type.
func ParsePhoneNumber(phone string) (PhoneNumber, error) {
	if phone == "" {
		return "", ErrInvalidPhoneNumber
	}

	pn, err := phonenumbers.Parse(phone, "US")
	if err != nil {
		return "", ErrInvalidPhoneNumber
	}

	if !phonenumbers.IsValidNumber(pn) {
		return "", ErrInvalidPhoneNumber
	}

	return PhoneNumber(phonenumbers.Format(pn, phonenumbers.E164)), nil
}

// String returns the string representation of the PhoneNumber.
func (p PhoneNumber) String() string {
	return string(p)
}

// Equals checks if two PhoneNumber values are equal.
func (p PhoneNumber) Equals(other PhoneNumber) bool {
	return p == other
}
