# API

## Overview

goboxd exposes a small HTTP API for compiling and executing code.

Base URL:

```text
http://localhost:8080
```

---

## Health Check

### GET /healthz

Returns the health status of the service.

### Response

Status Code:

```text
200 OK
```

Body:

```json
{
  "status": "ok"
}
```

---

## Execute Code

### POST /run

Compiles (if required) and executes source code against one or more test cases.

### Request

```json
{
  "language": "py3",
  "source": "print(input())",
  "tests": [
    {
      "stdin": "hello",
      "expected_stdout": "hello\n"
    }
  ]
}
```

### Fields

| Field    | Type   | Required | Description            |
| -------- | ------ | -------- | ---------------------- |
| language | string | Yes      | Language identifier    |
| source   | string | Yes      | Source code to execute |
| tests    | array  | Yes      | List of test cases     |

### Test Case

```json
{
  "stdin": "hello",
  "expected_stdout": "hello\n"
}
```

| Field           | Type   | Description                   |
| --------------- | ------ | ----------------------------- |
| stdin           | string | Input provided to the program |
| expected_stdout | string | Expected program output       |

---

## Successful Execution

```json
{
  "status": "accepted",
  "build": {
    "status": "ok",
    "stdout": "",
    "stderr": "",
    "duration_ms": 0
  },
  "tests": [
    {
      "status": "accepted",
      "stdout": "hello\n",
      "stderr": "",
      "duration_ms": 0
    }
  ]
}
```

---

## Wrong Output

```json
{
  "status": "wrong_output",
  "build": {
    "status": "ok"
  },
  "tests": [
    {
      "status": "wrong_output",
      "stdout": "actual output",
      "stderr": ""
    }
  ]
}
```

---

## Compilation Error

```json
{
  "status": "runtime_error",
  "build": {
    "status": "ok"
  },
  "tests": [
    {
      "status": "runtime_error",
      "stdout": "",
      "stderr": "compiler error output"
    }
  ]
}
```

---

## Unknown Language

Request:

```json
{
  "language": "rust",
  "source": "fn main(){}",
  "tests": [
    {
      "stdin": "",
      "expected_stdout": ""
    }
  ]
}
```

Response:

```json
{
  "error": {
    "code": "unknown_language",
    "message": "language not registered"
  }
}
```

---

## Status Values

### Top-Level Status

| Status        | Meaning                                     |
| ------------- | ------------------------------------------- |
| accepted      | All tests passed                            |
| wrong_output  | One or more tests produced incorrect output |
| runtime_error | Build or execution failed                   |

### Test Status

| Status        | Meaning                              |
| ------------- | ------------------------------------ |
| accepted      | Test passed                          |
| wrong_output  | Output did not match expected output |
| runtime_error | Program failed to build or execute   |
|               |                                      |
