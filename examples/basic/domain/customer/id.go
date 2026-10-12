package customer

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

// ID represents a validated customer ID.
type ID string

// ErrInvalidID is returned when a customer ID is invalid.
var ErrInvalidID = fmt.Errorf("invalid customer ID")

var regexUUID = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`)

// GenerateID generates a new unique customer ID.
func GenerateID() ID {
	return ID(uuid.NewString())
}

// ParseID validates and parses a customer ID string into an ID type.
func ParseID(id string) (ID, error) {
	if id == "" {
		return "", ErrInvalidID
	}

	if !regexUUID.MatchString(id) {
		return "", ErrInvalidID
	}

	return ID(id), nil
}

// String returns the string representation of the ID.
func (id ID) String() string {
	return string(id)
}

// Equals checks if two ID values are equal.
func (id ID) Equals(other ID) bool {
	return id == other
}
