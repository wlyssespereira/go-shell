#/bin/bash

export GOPATH=`pwd`
go build -o ./bin/gosh  ./src/main/main.go
