### Вендоринг дополнительных модулей

```shell
make vendor
```

### Генерация и билд отдельных сервисов

Подставляем в SERVICEDIR название желаемого сервиса

```shell
# пример для сервиса auth
SERVICEDIR=auth make generate build
```

Бинарники сохраняются в соответствующую директорию сервиса

```shell
# запуск сервера
./auth/bin/server
```

```shell
# запуск клиента
./auth/bin/client
```

### Поднятие сервисов в одной сети 

```shell
docker compose up
```