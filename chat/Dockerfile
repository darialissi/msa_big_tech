FROM golang:1.25-alpine AS builder

# Устанавливаем необходимые утилиты
RUN apk add --no-cache make git protoc protobuf-dev

WORKDIR /build

# Копируем только файлы, необходимые для установки зависимостей
COPY Makefile vendor.proto.mk ./       
COPY ./chat/go.mod ./chat/go.sum ./chat/

# Скачиваем зависимости (кешируемый слой)
RUN make vendor

# Копируем остальные файлы
COPY ./chat/ ./chat/

# Генерируем сервисы и собираем бинарник
RUN SERVICEDIR=chat make generate build

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/chat/bin/server ./

CMD ["./server"]