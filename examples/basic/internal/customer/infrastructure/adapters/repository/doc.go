// infrastructure/adapters/repository/doc.go

// Package repository contains persistence adapters implementing the
// repository ports defined in the domain layer.
//
// Responsibilities:
//   - Map between storage models and domain entities; storage tags and
//     schemas stay in this package.
//   - Translate driver errors (no rows, unique violation) into domain errors
//     such as common.ErrNotFound.
//   - Participate in transactions supplied via context or a unit-of-work
//     port, never by opening their own behind the caller's back.
//
// Business rules do not live here. If a query needs a decision, the
// decision belongs in the domain.
package repository
