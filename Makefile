include vendor.proto.mk

# Используем переменную окружения или значение по умолчанию
SUBDIR ?= $(SERVICEDIR)

CURDIR := $(CURDIR)/$(SUBDIR)

# Используем bin в текущей директории для установки плагинов protoc
LOCAL_BIN := $(CURDIR)/bin

# Добавляем bin в текущей директории в PATH при запуске protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

# Путь до protobuf файлов
PROTO_PATH := $(CURDIR)/api/proto

# Путь до сгенеренных .pb.go файлов
PKG_PROTO_PATH := $(CURDIR)/pkg

# устанавливаем необходимые плагины
.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	$(info Installing binary dependencies...)

	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
# go install github.com/bufbuild/buf/cmd/buf@latest

# генерация .go файлов с помощью protoc
.protoc-generate:
	$(PROTOC) -I $(VENDOR_PROTO_PATH) --proto_path=$(CURDIR) \
	--go_out=$(PKG_PROTO_PATH) --go_opt paths=source_relative \
	--go-grpc_out=$(PKG_PROTO_PATH) --go-grpc_opt paths=source_relative \
	--grpc-gateway_out=$(PKG_PROTO_PATH) --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true \
	$(PROTO_PATH)/service.proto \
	$(PROTO_PATH)/messages.proto

	$(PROTOC) -I $(VENDOR_PROTO_PATH) --proto_path=$(CURDIR) \
	--openapiv2_out=$(CURDIR) --openapiv2_opt logtostderr=true \
	$(PROTO_PATH)/service.proto

# 	$(PROTOC) -I $(VENDOR_PROTO_PATH) --proto_path=$(CURDIR) \
# 	--validate_out="lang=go,paths=source_relative:$(PKG_PROTO_PATH)" \
# 	$(PROTO_PATH)/messages.proto

# go mod tidy
.tidy:
	cd $(CURDIR) && GOBIN=$(LOCAL_BIN) go mod tidy

# Генерация кода из protobuf
generate: .bin-deps .protoc-generate .tidy

# Билд приложения
build:
	cd $(CURDIR) && go build -o $(LOCAL_BIN) $(CURDIR)/cmd/client
	cd $(CURDIR) && go build -o $(LOCAL_BIN) $(CURDIR)/cmd/server 
	
# Объявляем, что текущие команды не являются файлами и
# интсрументируем Makefile не искать изменения в файловой системе
.PHONY: \
	.bin-deps \
	.protoc-generate \
	.vendor-protovalidate \
	.tidy \
	vendor \
	generate \
	build