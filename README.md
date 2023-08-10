# chat room

## Run all application with docker

```
docker-compose build
docker-compose up

```

## Run application locally
### BACKEND
```
export URL_STOCK=https://stooq.com/q/l/?s=%s&f=sd2t2ohlcv&h&e=csv
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=postgres
go run ./back/main.go
```
### FRONTEND

```
cd web
ng serve
```

## Instalations

### Postgres tool, if necessary

```
go get github.com/lib/pq
```

### Mockgen and dependencies of test, if necessary
```
go install github.com/golang/mock/mockgen@v1.6.0
go get github.com/golang/mock/gomock
go get github.com/stretchr/testify/assert
```

## **Create mock**

### **Use Cases**
```
~/go/bin/mockgen -source=internal/app/usecases/chat/chat_usecase.go -destination=tests/usecases/chat_usecase.go -package=usecase
~/go/bin/mockgen -source=internal/app/usecases/user/user_usecase.go -destination=tests/usecases/user_usecase.go -package=usecase
```

### **Repository**
```
~/go/bin/mockgen -source=internal/gateways/bot_gateway.go -destination=tests/gateways/repositories/bot_gateway.go -package=repository
~/go/bin/mockgen -source=internal/gateways/message_gateway.go -destination=tests/gateways/repositories/message_gateway.go -package=repository
~/go/bin/mockgen -source=internal/gateways/user_gateway.go -destination=tests/gateways/repositories/user_gateway.go -package=repository
```

### **Producers**
```
~/go/bin/mockgen -source=internal/gateways/pubsub_gateway.go -destination=tests/gateways/producers/pubsub_gateway.go -package=producer
```
