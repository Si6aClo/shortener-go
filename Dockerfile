FROM golang:1.19-alpine as builder

WORKDIR /project/go-docker/

COPY go.* ./
RUN go mod download
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .
RUN go build -o /project/go-docker/build/myapp .

ENV CONNECTION_STRING="postgres://admin:admin@database:5432/shortener"

ENTRYPOINT ["sh", "docker-entrypoint.sh"]
