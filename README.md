### Вендоринг дополнительных модулей

```shell
make vendor
```

### Генерация и запуск отдельных сервисов

Подставляем в SERVICEDIR название желаемого сервиса

```shell
# кодогенерация для сервиса auth
SERVICEDIR=auth make generate
```

```shell
# билд для сервиса auth
SERVICEDIR=auth make build
```

Бинарники сохраняются в соответствующую директорию сервиса


```shell
# запуск сервера для сервиса auth
./auth/bin/server
```

```shell
# запуск клиента для сервиса auth
./auth/bin/client
```