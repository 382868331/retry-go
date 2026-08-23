#!/usr/bin/env bash
set -euo pipefail
docker build --platform linux/amd64 -f goletalab.Dockerfile -t "$1" .
docker run --rm "$1" go test ./...
