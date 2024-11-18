#!/bin/bash

export PROJECT_ROOT=$(pwd)
export GO_ENV=test

# validation
if [ -z "$1" ]; then
  echo "Please provide the path to the test file (e.g., ./modules/user/Tests/my_test.go)"
  exit 1
fi

test_path="$1"

# load env
if [ -f "$PROJECT_ROOT/.env.test" ]; then
    export $(grep -v '^#' "$PROJECT_ROOT/.env.test" | xargs)
else
    echo ".env.test file not found in project root!"
    exit 1
fi

# run
echo "Running test in $test_path"
CGO_ENABLED=1 go test -v "$test_path"
