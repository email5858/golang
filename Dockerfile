FROM golang:alpine
# FROM golang:latest

ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://proxy.golang.com.cn,direct"

WORKDIR $GOPATH/src

WORKDIR $GOPATH/src/myblog

COPY ./myblog $GOPATH/src/myblog

EXPOSE 8080

# RUN sh "/go/src/myblog/gin_init.sh"
# RUN chmod +x $GOPATH/src/myblog

ENTRYPOINT sh gin_init.sh