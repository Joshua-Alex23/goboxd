# Security

## Overview

goboxd executes user-supplied code and therefore applies several safeguards to reduce risk and prevent abuse.

The current implementation focuses on the requirements of Stage 1.

---

## Execution Isolation

Each request receives a dedicated temporary working directory.

Source files and build artifacts are written only inside that directory.

The directory is removed after execution completes.

Benefits:

* Prevents collisions between concurrent executions.
* Ensures artifacts from one request are not reused by another.
* Simplifies cleanup.

---

## Filename Validation

Language configuration values such as:

```text
source_filename
artifact
```

are validated before use.

This prevents invalid or dangerous paths from being used during execution.

Examples of rejected values:

```text
../../etc/passwd
/tmp/file
../solution
```

---

## Execution Timeout

Programs are executed using a context with a fixed timeout.

If execution exceeds the allowed time, the process is terminated.

Current timeout:

```text
2 seconds
```

Purpose:

* Prevent infinite loops.
* Prevent long-running programs from occupying execution slots indefinitely.

---

## Output Limits

Program output is truncated to a maximum size.

Current limit:

```text
1 MB
```

Purpose:

* Prevent excessive memory usage.
* Prevent oversized API responses.

---

## Request Concurrency

The service uses a bounded execution queue.

Requests must acquire a slot before execution begins.

Benefits:

* Prevents unbounded resource consumption.
* Provides predictable server behaviour under load.

---

## Container Boundary

All executions occur inside the Docker container running goboxd.

Language toolchains and runtime dependencies are installed inside the container image.

This allows the application to run consistently across environments.

---

## Current Limitations

The current implementation is intended for Stage 1 of the project.

Areas planned for future hardening include:

* Stronger process isolation
* Memory limits
* Process count limits
* Additional filesystem restrictions
* Enhanced sandboxing controls

These improvements are expected in later stages of the project.
