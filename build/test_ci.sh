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

# Run tests and generate reports in JSON format, then copy them to path/file expected CodeClimate and SonarCloud
echo "Running tests and Generating reports..."
go test -coverprofile=/reports/coverage.out -installsuffix "static" ${TARGETS} -json > /reports/test-report.out
cp /reports/coverage.out /coverage/c.out
echo

# Check coverage using test report from previous step
echo "Coverage:"
go tool cover -func=/coverage/c.out
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
ERRS=$(go vet ${TARGETS} 2>&1 | tee "/reports/vet.out" || true)
if [ -n "${ERRS}" ]; then
    echo "FAIL"
    echo "${ERRS}"
    echo
    exit 1
fi
echo "PASS"
echo
