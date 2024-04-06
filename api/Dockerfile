FROM golang:1.22.2-alpine3.19

WORKDIR /app/api

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]