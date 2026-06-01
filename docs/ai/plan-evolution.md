# Plan Evolution

## 2026-05-24 Initial Design

### What we thought we'd do

Create a simple HTTP server that directly executes commands using Go's os/exec package.

### What we actually did

Introduced a dedicated executor package and YAML-driven language configuration.

### Why it changed

The specification required plug-and-play language support. Hardcoded language definitions would not scale.

---

## 2026-05-27 Language Support

### What we thought we'd try out

Supported only Python initially.

### What we then actually did

Implemented both Python and C++ support.

### Why it changed

Stage 1 requires one interpreted language and one compiled language.

---

## 2026-05-30 Execution Design

### What we thought we could do

Execute programs directly from the current working directory.

### What we actually did

Create isolated temporary directories for each submission.

### Why it changed

This prevents conflicts between concurrent executions and simplifies cleanup.

---

## 2026-05-31 Artifact Execution Fix

### What we thought we could do

Execute compiled binaries using generated paths.

### What we actually did

Adjusted artifact path handling and execution working directory.

### Why it changed

Compiled binaries were being generated correctly but execution paths were incorrect, causing empty outputs.