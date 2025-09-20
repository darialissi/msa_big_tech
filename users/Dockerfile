FROM golang:1.25-alpine AS builder

# Устанавливаем необходимые утилиты
RUN apk add --no-cache make git protoc protobuf-dev

WORKDIR /build

# Копируем только файлы, необходимые для установки зависимостей
COPY Makefile vendor.proto.mk ./       
COPY ./users/go.mod ./users/go.sum ./users/

# Скачиваем зависимости (кешируемый слой)
RUN make vendor

# Копируем остальные файлы
COPY ./users/ ./users/

# Генерируем сервисы и собираем бинарник
RUN SERVICEDIR=users make generate build

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/users/bin/server ./

CMD ["./server"]