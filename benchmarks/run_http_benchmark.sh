#!/bin/bash
# HTTP Server Benchmark — self-contained version for public use.
# Compares raw request throughput on a "Hello, World!" endpoint across runtimes.
#
# Requires: chuks, go, java/javac, bun, node, python3, wrk
# Usage: ./run_http_benchmark.sh

set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$SCRIPT_DIR"

WRK_THREADS=4
WRK_CONNS=100
WRK_DURATION=10s
PORT=9090
WARMUP=2
COOLDOWN=1

RED='\033[0;31m'; GREEN='\033[0;32m'; CYAN='\033[0;36m'; BOLD='\033[1m'; NC='\033[0m'

declare -a NAMES RPS_VALUES LATENCY_AVG TRANSFER

kill_port() {
    lsof -ti:$PORT 2>/dev/null | xargs kill -9 2>/dev/null || true
    sleep "$COOLDOWN"
}

wait_for_server() {
    local waited=0
    while ! curl -s "http://localhost:$PORT/" > /dev/null 2>&1; do
        sleep 0.2; waited=$((waited + 1))
        [ $waited -ge 50 ] && { echo -e "  ${RED}✗ server failed to start${NC}"; return 1; }
    done
    sleep "$WARMUP"; return 0
}

run_wrk() {
    local label="$1"
    local output; output=$(wrk -t"$WRK_THREADS" -c"$WRK_CONNS" -d"$WRK_DURATION" "http://localhost:$PORT/" 2>&1)
    echo "$output"
    NAMES+=("$label")
    RPS_VALUES+=("$(echo "$output" | grep 'Requests/sec:' | awk '{print $2}')")
    LATENCY_AVG+=("$(echo "$output" | grep 'Latency' | awk '{print $2}')")
    TRANSFER+=("$(echo "$output" | grep 'Transfer/sec:' | awk '{print $2}')")
    echo ""
}

have() { command -v "$1" > /dev/null 2>&1; }

if have chuks; then
    echo -e "${BOLD}▶ Chuks VM${NC}"
    kill_port
    chuks run bench_http_server.chuks &
    SERVER_PID=$!
    wait_for_server && run_wrk "Chuks VM"
    kill $SERVER_PID 2>/dev/null; wait $SERVER_PID 2>/dev/null || true

    echo -e "${BOLD}▶ Chuks AOT${NC}"
    kill_port
    chuks build bench_http_server.chuks > /dev/null 2>&1 || true
    if [ -x build/bench_http_server.chuks.bin ]; then
        xattr -cr build/bench_http_server.chuks.bin 2>/dev/null || true
        codesign --force --sign - build/bench_http_server.chuks.bin 2>/dev/null || true
        ./build/bench_http_server.chuks.bin &
        SERVER_PID=$!
        wait_for_server && run_wrk "Chuks AOT"
        kill $SERVER_PID 2>/dev/null; wait $SERVER_PID 2>/dev/null || true
    fi
fi

if have go; then
    echo -e "${BOLD}▶ Go (net/http)${NC}"
    kill_port
    go build -o go_stdlib_server go_server_stdlib.go 2>&1
    ./go_stdlib_server &
    SERVER_PID=$!
    wait_for_server && run_wrk "Go (net/http)"
    kill $SERVER_PID 2>/dev/null; wait $SERVER_PID 2>/dev/null || true

    echo -e "${BOLD}▶ Go (gnet+ants)${NC}"
    kill_port
    if go build -o bench_gnet_server raw_gnet.go 2>&1; then
        ./bench_gnet_server &
        SERVER_PID=$!
        wait_for_server && run_wrk "Go (gnet+ants)"
        kill $SERVER_PID 2>/dev/null; wait $SERVER_PID 2>/dev/null || true
    fi
fi

if have javac && have java; then
    echo -e "${BOLD}▶ Java (HttpServer)${NC}"
    kill_port
    javac JavaServer.java 2>&1
    java JavaServer &
    SERVER_PID=$!
    wait_for_server && run_wrk "Java"
    kill $SERVER_PID 2>/dev/null; wait $SERVER_PID 2>/dev/null || true
fi

if have bun; then
    echo -e "${BOLD}▶ Bun${NC}"
    kill_port
    bun bun_server.js &
    SERVER_PID=$!
    wait_for_server && run_wrk "Bun"
    kill $SERVER_PID 2>/dev/null; wait $SERVER_PID 2>/dev/null || true
fi

if have node; then
    echo -e "${BOLD}▶ Node.js${NC}"
    kill_port
    node node_server.js &
    SERVER_PID=$!
    wait_for_server && run_wrk "Node.js"
    kill $SERVER_PID 2>/dev/null; wait $SERVER_PID 2>/dev/null || true
fi

if have python3; then
    echo -e "${BOLD}▶ Python 3${NC}"
    kill_port
    python3 python_server.py &
    SERVER_PID=$!
    wait_for_server && run_wrk "Python 3"
    kill $SERVER_PID 2>/dev/null; wait $SERVER_PID 2>/dev/null || true
fi

echo ""
echo -e "${BOLD}╔══════════════════════════════════════════════════════════════════════════╗${NC}"
echo -e "${BOLD}║                         BENCHMARK RESULTS                                ║${NC}"
echo -e "${BOLD}║            wrk -t${WRK_THREADS} -c${WRK_CONNS} -d${WRK_DURATION}  •  GET / → \"Hello, World!\"                 ║${NC}"
echo -e "${BOLD}╠══════════════════════════════════════════════════════════════════════════╣${NC}"
printf "║  %-22s   %10s   %12s    %8s   ║\n" "Runtime" "Req/sec" "Avg Latency" "Transfer/s"
echo -e "${BOLD}╠══════════════════════════════════════════════════════════════════════════╣${NC}"
for i in "${!NAMES[@]}"; do
    printf "║  %-22s   %10s   %12s    %8s   ║\n" "${NAMES[$i]}" "${RPS_VALUES[$i]}" "${LATENCY_AVG[$i]}" "${TRANSFER[$i]}"
done
echo -e "${BOLD}╚══════════════════════════════════════════════════════════════════════════╝${NC}"
