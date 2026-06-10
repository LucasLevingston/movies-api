FROM golang:1.22-alpine AS proto-gen
RUN apk add --no-cache protobuf && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

WORKDIR /workspace
COPY proto/ proto/
RUN mkdir -p gen/movies && \
    protoc \
      --go_out=. \
      --go_opt=module=movies-api/movies-service \
      --go-grpc_out=. \
      --go-grpc_opt=module=movies-api/movies-service \
      -I proto \
      proto/movies.proto

FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY movies-service/ .
COPY --from=proto-gen /workspace/gen ./gen
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .
COPY movies.json .
EXPOSE 50051
CMD ["./server"]
