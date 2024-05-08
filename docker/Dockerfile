FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY ./docker/init.sql /docker-entrypoint-initdb.d/init.sql

RUN GOOS=linux GOARCH=amd64 go build -o main ./cmd

RUN chmod +x ./main

CMD ["./main"]
