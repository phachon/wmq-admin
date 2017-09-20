#!/bin/sh

tarfile="wmq-admin-$1.tar.gz"

echo "开始打包$tarfile..."

export GOARCH=amd64
export GOOS=linux

go build ./

bee pack

mv wmq-admin.tar.gz $tarfile
