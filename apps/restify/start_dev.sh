#!/bin/bash

echo "🚀 Starting dev server..."

if ! command -v air &> /dev/null; then
    echo "❌ Air is not installed. Run: go install github.com/cosmtrek/air@latest"
    exit 1
fi

air -c .air.toml

