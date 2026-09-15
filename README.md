# go-keelson

**A Go service template for teams who want hexagonal architecture done right — not just diagrammed, but enforced.**

Keelson gives you a working, DDD-informed hexagonal (ports & adapters) service skeleton with the boundaries that usually get skipped: explicit bounded-context isolation, anti-corruption layers at every seam, and a dependency-injected composition root that keeps `main.go` boring on purpose.

Clone it and ship, or read it as a reference for structuring your own service — it's built to hold up either way.

---

## Why this exists

Most "clean architecture" templates nail the layering and stop there. The failure mode that actually costs teams money isn't bad layering — it's **boundary erosion**: contexts quietly importing each other's domain types, interfaces defined wherever's convenient, and no enforced line between "our model" and "everyone else's."

Keelson exists to make that boundary a compiler error, not a code-review opinion.

## What's inside

**Hexagonal core** — Domain logic has zero knowledge of HTTP, SQL, or any transport. Ports are interfaces owned by the *application* layer (the consumer), not infrastructure — infrastructure only ever implements, never defines.

**DDD tactical patterns** — Entities, value objects, and aggregates live in an isolated `domain/` package per bounded context, free of framework or persistence concerns.

**Anti-corruption layers, everywhere they're needed** — Every boundary a context doesn't own — an external API, another bounded context — gets an explicit ACL. Foreign types stop at the ACL; they never reach the domain.

**Composition root, not a junk drawer** — `main.go` calls `app.Run`. `app.Run` loads config, wires dependencies, mounts routes, and listens with graceful shutdown. Nothing else. Wiring lives in one auditable place.

**Enforced, not just documented** — Package boundaries are structured to be enforceable in CI (`internal/` visibility, boundary linting), so the architecture survives contact with a deadline.

## Architecture at a glance
```
main.go
  └── app.Run                composition root: config → deps → routes → serve

internal/contexts/<name>/
  ├── domain/                entities, value objects, aggregates — no external deps
  ├── application/
  │   ├── ports/              interfaces, owned by the consumer
  │   ├── command/            write use cases
  │   ├── query/              read use cases
  │   └── acl/                anti-corruption layer for cross-context calls
  └── infrastructure/
      ├── repository/         implements the ports — persistence
      └── http/               implements the ports — transport
```


Dependencies point one direction only: **infrastructure → application → domain.** The domain never imports outward, and no context imports another context's `domain/` package directly — only its published contracts, through an ACL.

## Getting started

```bash
git clone https://github.com/mace-codes/go-keelson.git
cd go-keelson
go mod tidy
go run ./cmd/service
```

Use it as your starting point — rename the example context, add your own — or read through `internal/contexts/` as a reference implementation before structuring your own service.

## Design principles

1. **Ports belong to consumers.** An interface is defined where it's used, never by the thing implementing it.
2. **Nothing crosses a boundary untranslated.** External systems and other bounded contexts are only ever seen through an ACL.
3. **The composition root is the only place that knows everything.** Domain and application code stay ignorant of how they're wired together.
4. **If it's a rule, it's enforced.** Boundaries are structural (`internal/` packages, lint rules in CI) — not tribal knowledge that erodes under deadline pressure.

## Status

Actively maintained as a working template. Contributions and architectural discussion welcome via issues.

## License

Apache 2.0 — see [LICENSE](./LICENSE).