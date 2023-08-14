# FROM golang:1.20.6

# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .

# RUN go build -o /app/classroom

# EXPOSE 8080

# CMD ["/app/classroom"]


FROM golang:1.20.6

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/classroom

# Copie o script wait-for-it.sh para o diretório /app
COPY wait-for-it.sh /app

# Garanta que o script tenha permissões de execução
RUN chmod +x /app/wait-for-it.sh

EXPOSE 8080

# Altere o comando CMD para incluir a execução do script wait-for-it.sh e, em seguida, iniciar o serviço Go
CMD ["/app/wait-for-it.sh", "db:5432", "--", "/app/classroom"]