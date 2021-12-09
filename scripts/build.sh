#!/usr/bin/env bash
set -eo pipefail

[ ! -d "./mocks" ] && echo "Mocks DO NOT exist." && exit 1

echo "Install vendor modules..."
go mod vendor

echo "Build device-commander binary..."
go build -o device-commander
