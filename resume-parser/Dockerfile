# syntax=docker/dockerfile:1
FROM golang:1.24-alpine

WORKDIR /resume-parser

RUN go mod download

RUN go build -o resume-parser .

EXPOSE 8082

CMD ["go", "run", "main.go"]
