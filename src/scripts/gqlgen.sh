#!/bin/bash
printf "\nRegenerating gqlgen files\n"
rm -f graph/generated/generated.go \
    graph/model/generated.go \
    graph/resolvers/generated.go
time go run -v github.com/99designs/gqlgen $1
printf "\nDone.\n\n"
