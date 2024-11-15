#!/bin/bash

export PROJECT_ROOT=$(pwd)

export GO_ENV=test

test_paths=(
    "./modules/user/Tests"
    "./modules/auth/Tests"
)

if [ -f "$PROJECT_ROOT/.env.test" ]; then
    export $(grep -v '^#' "$PROJECT_ROOT/.env.test" | xargs)
else
    echo ".env.test file not found in project root!"
    exit 1
fi

for path in "${test_paths[@]}"
do
  echo "Running tests in $path"
  CGO_ENABLED=1 go test -v "$path"
done