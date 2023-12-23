#!/bin/bash

VERSION="1.0.0"

PLATFORMS=(
	"darwin/amd64"
	"darwin/arm64"
	"linux/386"
	"linux/amd64"
	"linux/arm64"
)

mkdir -p build

for PLATFORM in "${PLATFORMS[@]}"; do
    GOOS=${PLATFORM%/*}
    GOARCH=${PLATFORM#*/}

    OUTPUT_NAME="build/cloudstart_${VERSION}_${GOOS}_${GOARCH}"

    echo "compile for $GOOS/$GOARCH..."
    env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "-X cloudstart/version.Version=$VERSION" -o $OUTPUT_NAME

    if [ $? -ne 0 ]; then
        echo "error during compilation for $GOOS $GOARCH"
        exit 1
    fi
done

echo "done"