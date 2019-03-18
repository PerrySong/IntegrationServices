#!/usr/bin/env bash
export PATH=$PATH:$GOPATH/bin
protoc user.proto --go_out=plugins=grpc:.