#!/bin/bash

echo "🐞 Starting debugger..."

if ! command -v dlv &> /dev/null; then
    echo "❌ Delve (dlv) is not installed. Run: go install github.com/go-delve/delve/cmd/dlv@latest"
    exit 1
fi

DLV_PORT=40000

dlv debug ./cmd/main.go --headless --listen=:${DLV_PORT} --api-version=2 --accept-multiclient --log
