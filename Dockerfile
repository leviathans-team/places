FROM golang:1.19-alpine3.14

WORKDIR /app

RUN ls -la

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
