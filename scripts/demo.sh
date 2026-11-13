#!/usr/bin/env bash
# End-to-end check of the SDP lab started with `docker compose up -d --build`.
#
#   1. The client cannot reach the protected service (gateway drops everything).
#   2. The client cannot reach the private network directly.
#   3. The client sends an SPA packet (fwknop -n service_gate).
#   4. The gateway opens access for that client only, and the request succeeds.
#
# Exits non-zero if any expectation fails, so it can be used in CI.
set -uo pipefail

GATEWAY_IP="${GATEWAY_IP:-10.231.0.20}"
PROTECTED_SERVICE_IP="${PROTECTED_SERVICE_IP:-10.231.1.30}"
PROTECTED_PORT="${PROTECTED_PORT:-8080}"
URL="http://${GATEWAY_IP}:${PROTECTED_PORT}/"

failures=0
pass() { printf '  \033[32mPASS\033[0m %s\n' "$1"; }
fail() { printf '  \033[31mFAIL\033[0m %s\n' "$1"; failures=$((failures + 1)); }
client() { docker compose exec -T sdp-client "$@"; }
http_code() { client curl -s -m 4 -o /dev/null -w '%{http_code}' "$1" 2>/dev/null; }

echo "[1/4] Request to $URL before SPA (expect: blocked)"
code="$(http_code "$URL")"
if [ "$code" = "000" ]; then pass "no response, the service is hidden"; else fail "got HTTP $code"; fi

echo "[2/4] Direct request to the private network (expect: unreachable)"
code="$(http_code "http://${PROTECTED_SERVICE_IP}/")"
if [ "$code" = "000" ]; then pass "private network not reachable from the client"; else fail "got HTTP $code"; fi

echo "[3/4] Sending SPA packet: fwknop -n service_gate"
if client fwknop -n service_gate; then pass "SPA sent and controller session completed"; else fail "fwknop exited with an error"; fi
sleep 1

echo "[4/4] Request to $URL after SPA (expect: HTTP 200)"
code="$(http_code "$URL")"
if [ "$code" = "200" ]; then pass "protected service reachable through the gateway"; else fail "got HTTP $code"; fi

echo
if [ "$failures" -eq 0 ]; then
    echo "All checks passed."
else
    echo "$failures check(s) failed. Inspect: docker compose logs sdp-gateway sdp-controller"
    exit 1
fi
