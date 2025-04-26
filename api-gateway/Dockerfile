# syntax=docker/dockerfile:1
FROM golang:1.24-alpine

WORKDIR /api-gateway

RUN go mod download

RUN go build -o api-gateway .

EXPOSE 8080

CMD ["go", "run", "main.go"]
