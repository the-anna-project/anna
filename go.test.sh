#!/usr/bin/env bash

set -e
> coverage.txt

for d in $(find ./* -maxdepth 10 -type f -name '*test.go' -not -path "./.workspace/*" -not -path "./.git/*" -exec dirname {} \; | sort -u); do
  go test -covermode=count -coverprofile=profile.out $d
  if [ -f profile.out ]; then
      cat profile.out >> coverage.txt
      rm profile.out
  fi
done
