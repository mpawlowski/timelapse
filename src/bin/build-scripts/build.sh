#!/bin/bash

set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT_DIR=$DIR/../../..

go fmt $ROOT_DIR/...

rm -rf $ROOT_DIR/build/
mkdir -p $ROOT_DIR/build/

GIT_HASH=$(git rev-parse HEAD)
GIT_DIRTY=$(git diff --quiet; [ $? -eq 0 ] && echo false || echo true)

function build() {
  local os=$1
  local arch=$2
  local name=$3
  CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build -o build/$name -ldflags "-X main.gitHash=$GIT_HASH -X main.gitDirty=$GIT_DIRTY" $ROOT_DIR/src/bin/cli/cli.go $ROOT_DIR/src/bin/cli/build.go $ROOT_DIR/src/bin/cli/help.go
}

build linux amd64 timelapse-linux-amd64