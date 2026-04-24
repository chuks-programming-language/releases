#!/bin/bash
# Fair benchmark harness: warmup + best-of-5 per language, wall-clock ms
set -e
cd "$(dirname "$0")"

BENCHMARKS=(
    "fib_benchmark:fib:Fib:Fibonacci (fib(38))"
    "matrix_multiply:matrix:Matrix:Matrix Multiply (200x200)"
    "prime_sieve:sieve:Sieve:Prime Sieve (1M)"
    "quicksort:sort:Sort:Quicksort (100K elements)"
    "string_concat:strings:Strings:String Concatenation (100K)"
    "binary_trees:trees:Trees:Binary Trees (depth 16)"
    "nbody:nbody:Nbody:N-Body Simulation (500K steps)"
    "oop_benchmark:oop:Oop:OOP Class Hierarchy (100K objects)"
    "map_operations:maps:Maps:Map Operations (100K entries)"
)

REPORT="benchmark_report.txt"
: > "$REPORT"

{
    echo "================================================================"
    echo "  Chuks — Benchmark (warmup + best-of-5, ms)"
    echo "  Date: $(date)"
    echo "  Platform: $(uname -sm)"
    echo "================================================================"
    echo
} >> "$REPORT"

bench() {
    local name="$1"; shift
    python3 - "$@" <<'PY' "$name"
import subprocess, time, sys
args = sys.argv[1:-1]
name = sys.argv[-1]
try:
    subprocess.run(args, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL, check=False)
except FileNotFoundError:
    print(f"  {name:<18}        n/a")
    sys.exit(0)
best = float("inf")
for _ in range(5):
    t0 = time.perf_counter()
    r = subprocess.run(args, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
    dt = (time.perf_counter() - t0) * 1000.0
    if r.returncode != 0:
        print(f"  {name:<18}   FAILED")
        sys.exit(0)
    if dt < best:
        best = dt
print(f"  {name:<18} {best:10.2f} ms")
PY
}

for entry in "${BENCHMARKS[@]}"; do
    IFS=':' read -r dir basename javaname display <<< "$entry"
    echo "────────────────────────────────────────────────────────────────" >> "$REPORT"
    echo "  $display" >> "$REPORT"
    echo "────────────────────────────────────────────────────────────────" >> "$REPORT"
    pushd "$dir" > /dev/null

    chuks build "$basename.chuks" > /dev/null 2>&1 || true
    if [[ -f "build/$basename.chuks.bin" ]]; then
        xattr -cr "build/$basename.chuks.bin" 2>/dev/null || true
        codesign --force --sign - "build/$basename.chuks.bin" 2>/dev/null || true
    fi
    go build -o "${basename}_go" "$basename.go" 2>/dev/null || true
    if [[ -f "${basename}_go" ]]; then
        xattr -cr "${basename}_go" 2>/dev/null || true
        codesign --force --sign - "${basename}_go" 2>/dev/null || true
    fi
    javac "$javaname.java" 2>/dev/null || true

    {
        bench "Chuks VM"    chuks run "$basename.chuks"
        [[ -x "build/$basename.chuks.bin" ]] && bench "Chuks AOT"   "./build/$basename.chuks.bin"
        [[ -x "${basename}_go" ]]            && bench "Go"          "./${basename}_go"
        [[ -f "$javaname.class" ]]           && bench "Java"        java "$javaname"
        bench "Node.js"     node "$basename.js"
        bench "Bun"         bun "$basename.js"
        bench "Python"      python3 "$basename.py"
    } >> "../$REPORT"

    echo >> "../$REPORT"
    popd > /dev/null
done

echo "================================================================" >> "$REPORT"
echo "  Fair benchmark complete." >> "$REPORT"
echo "================================================================" >> "$REPORT"
cat "$REPORT"
