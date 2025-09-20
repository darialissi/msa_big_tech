FROM golang:1.25-alpine AS builder

# Устанавливаем необходимые утилиты
RUN apk add --no-cache make git protoc protobuf-dev

WORKDIR /build

# Копируем только файлы, необходимые для установки зависимостей
COPY Makefile vendor.proto.mk ./       
COPY ./social/go.mod ./social/go.sum ./social/

# Скачиваем зависимости (кешируемый слой)
RUN make vendor

# Копируем остальные файлы
COPY ./social/ ./social/

# Генерируем сервисы и собираем бинарник
RUN SERVICEDIR=social make generate build

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/social/bin/server ./

CMD ["./server"]