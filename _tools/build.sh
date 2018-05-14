#!/bin/sh
go build -o ./bin/esso ./cmd/esso/main.go
go build -o ./bin/amex ./cmd/amex/main.go
# TODO makefile