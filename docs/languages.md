# Languages

Languages are configured through `configs/languages.yaml`.

The service loads language definitions during startup and registers them in memory.

Adding support for a new language requires only a new configuration entry and the corresponding toolchain inside the Docker image.

---

# Python 3

Identifier:

```yaml
id: py3
```

Source file:

```text
solution.py
```

Execution command:

```text
python3 solution.py
```

Capabilities:

* Interpreted language
* No build step required
* Supports stdin input
* Supports multiple test cases

Example:

```python
name = input()
print(name)
```

---

# C++

Identifier:

```yaml
id: cpp
```

Source file:

```text
solution.cpp
```

Build command:

```text
g++ solution.cpp -o solution
```

Execution command:

```text
./solution
```

Capabilities:

* Compiled language
* Build step required before execution
* Supports stdin input
* Supports multiple test cases

Example:

```cpp
#include <iostream>

int main() {
    std::cout << "hello";
}
```

---

# Language Registry

Each language definition contains:

```yaml
id:
name:
source_filename:

build:
  cmd:
  args:

run:
  cmd:
  args:
```

The registry is loaded when the service starts.

Invalid language definitions prevent startup.

---

# Adding a New Language

To add a new language:

1. Install the compiler/interpreter inside the Docker image.
2. Add a new entry to `configs/languages.yaml`.
3. Rebuild the container.
4. Restart the service.

No application code changes are required.
