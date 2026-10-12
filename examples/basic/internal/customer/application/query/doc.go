// application/query/doc.go

// Package query contains read-side use cases: handlers that return data
// and never change state.
//
// Each query is a plain struct of criteria (GetOrder, ListOrders) paired
// with a handler returning a purpose-built read model (DTO), not a domain
// entity. Handlers may read through dedicated read-store ports, bypassing
// aggregates and repositories when that is simpler or faster.
//
// Queries must be free of side effects and safe to retry or cache. They
// contain no business rules; if one is needed, it belongs in the domain.
//
// Depends on the domain and on ports only, never on infrastructure.
package query
