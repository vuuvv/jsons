#!/bin/sh
set -ex
go get -u github.com/bytedance/sonic
go mod tidy
git add -A .
git commit -m "update dependencies version"
git push
git tag $1
git push origin tag $1