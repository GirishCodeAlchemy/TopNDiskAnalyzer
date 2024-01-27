#!/bin/bash

platforms=("linux/amd64" "linux/386" "darwin/amd64" "windows/amd64" "windows/386")

# Create a package folder if it doesn't exist
mkdir -p package

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name="tdf-${GOOS}-${GOARCH}"

    if [ $GOOS = "windows" ]; then
        output_name="tdf.exe"
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o package/$output_name main.go
done
