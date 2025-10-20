## Мессенджер (аналог Discord)

[Полное описание приложения](./description.md)

### Поднятие сервисов в одной сети 

```shell
make up
```

### Накатка миграций отдельных сервисов 

Подставляем в SERVICEDIR название желаемого сервиса.

```shell
SERVICEDIR=auth make migrate
```

```shell
# без установки зависимостей
SERVICEDIR=auth make fast-migrate
```

### Примеры запросов

Запросы можно прогнать через скрипт клиента.

Подставляем в SERVICEDIR название желаемого сервиса.

```shell
# билд
SERVICEDIR=chat make client
```

```shell
# запуск
./chat/bin/client
```

#### AuthService

```shell
grpcurl -plaintext localhost:8083 list github.com.darialissi.msa_big_tech.auth.AuthService
```

```shell
grpcurl -plaintext -d '{"email": "test@example.com", "password": "paSS123"}' localhost:8083 github.com.darialissi.msa_big_tech.auth.AuthService/Register
```

#### ChatService

```shell
grpcurl -plaintext localhost:8084 list github.com.darialissi.msa_big_tech.chat.ChatService
```

```shell
grpcurl -plaintext -d '{"participant_id": "ba209999-0c6c-11d2-97cf-00c04f8eea45"}' localhost:8084 github.com.darialissi.msa_big_tech.chat.ChatService.CreateDirectChat
```

```shell
grpcurl -plaintext -d '{"user_id": "ba209999-0c6c-11d2-97cf-00c04f8eea45"}' localhost:8084 github.com.darialissi.msa_big_tech.chat.ChatService.ListUserChats
```

```shell
grpcurl -plaintext -d '{"chat_id": "da4a578b-e952-4e16-a335-22a46664a5f9"}' localhost:8084 github.com.darialissi.msa_big_tech.chat.ChatService.GetChat
```

```shell
grpcurl -plaintext -d '{"chat_id": "da4a578b-e952-4e16-a335-22a46664a5f9"}' localhost:8084 github.com.darialissi.msa_big_tech.chat.ChatService.ListChatMembers
```

```shell
grpcurl -plaintext -d '{"chat_id": "e00ea407-70b4-4f90-9777-62638519b0bf", "text": "hi there!"}' localhost:8084 github.com.darialissi.msa_big_tech.chat.ChatService.SendMessage
```

```shell
grpcurl -plaintext -d '{"chat_id": "e00ea407-70b4-4f90-9777-62638519b0bf", "limit": "5"}' localhost:8084 github.com.darialissi.msa_big_tech.chat.ChatService.ListMessages
```

```shell
grpcurl -plaintext -d '{"chat_id": "e00ea407-70b4-4f90-9777-62638519b0bf", "since_unix_ms": "1760210300"}' localhost:8084 github.com.darialissi.msa_big_tech.chat.ChatService.StreamMessages
```

#### SocialService

```shell
grpcurl -plaintext localhost:8085 list github.com.darialissi.msa_big_tech.social.SocialService
```

```shell
grpcurl -plaintext -d '{"user_id": "ba209999-0c6c-11d2-97cf-00c04f8eea45"}' localhost:8085 github.com.darialissi.msa_big_tech.social.SocialService.SendFriendRequest
```

```shell
grpcurl -plaintext -d '{"user_id": "ba209999-0c6c-11d2-97cf-00c04f8eea45"}' localhost:8085 github.com.darialissi.msa_big_tech.social.SocialService.ListFriendRequests
```

```shell
grpcurl -plaintext -d '{"friend_request_id": "cb8bc828-353e-4d7a-a3d8-f3145bfd1b67"}' localhost:8085 github.com.darialissi.msa_big_tech.social.SocialService.AcceptFriendRequest
```

```shell
grpcurl -plaintext -d '{"friend_request_id": "c7da8bee-3a8d-4a43-907d-0d4f7655cef9"}' localhost:8085 github.com.darialissi.msa_big_tech.social.SocialService.DeclineFriendRequest
```

```shell
grpcurl -plaintext -d '{"user_id": "ba209999-0c6c-11d2-97cf-00c04f8eea45", "cursor": "2025-10-19T13:54:27+03:00"}' localhost:8085 github.com.darialissi.msa_big_tech.social.SocialService.ListFriends
```

#### UsersService

```shell
grpcurl -plaintext localhost:8086 list github.com.darialissi.msa_big_tech.users.UsersService
```

```shell
grpcurl -plaintext -d '{"nickname": "helloworld", "bio": "about me"}' localhost:8086 github.com.darialissi.msa_big_tech.users.UsersService.CreateProfile
```

```shell
grpcurl -plaintext -d '{"user_id": "dd3d39f1-7039-4524-a8ce-0b0555929b1f", "nickname": "hello-world", "avatar_url": "http://avatar.me"}' localhost:8086 github.com.darialissi.msa_big_tech.users.UsersService.UpdateProfile
```

```shell
grpcurl -plaintext -d '{"user_id": "dd3d39f1-7039-4524-a8ce-0b0555929b1f"}' localhost:8086 github.com.darialissi.msa_big_tech.users.UsersService.GetProfileByID
```

```shell
grpcurl -plaintext -d '{"nickname": "hello-world"}' localhost:8086 github.com.darialissi.msa_big_tech.users.UsersService.GetProfileByNickname
```

#### API Gateway

```shell
curl -I http://localhost:8087/
```

```shell
curl -X POST http://localhost:8087/api/v1/register \
  -H "Content-Type: application/json" \
  -d '{"email": "test2@example.com", "password": "paSS123"}'
```

```shell
curl -X POST http://localhost:8087/api/v1/friends/requests \
  -H "Content-Type: application/json" \
  -d '{"user_id": "ba209999-0c6c-11d2-97cf-00c04f8eea45"}'
```

```shell
curl -X POST http://localhost:8087/api/v1/chats/5b46d785-ac49-480b-b1b3-bdd697267487/message \
  -H "Content-Type: application/json" \
  -d '{"text": "hi there!"}'
```

```shell
curl -X GET http://localhost:8087/api/v1/profile/ba209999-0c6c-11d2-97cf-00c04f8eea45
```

### Локальное тестирование

Пока написаны только unit тесты на *usecases* сервиса *social*

```shell
SERVICEDIR=social make test
```