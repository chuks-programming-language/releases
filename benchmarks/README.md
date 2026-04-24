# Chuks Benchmarks

Cross-language micro-benchmarks and HTTP throughput tests comparing **Chuks** against **Go**, **Java**, **Node.js**, **Bun**, and **Python**. All benchmarks are single-file, single-purpose, and use identical algorithms in every language so results reflect runtime/codegen differences — not implementation choices.

These are the exact scripts we use in our own CI to track regressions. Everyone is welcome to run them locally and report results.

## Requirements

Install whichever runtimes you want to compare. Missing ones are skipped automatically.

| Tool       | Install                                                   |
|------------|-----------------------------------------------------------|
| `chuks`    | `curl -fsSL https://chuks.org/install.sh \| bash`         |
| `go`       | https://go.dev/dl/                                        |
| `java`/`javac` | `brew install openjdk` / distro JDK                   |
| `node`     | https://nodejs.org                                        |
| `bun`      | `curl -fsSL https://bun.sh/install \| bash`               |
| `python3`  | Usually preinstalled                                      |
| `wrk`      | `brew install wrk` / `apt install wrk` (HTTP test only)   |

## Quick start

```bash
git clone https://github.com/chuks-programming-language/releases.git
cd releases/benchmarks

# Compute benchmarks (9 benchmarks × up to 7 runtimes, best-of-5)
./run_all_benchmarks.sh

# HTTP throughput (wrk -t4 -c100 -d10s, GET /)
./run_http_benchmark.sh
```

Compute results are written to `benchmark_report.txt` and printed on the console. HTTP results are printed as a summary table.

## Methodology

The harness is designed to be **fair** and reproducible:

1. **Warmup run** before timing to page in binaries, load caches, and let JITs spin up.
2. **Best-of-5 timing** per benchmark/runtime. Takes the fastest of 5 wall-clock runs. Smooths out OS jitter without rewarding outliers.
3. **macOS Gatekeeper mitigation** — `xattr -cr` and `codesign --force --sign -` on each freshly-built binary, so first-run quarantine latency doesn't get charged to whatever language built last.
4. **Same algorithm** across all languages. Check `*/fib.go` vs `*/fib.chuks` vs `*/Fib.java` — you will see the code is the same shape in every language.
5. **Wall-clock from outside the process** (Python `perf_counter`), not internal timers. Includes startup cost, which is part of real-world performance.

## What's measured

### Compute (`run_all_benchmarks.sh`)

| Benchmark       | Problem Size          | Stresses                       |
|-----------------|-----------------------|--------------------------------|
| Fibonacci       | `fib(38)`             | recursive call dispatch        |
| Matrix Multiply | 200×200 × 200×200     | nested loops, double math      |
| Prime Sieve     | 1,000,000             | boolean array access           |
| Quicksort       | 100,000 ints          | recursion + array swap         |
| String Concat   | 100,000 appends       | string buffer growth           |
| Binary Trees    | depth 16              | allocation + GC pressure       |
| N-Body          | 500,000 steps         | float math, loop hoisting      |
| OOP Hierarchy   | 100,000 objects       | virtual dispatch, inheritance  |
| Map Operations  | 100,000 entries       | hash ops (insert/lookup/del)   |

### HTTP (`run_http_benchmark.sh`)

A minimal `"Hello, World!"` endpoint. Measures raw request throughput and average latency under `wrk -t4 -c100 -d10s`. Compares Chuks's `std/http` (built on `gnet` event loop) against Go stdlib `net/http`, raw `gnet+ants`, Java `HttpServer`, Bun, Node.js, and Python's `http.server`.

## Reporting your results

We'd love to see numbers from other hardware. Open an issue at:
https://github.com/chuks-programming-language/releases/issues

Please include:

- Output of `benchmark_report.txt` and the HTTP summary table
- `uname -srm`
- CPU model and core count (`sysctl -n machdep.cpu.brand_string` on macOS, `lscpu` on Linux)
- Runtime versions (`chuks --version`, `go version`, `java --version`, `node --version`, `bun --version`, `python3 --version`)

## Notes

- **Python's `http.server` is single-threaded** and has no connection pool. It will timeout under `wrk`'s load (few requests served). That's expected — it's a reference, not a contender.
- **Java and Go `net/http`** pay a cold-start tax on the first request. This is visible in `wrk`'s first-second latency but is amortized across the 10s window.
- **Chuks AOT binaries** are cached in `$HOME/.chuks/cache/builds/` keyed by source path. Delete that directory to force a fresh transpile + `go build`.

## License

The benchmark sources are released into the public domain (CC0). Port them, modify them, publish your own numbers.
