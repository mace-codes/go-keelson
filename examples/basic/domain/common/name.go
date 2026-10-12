package common

import (
	"fmt"
)

// Name represents a validated name.
type Name string

// ErrInvalidName is returned when a name is invalid.
var ErrInvalidName = fmt.Errorf("invalid name")

// ParseName validates and parses a name string into a Name type.
func ParseName(name string) (Name, error) {
	if name == "" {
		return "", ErrInvalidName
	}

	return Name(name), nil
}

// String returns the string representation of the Name.
func (n Name) String() string {
	return string(n)
}

// Equals checks if two Name values are equal.
func (n Name) Equals(other Name) bool {
	return n == other
}
