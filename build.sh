#!/bin/bash

#!/bin/bash
set -e

GOOS_ARCHES=(
  "linux amd64"
  "darwin amd64"
  "darwin arm64"
)

for combo in "${GOOS_ARCHES[@]}"; do
  read -r GOOS GOARCH <<< "$combo"
  out="dok-${GOOS}-${GOARCH}"
  [[ "$GOOS" == "windows" ]] && out="${out}.exe"
  echo "Building $out..."
  GOOS=$GOOS GOARCH=$GOARCH go build -o "dist/$out" main.go
done
