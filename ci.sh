#!/bin/bash

set -e
set -o pipefail

dist=/tmp/website

case $1 in
    g|generate)
	echo "generate"	
	pushd ../
	tree -P "*.go" -I "*_test.go" navstar | \
	    grep -v directories > website/example/navstar.tree
	popd
	;;
    b|build)
	echo "build"	
	go build ./...
	
	mkdir -p $dist
	go run ./cmd/mksite -p $dist/docs
	# update static files
	rsync -aC ./docs $dist/
	echo "dist: $dist"
	;;
    t|test)
	echo "test"	
	go test -coverprofile /tmp/c.out ./... 2>&1 | \
	    sed -e 's| of statements||g' \
		-e 's|coverage: ||g' \
		-e 's|github.com/gregoryv/website|.|g' | \
	    grep -v "no test"
	;;
    publish)
	echo "publish"
	go run ./cmd/mksite -c # guard
	rsync -avC $dist/docs/ www.7de.se:/var/www/www.sogvin.com/
	;;
    clean)
	echo "clean"	
	rm -rf $dist
	;;
    -h)
	echo "Usage: $0 build|publish"
	;;
    *)
	$0 build test
	;;
esac


# Run next target if any
shift
[[ -z "$@" ]] && exit 0
$0 $@

