FROM golang:1.13-alpine AS builder

RUN mkdir /build

ADD . /build/

WORKDIR /build

#go get -d -v ./... && go install -v ./... && \

# 通过设置环境变量实现跨平台编译，此处目标环境为 amd64 cpu 的 linux
RUN  go env -w GOPROXY=https://goproxy.cn,direct; \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo  -o main ./infra/gateway

# FROM scratch literally is an empty, zero-byte image / filesystem, where you add everything yourself.
FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]
