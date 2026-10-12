// application/command/doc.go

// Package command contains write-side use cases: handlers that change
// system state.
//
// Each command is a plain struct of intent (CreateOrder, CancelOrder) paired
// with a handler that:
//   - validates input and loads aggregates through repository ports,
//   - invokes domain behaviour, where the business rules live,
//   - persists the result and publishes domain events.
//
// Handlers own the transaction boundary. They return an error or, at most,
// a newly created identifier; they never return read models. Callers that
// need data afterwards use the query package.
//
// Depends on the domain and on ports only, never on infrastructure.
package command
