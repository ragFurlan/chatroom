FROM golang:1.20.6

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/classroom
RUN dos2unix wait-for-it.sh 
COPY wait-for-it.sh /app

RUN chmod +x /app/wait-for-it.sh
EXPOSE 8080

CMD ["/app/wait-for-it.sh", "db:5432", "--", "/app/classroom"]
