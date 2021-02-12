#!/bin/sh

protoc -I . flow.proto --go_out=plugins=grpc:.