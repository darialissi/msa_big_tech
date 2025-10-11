### Билд отдельных сервисов

Подставляем в SERVICEDIR название желаемого сервиса.

```shell
SERVICEDIR=auth make server
```

Запуск сервера

```shell
./auth/bin/server
```

### Поднятие сервисов в одной сети 

```shell
make up
```

### Примеры запросов

#### AuthService

```shell
grpcurl -plaintext localhost:8083 list
```

```shell
grpcurl -plaintext -d '{"email": "test@example.com", "password": "paSS123"}' localhost:8083 github.com.darialissi.msa_big_tech.auth.AuthService/Register
```


### Локальное тестирование

Пока написаны только unit тесты на *usecases* сервиса *social*