FROM golang:1.27.2-alpine AS builder
WORKDIR /app
ENV CGO_ENABLED=0
COPY ./core/go.mod ./core/go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY ./core .
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
	go build -ldflags="-s -w" -o ./bin/core ./cmd/core

FROM alpine:latest
WORKDIR /app
EXPOSE 9000
COPY --from=builder /app/bin/core ./bin/core
CMD ["./bin/core"]
