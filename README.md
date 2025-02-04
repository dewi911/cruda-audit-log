# cruda-audit-log
## Оглавление
- [Генерация](#генерация)
- [Использование](#использование)
- [Конфигурация](#конфигурация)

## Генерация

```bash
protoc --go_out=. --go-grpc_out=. proto/audit.proto
```
## Использование

Сборка
```bash
source .env && go build -o app cmd/main.go && ./app
```

Запуск в докере монго

```bash
docker run --rm -d --name audit-log-mongo \
-e MONGO_INITDB_ROOT_USERNAME=admin \
-e MONGO_INITDB_ROOT_PASSWORD=qwerty123 \
-p 27017:27017 mongo:latest
```

Создание в обычной монге
```bash
db.createUser(
  {
    user: "admin",
    pwd: "qwerty123",
    roles: [ { role: "readWrite", db: "logs"} ]
  }
)
```

## Конфигурация

Конфигурация в .env файле.
.env
```bash
export DB_URI=mongodb://localhost:27017
export DB_USERNAME=admin
export DB_PASSWORD=qwerty123
export DB_DATABASE=audit

export SERVER_PORT=9000
```
