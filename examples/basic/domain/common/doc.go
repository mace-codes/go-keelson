// Package common holds the shared kernel of the domain layer: value objects,
// identifiers, domain errors, and small interfaces used by more than one
// aggregate.
//
// # Rules
//
//   - Pure domain only: no I/O, no framework, no infrastructure imports.
//   - Standard library only; never imports a sibling domain package.
//   - Admission is earned: a type lands here once two or more aggregates need it.
//
// # Contents
//
//   - Identifiers and value objects (ID, Money, Timestamp)
//   - Sentinel domain errors (ErrNotFound, ErrInvalid)
//   - Shared interfaces (Clock, IDGenerator) implemented by adapters
package common
