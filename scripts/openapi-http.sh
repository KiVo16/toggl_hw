#!/bin/bash
set -e

readonly api="$1"
readonly output_dir="$2"
readonly package="$3"

oapi-codegen -old-config-style -generate server -o "$output_dir/openapi_types.gen.go" -package "$package" "./api/openapi/$api.yml"
