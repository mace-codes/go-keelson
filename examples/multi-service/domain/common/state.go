package common

import (
	"fmt"
	"regexp"
)

// State represents a validated state.
type State string

// ErrInvalidState is returned when a state is invalid.
var ErrInvalidState = fmt.Errorf("invalid state")

var validStates = regexp.MustCompile(`^[A-Z]{2}$`)

// ParseState validates and parses a state string into a State type.
func ParseState(state string) (State, error) {
	if state == "" {
		return "", ErrInvalidState
	}

	if !validStates.MatchString(state) {
		return "", ErrInvalidState
	}

	return State(state), nil
}

// String returns the string representation of the State.
func (s State) String() string {
	return string(s)
}

// Equals checks if two State values are equal.
func (s State) Equals(other State) bool {
	return s == other
}
