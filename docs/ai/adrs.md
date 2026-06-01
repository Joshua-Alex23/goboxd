# Architectural Decision Records

## ADR-001: Use YAML For Language Definitions

### Context

The system must support multiple programming languages.

### Options Considered

1. Hardcoded Go switch statements
2. External YAML configuration

### Decision

Use YAML configuration.

### Rationale

Adding languages becomes a configuration change rather than a code change. While also going along with instructions of the hackathon.

---

## ADR-002: Use Temporary Directories Per Execution

### Context

Each execution creates source files and artifacts.

### Options Considered

1. Shared execution directory
2. Per-request temporary directory

### Decision

Use temporary directories.

### Rationale

Avoids collisions between executions and simplifies cleanup. And creates much needed isolation.

---

## ADR-003: Separate Validation From Execution

### Context

Requests require validation before running code.

### Options Considered

1. Validation inside executor
2. Dedicated validation package

### Decision

Dedicated validation package.

### Rationale

Keeps execution logic focused only on build and run operations.

---

## ADR-004: Queue-Based Concurrency Control

### Context

Multiple users may submit code simultaneously.

### Options Considered

1. Unlimited concurrent execution
2. Controlled concurrency using queue slots

### Decision

Use queue-based concurrency control.

### Rationale

Protects system resources and prevents overload.