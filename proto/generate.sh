#!/usr/bin/env bash
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
protoc user.proto --go_out=plugins=grpc:.