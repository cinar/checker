#!/bin/bash

COVERAGE_FILE=coverage.out

if [ -z "$CODE_COVERAGE_THRESHOLD" ]
then
  echo "No code coverage threshold is set."
  exit 0
fi

echo "Code coverage threshold ${CODE_COVERAGE_THRESHOLD}%"

if [ ! -e "$COVERAGE_FILE" ]
then
  echo "Coverage file ${COVERAGE_FILE} is not found."
  exit 1
fi

CODE_COVERAGE=$(go tool cover -func=coverage.out | grep "total:" | grep -E -o '[0-9]+\.[0-9]+')

if awk "BEGIN { exit !($CODE_COVERAGE < $CODE_COVERAGE_THRESHOLD)}"
then
  echo "Current code coverage ${CODE_COVERAGE}% is below the ${CODE_COVERAGE_THRESHOLD}% threshold."
  exit 1
fi
