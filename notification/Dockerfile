FROM golang:1.25-alpine AS builder

# Configure Go proxy with direct fallback for better reliability
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=sum.golang.org

WORKDIR /build

# Кэширование зависимостей
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/ ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/bin/server ./

CMD ["./server"]