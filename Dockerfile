FROM golang:alpine

RUN apk update && apk add --no-cache make protobuf-dev

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
RUN wget https://github.com/grpc/grpc-web/releases/download/1.4.2/protoc-gen-grpc-web-1.4.2-linux-x86_64
RUN mv protoc-gen-grpc-web-1.4.2-linux-x86_64 /usr/local/bin/protoc-gen-go-grpc
RUN chmod 777 /usr/local/bin/protoc-gen-go-grpc

# Create app directory
WORKDIR /app/go-grpc-api-gateway-main

ENTRYPOINT ["make", "server"]
