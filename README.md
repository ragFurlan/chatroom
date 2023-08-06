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

