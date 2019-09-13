FROM golang:1.12-alpine

RUN apk add gcc g++

# Docker image for running tests. This image is needed because tests use SQLite3 as in-memory database
# and that requires CGO to be enabled, which in turn requires GCC and G++ to be installed.
