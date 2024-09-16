FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/hotPotBot

EXPOSE 8080

CMD ["./main"]