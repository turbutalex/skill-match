# syntax=docker/dockerfile:1
FROM golang:1.24-alpine

WORKDIR /auth-service

RUN go mod download

RUN go build -o auth-service .

EXPOSE 8081

CMD ["go", "run", "main.go"]
