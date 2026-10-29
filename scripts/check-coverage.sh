#!/bin/bash

THRESHOLD=80

go test -coverprofile=coverage.out ./... > /dev/null 2>&1


# coverage=$(go test -cover ./... | grep -o '[0-9.]*%' | tr -d '%' | awk '{s+=$1; n++} END {printf "%.2f", s/n}')
coverage=$(go tool cover -func=coverage.out | grep total: | grep -o '[0-9.]*%' | tr -d '%')
rm ./coverage.out
echo "Coverage Total: $coverage%"
if (( $(echo "$coverage < $THRESHOLD" | bc -l) )); then
  echo "Current test coverage is below $THRESHOLD%"
  exit 1
fi

echo "Current coverage is accepted."
