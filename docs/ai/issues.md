# Issues Encountered

## 2026-05-30 Compiled Programs Produced Empty Output

### What we were trying to do

Run compiled C++ programs and capture stdout.

### What went wrong

Compilation succeeded but execution returned empty output.

### How we resolved it

Investigated artifact generation and execution paths. Updated execution to use the correct working directory and artifact path.

### What we learned

Successful compilation does not guarantee correct execution paths.

---

## 2026-05-31 Compiler Errors Were Reported As Accepted

### What we were trying to do

Return failed status when compilation errors occur.

### What went wrong

Compilation stderr was returned but test status logic still allowed acceptance under some conditions.

### How we resolved it

Updated status calculation to fail whenever stderr is present.

### What we learned

Status evaluation must consider build and runtime failures separately.

---

## 2026-05-31 NSJail Integration Investigation

### What we were trying to do

Execute programs through nsjail.

### What went wrong

nsjail could not locate executables because filesystem mounts and root configuration were incomplete.

### How we resolved it

Temporarily reverted to direct execution while identifying the correct nsjail configuration.

### What we learned

Sandbox integration should be validated independently before integrating into the execution pipeline.