FROM golang:1.22.3-alpine3.19

WORKDIR /app/api

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]