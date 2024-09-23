FROM golang:1.23-alpine as gobuild

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/hotPotBot

EXPOSE 8080

CMD ["./main"]
