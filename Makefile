include scratch.mk
include vendor.proto.mk

# Локальная накатка миграций
migrate: .bin-deps-migrate .migrate-up
fast-migrate: .migrate-up

# Поднятие всех сервисов
up:
	docker compose up

# Остановка всех сервисов
down:
	docker compose down
    
# Билд отдельных сервисов
server:
	cd $(CURDIR) && go build -o $(LOCAL_BIN) $(CURDIR)/cmd/server 

client:
	cd $(CURDIR) && go build -o $(LOCAL_BIN) $(CURDIR)/cmd/client

# Тестирование usecase отдельного сервиса
test:
	cd $(CURDIR) && go test -v ./internal/app/usecases


# Объявляем, что текущие команды не являются файлами и
# инструментируем Makefile не искать изменения в файловой системе
.PHONY: \
	migrate \
	up \
	down \
	server \
	client \
	test