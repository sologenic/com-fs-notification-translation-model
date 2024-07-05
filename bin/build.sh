#!/bin/bash
set -e

# move to the root dir of the package
rd=$(git rev-parse --show-toplevel)
cd $rd

protoc \
  --proto_path=${HOME}/go/src/github.com/sologenic \
  --proto_path=. \
  --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  translation.proto