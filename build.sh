#!/bin/bash

#!/bin/bash
set -e

GOOS_ARCHES=(
  "linux amd64"
  "darwin amd64"
  "darwin arm64"
)

if [[ $# -ne 1 ]]; then
  echo "Usage: $0 [version, e.g. v0.1.0]"
  exit 1
fi

VERSION=$1

mkdir -p dist/$1


for combo in "${GOOS_ARCHES[@]}"; do
  read -r GOOS GOARCH <<< "$combo"
  out="dok-${GOOS}-${GOARCH}"
  [[ "$GOOS" == "windows" ]] && out="${out}.exe"
  echo "Building $out..."
  GOOS=$GOOS GOARCH=$GOARCH go build -o "dist/$1/$out" main.go
done

ghr $1 dist/$1
echo "Release $VERSION created successfully."
