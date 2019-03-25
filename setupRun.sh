#!/usr/bin/env bash
dep ensure
go run DBMigrate.go
go run main.go &
go run servers/server.go &