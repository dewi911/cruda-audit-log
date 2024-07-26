# cruda-audit-log

### protoc --go_out=. --go-grpc_out=. proto/audit.proto

### export DB_URI=mongodb://localhost:27017
### export DB_USERNAME=admin
### export DB_PASSWORD=qwerty123
### export DB_DATABASE=audit

### export SERVER_PORT=9000

### docker run --rm -d --name audit-log-mongo \
### -e MONGO_INITDB_ROOT_USERNAME=admin \
### -e MONGO_INITDB_ROOT_PASSWORD=qwerty123 \
### -p 27017:27017 mongo:latest

