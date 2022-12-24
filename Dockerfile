FROM golang:1.18.6-alpine3.16

WORKDIR /app

RUN ls -la

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "/app/cmd/bot/main.go"]