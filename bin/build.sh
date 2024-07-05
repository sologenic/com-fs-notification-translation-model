#!/bin/bash
set -e

# move to the root dir of the package
rd=$(git rev-parse --show-toplevel)
cd $rd

# consider a TS project exists in the package, when both package.json and tsconfig.json
if [[ -f ./package.json && -f ./tsconfig.json ]]; then has_ts=1; else has_ts=0; fi

# count the number of ./*.proto files and trim white space
has_proto=$(ls -1 *.proto|grep -v grpc 2>/dev/null | wc -l | xargs)

# get the service name by substring
service=$(ls *.proto| grep -v grpc | cut -d'/' -f2 | cut -d'.' -f1)

protoc \
--proto_path=. "translation.proto" \
--proto_path=$(dirname $(dirname "$rd")) \
"--go_out=." --go_opt=paths=source_relative \
--go-grpc_opt=require_unimplemented_servers=false \
"--go-grpc_out=." --go-grpc_opt=paths=source_relative