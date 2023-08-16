FROM golang:1.20.6

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/classroom

RUN sed 's/\r$//' wait-for-it.sh > wait-for-it-tmp.sh && \
    chmod +x wait-for-it-tmp.sh && \
    mv wait-for-it-tmp.sh /app/wait-for-it.sh

# RUN dos2unix wait-for-it.sh 
# COPY wait-for-it.sh /app

# RUN chmod +x /app/wait-for-it.sh
EXPOSE 8080

CMD ["/app/wait-for-it.sh", "db:5432", "--", "/app/classroom"]
