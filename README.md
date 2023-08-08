# go-clean-architecture-template

### Run the application

```
export URL_STOCK=https://stooq.com/q/l/?s=%s&f=sd2t2ohlcv&h&e=csv
export TOPIC=topic1
go run ./cmd/main.go

```

TODO: 
- Adicionar mysql
- testes
- Adicionar Makefile"
- adicionar na tela a possibilidade de escolher até 3 salas para entrar  que vai ser o nome do tópico


### run the project with Docker
```
docker build -t classroom .

docker run -p 8080:8080 classroom
```


Rodar web localmente 
```
ng serve
```

```

docker build -t classroom-web-app .

docker run -p 5000:80 --name classroom-web classroom-web-app
```

### install if necessary
go get github.com/lib/pq

go get -u github.com/golang-migrate/migrate/v4/cmd/migrate



### Install mockgen and dependencies, if necessary
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
~/go/bin/mockgen -source=internal/gateways/user_gateway.go -destination=tests/gateways/user_gateway.go -package=repository
```

### **Producers**
```
~/go/bin/mockgen -source=internal/gateways/pubsub_gateway.go -destination=tests/gateways/producers/pubsub_gateway.go -package=producer
```
