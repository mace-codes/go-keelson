// infrastructure/adapters/client/doc.go

// Package client contains outbound adapters that call external services
// (HTTP, gRPC, APIs, or SDK-based) on behalf of the application layer.
//
// Each type here implements a port declared in the application or domain
// layer; the port is owned by the consumer, never by this package.
//
// Responsibilities:
//   - Translate between wire models and domain types, so transport DTOs
//     never leak inward.
//   - Map transport and status errors onto domain errors.
//   - Own timeouts, retries, and auth for the downstream service.
//
// Context is always the first parameter and is honoured for cancellation.
package client
