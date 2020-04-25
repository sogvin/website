#!/bin/bash
path=$1
dir=$(dirname "$path")
filename=$(basename "$path")
extension="${filename##*.}"
nameonly="${filename%.*}"

case $extension in
    go)
        goimports -w $path
	go test -coverprofile /tmp/c.out ./...
	uncover /tmp/c.out
        ;;
esac
reloadsurf.sh