#!/bin/bash

# Compile stm for all targets

# Clean bin directory
rm -rf bin
mkdir bin

# Linux 386
GOARCH=386 GOOS=linux go build
mv stm bin/stm-linux-386

# Linux amd64
GOARCH=amd64 GOOS=linux go build
mv stm bin/stm-linux-amd64

# Linux arm
GOARCH=arm GOOS=linux go build
mv stm bin/stm-linux-arm

# Mac amd64
GOARCH=amd64 GOOS=darwin go build
mv stm bin/stm-darwin-amd64

echo Done compiling.
