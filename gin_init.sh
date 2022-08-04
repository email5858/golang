#!/bin/bash

go mod init gin

go mod edit -require github.com/gin-gonic/gin@latest

go get -u -v github.com/gin-gonic/gin && go run main.go