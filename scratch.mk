# Используем переменную окружения или значение по умолчанию
SUBDIR ?= $(SERVICEDIR)

CURDIR := $(CURDIR)/$(SUBDIR)

# Используем bin в текущей директории для установки плагинов
LOCAL_BIN := $(CURDIR)/bin/

# Путь до сгенеренных .pb.go файлов
PKG_PROTO_PATH := $(CURDIR)/pkg

# Установка плагинов для proto генерации
.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	$(info Installing binary dependencies...)

	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/bufbuild/buf/cmd/buf@latest

.buf-generate:
	cd $(CURDIR) && $(LOCAL_BIN)/buf generate

# go mod tidy
.tidy:
	cd $(CURDIR) && GOBIN=$(LOCAL_BIN) go mod tidy

# Генерация кода из proto
generate: .bin-deps .buf-generate .tidy
fast-generate: .buf-generate .tidy

# Установка плагинов для миграций БД
.bin-deps-migrate: export GOBIN := $(LOCAL_BIN)
.bin-deps-migrate:
	$(info Installing binary dependencies...)
	go install github.com/pressly/goose/v3/cmd/goose@latest

.migrate-up:
	cd $(CURDIR) && goose -dir migrations up

.migrate-down:
	cd $(CURDIR) && goose -dir migrations down


# Объявляем, что текущие команды не являются файлами и
# инструментируем Makefile не искать изменения в файловой системе
.PHONY: \
	.bin-deps \
	.protoc-generate \
	.tidy \
	generate