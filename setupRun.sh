#!/usr/bin/env bash
dep ensure
docker-compose up
go run DBMigrate.go
go run main.go &
go run servers/server.go &