#!/bin/bash
echo "first run"
go run main.go $@
ls -al /tmp/hello
echo "second run"
go run main.go $@
ls -al /tmp/hello
