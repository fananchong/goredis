#!/bin/bash

set -ex


# update
./legacy.sh

cp -f ./Godeps.json.template ./Godeps.json
cd ..
rm -rf ./vendor
docker run --rm -e GOPATH=/go/:/temp/ -v /temp/:/temp/ -v "$PWD":/go/src/github.com/fananchong/goredis -w /go/src/github.com/fananchong/goredis fananchong/godep save ./...

