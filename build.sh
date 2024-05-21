#!/bin/bash

set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

go fmt $DIR/...

rm -rf $DIR/build/
mkdir -p $DIR/build/

GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
GIT_HASH=$(git rev-parse HEAD)
GIT_DIRTY=$(git diff --quiet; [ $? -eq 0 ] && echo false || echo true)

function build() {
  local os=$1
  local arch=$2
  local name=$3
  CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build -o build/$name -ldflags "-X main.gitBranch=$GIT_BRANCH -X main.gitHash=$GIT_HASH -X main.gitDirty=$GIT_DIRTY" $DIR/src/bin/cli/cli.go $DIR/src/bin/cli/build.go $DIR/src/bin/cli/help.go
}

build linux amd64 timelapse-linux-amd64
build windows amd64 timelapse-windows-amd64.exe