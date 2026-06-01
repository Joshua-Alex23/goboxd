# Architecture

## Overview

goboxd is a Go HTTP service that accepts source code submissions, executes them in a controlled environment, and evaluates them against one or more test cases.

The system is divided into four main layers:

1. HTTP API layer
2. Configuration layer
3. Execution layer
4. Validation layer

---

## Request Flow

1. Client sends a POST request to `/run`.
2. The API validates the request payload.
3. The language configuration is loaded from the registry.
4. A temporary working directory is created.
5. Source code is written to disk.
6. If the language requires compilation, the build step is executed.
7. The program is executed with the supplied stdin.
8. Output is compared against the expected output.
9. Results are returned as JSON.

---

## Language Configuration

Languages are defined in `configs/languages.yaml`.

Each language specifies:

- Source filename
- Optional build step
- Run command
- Resource limits
- Allowed compiler flags

Example:

- Python uses only a run step.
- C++ uses a build step followed by execution of the generated artifact.

This allows new languages to be added without changing application code.

---

## Validation

Validation is separated from execution.

Current validation includes:

- Request validation
- Filename validation
- Compiler flag allowlist validation

Validation occurs before any code is executed.

---

## Execution Model

Execution is handled by the executor package.

For every request:

- A temporary directory is created.
- Source code is written into the directory.
- Compilation is performed when required.
- Execution occurs inside the temporary directory.
- Results are collected.
- Temporary files are removed after completion.

Execution time is limited using Go contexts and timeouts.

---

## Concurrency Control

The service uses a queue package to limit concurrent executions.

Before a run starts:

- A slot is acquired.

After completion:

- The slot is released.

This prevents excessive resource usage from many simultaneous requests.

---

## Error Handling

Compilation failures and runtime failures are returned to the client through stderr.

Supported result states currently include:

- accepted
- wrong_output
- runtime_error

---

## Health Monitoring

The service exposes:

- `/healthz`
- `/readyz`
- `/info`

These endpoints allow external systems to verify service availability.