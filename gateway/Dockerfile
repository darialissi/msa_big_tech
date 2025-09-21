FROM golang:1.25-alpine AS builder

WORKDIR /build

COPY ./gateway/ ./gateway/

# Зависимые сервисы
COPY ./auth/go.mod ./auth/
COPY ./auth/pkg/ ./auth/pkg/
COPY ./chat/go.mod ./chat/
COPY ./chat/pkg/ ./chat/pkg/

RUN cd gateway && go build -o ./bin/ ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/gateway/bin/server ./

CMD ["./server"]