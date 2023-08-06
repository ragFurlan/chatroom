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

