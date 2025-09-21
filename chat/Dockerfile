FROM golang:1.25-alpine AS builder

WORKDIR /build

COPY ./chat/ ./

RUN go build -o ./bin/ ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/bin/server ./

CMD ["./server"]