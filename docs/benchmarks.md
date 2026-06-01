# Benchmarks

## Environment

Benchmarks were executed inside Docker using the project runtime image.

Hardware used during development:

* CPU: Apple M2
* Memory: 8GB RAM
* OS: Mac OS Sonoma
* Docker: Docker Compose v2

## Methodology

Load testing was performed by issuing concurrent requests to the `/run` endpoint using representative Python and C++ programs.

Measurements focused on:

* Request latency
* Successful execution rate
* Service stability under concurrent load

## Results

The service remained responsive during concurrent execution tests and continued serving health checks through `/healthz`.

Representative workloads:

| Language | Program Type      | Result             |
| -------- | ----------------- | ------------------ |
| Python   | Hello World       | Passed             |
| Python   | Infinite Loop     | Timed Out          |
| C++      | Hello World       | Passed             |
| C++      | Compilation Error | Correctly Reported |

## Notes

The benchmark suite is intended to verify correctness and service stability rather than establish absolute throughput numbers. Performance depends on host hardware, container runtime configuration, and resource limits.
