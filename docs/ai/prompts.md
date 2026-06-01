# AI Usage Log

## 2026-05-29 · Designing the Language Registry

**Prompt:**

How should I structure a Go service that supports multiple programming languages with different build and run commands without hardcoding everything?

**Response summary:**

Suggested creating a configuration-driven design where language definitions are stored externally and loaded into a registry at startup.

**What we used / didn't use:**

Used the configuration-driven registry idea and implemented language definitions in YAML. Did not use a hardcoded switch-based approach.

---

## 2026-05-29 · Request Validation Strategy

**Prompt:**

Where should validation happen in a code execution service? Should validation be done inside the executor or before execution starts?

**Response summary:**

Recommended separating validation from execution so invalid requests are rejected before any build or run operation begins.

**What we used / didn't use:**

Created a dedicated validation package and kept execution logic focused only on build and run responsibilities.

---

## 2026-05-29 · Temporary Directory Isolation

**Prompt:**

How can I safely handle multiple code submissions without file conflicts between executions?

**Response summary:**

Suggested creating a unique temporary working directory for every execution and removing it afterwards.

**What we used / didn't use:**

Implemented per-request temporary directories using Go's os.MkdirTemp and cleanup with defer.

---

## 2026-05-29 · Compiler Flag Security

**Prompt:**

Users can provide compiler flags. How do I prevent dangerous compiler arguments from being passed to the build process?

**Response summary:**

Suggested maintaining an allowlist of approved compiler flags and rejecting everything else.

**What we used / didn't use:**

Implemented flag validation and compiler flag allowlists for supported languages.

---

## 2026-05-29 · C++ Compilation Succeeds But Output Is Empty

**Prompt:**

My C++ code compiles successfully but execution returns no output. How should I debug this?

**Response summary:**

Suggested verifying artifact paths, working directories, and how executable paths are generated after compilation.

**What we used / didn't use:**

Used the debugging approach to trace the problem. Discovered the compiled artifact was not being executed from the correct location.

---

## 2026-05-30 · Executable Artifact Path Bug

**Prompt:**

The compiler produces the binary successfully but the runner cannot find it. What could cause this?

**Response summary:**

Explained that relative executable paths can break when the working directory differs from the artifact location.

**What we used / didn't use:**

Updated artifact handling and execution paths so compiled programs are executed from the correct directory.

---

## 2026-05-30 · NSJail Integration Debugging

**Prompt:**

NSJail is reporting "No such file or directory" even though the executable exists. How can I investigate this?

**Response summary:**

Suggested checking chroot behavior, filesystem mounts, executable visibility, and whether the process can access required binaries.

**What we used / didn't use:**

Used the debugging process to identify filesystem visibility issues. Deferred full sandbox integration until the execution pipeline was stable.

---

## 2026-05-31 · Handling Compilation Errors Correctly

**Prompt:**

How should compilation failures be represented in API responses?

**Response summary:**

Recommended treating build failures separately from successful executions and surfacing compiler stderr to the client.

**What we used / didn't use:**

Updated execution handling to return compiler diagnostics instead of silently treating failures as empty output.

---

## 2026-05-31 · Determining Test Status

**Prompt:**

How should test status be calculated when stderr exists but stdout matches expected output?

**Response summary:**

Suggested treating stderr as a failure condition because successful output does not necessarily mean successful execution.

**What we used / didn't use:**

Updated status calculation to fail whenever execution or compilation produced stderr.

---

## 2026-06-01 · Debugging Incorrect Runtime Results

**Prompt:**

The API returns wrong_output even though compilation succeeds. What areas should I inspect first?

**Response summary:**

Suggested checking execution paths, working directories, command rendering, and runtime argument construction.

**What we used / didn't use:**

Used this approach to identify a path-generation issue that prevented the compiled executable from being launched correctly.

---

## 2026-06-01 · Stage 1 Readiness Review

**Prompt:**

Review the implementation against the hackathon Stage 1 specification and identify missing requirements.

**Response summary:**

Compared the implementation against the Stage 1 requirements including health checks, interpreted language support, compiled language support, testing, documentation, and repository structure.

**What we used / didn't use:**

Used the review checklist to verify completion status and identify documentation work that remained.
