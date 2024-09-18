FROM golang:1.23-alpine as gobuild

RUN apk add --no-cache tzdata ca-certificates

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/hotPotBot

FROM scratch

COPY --from=gobuild /app/main /main
COPY --from=gobuild /app/daily.log /daily.log
COPY --from=gobuild /app/.env /.env

COPY --from=gobuild /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=gobuild /usr/share/zoneinfo /usr/share/zoneinfo

EXPOSE 8080

CMD ["./main"]