#!/usr/bin/env bash
set -eo pipefail

echo "Mocking interfaces in 'src'..."
mockery --dir ./src --all --keeptree --output mocks
