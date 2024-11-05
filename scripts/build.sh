#! /usr/bin/env bash 


mod=$(ls|grep go.mod)
if [ "$mod" != "go.mod" ] ; then 
    go mod init Golang-bc8-quera/web_crawler
    go mod tidy
fi 


go build -o web-crawler ../cmd/web-crawler/main.go


