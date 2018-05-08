#!/usr/bin/env
set -e
version=$(git describe --tags --always)
go build -ldflags "-X markdown/version.VERSION=$version" -o output/markdown
output/markdown
