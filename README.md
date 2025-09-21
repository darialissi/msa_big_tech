### Вендоринг + кодогенерация отдельных сервисов

Подставляем в SERVICEDIR название желаемого сервиса.

```shell
# пример для сервиса auth
SERVICEDIR=auth make vendor generate
```

### Билд отдельных сервисов

Подставляем в SERVICEDIR название желаемого сервиса.

```shell
# пример для сервиса auth
SERVICEDIR=auth make build
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