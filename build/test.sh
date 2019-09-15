#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

# Enable C code, as it is needed for SQLite3 database binary
# Enable go modules
export CGO_ENABLED=1
export GO111MODULE=on
export GOFLAGS="-mod=vendor"

# Collect test targets
TARGETS=$(for d in "$@"; do echo ./$d/...; done)

# Run tests
echo "Running tests:"
go test -installsuffix "static" ${TARGETS} 2>&1
echo

# Collect all `.go` files and `gofmt` against them. If some need formatting - print them.
echo -n "Checking gofmt: "
ERRS=$(find "$@" -type f -name \*.go | xargs gofmt -l 2>&1 || true)
if [ -n "${ERRS}" ]; then
    echo "FAIL - the following files need to be gofmt'ed:"
    for e in ${ERRS}; do
        echo "    $e"
    done
    echo
    exit 1
fi
echo "PASS"
echo

# Run `go vet` against all targets. If problems are found - print them.
echo -n "Checking go vet: "
ERRS=$(go vet ${TARGETS} 2>&1 || true)
if [ -n "${ERRS}" ]; then
    echo "FAIL"
    echo "${ERRS}"
    echo
    exit 1
fi
echo "PASS"
echo
