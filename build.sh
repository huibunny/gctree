#!/bin/bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gctree cmd/app/main.go
