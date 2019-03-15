#!/usr/bin/env bash
export PATH=$PATH:$GOPATH/bin
protoc proto/user.proto --go_out=plugins=grpc:.