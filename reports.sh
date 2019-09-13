#!/usr/bin/env bash
# Prepares files and directories needed when generating reports (mainly used in CI/CD build)

mkdir -p reports
touch reports/coverage.out reports/test-report.out reports/vet.out
touch c.out
