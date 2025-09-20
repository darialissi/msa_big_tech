FROM golang:1.25-alpine AS builder

# Устанавливаем необходимые утилиты
RUN apk add --no-cache make git protoc protobuf-dev

WORKDIR /build

# Копируем только файлы, необходимые для установки зависимостей
COPY Makefile vendor.proto.mk ./       
COPY ./auth/go.mod ./auth/go.sum ./auth/

# Скачиваем зависимости (кешируемый слой)
RUN make vendor

# Копируем остальные файлы
COPY ./auth/ ./auth/

# Генерируем сервисы и собираем бинарник
RUN SERVICEDIR=auth make generate build

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/auth/bin/server ./

CMD ["./server"]