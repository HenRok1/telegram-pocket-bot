FROM golang:1.15-alpine3.12 AS builder

COPY . /github.com/henRok1/telegram-bot/
WORKDIR /github.com/henRok1/telegram-bot/

RUN go mod download
RUN go build -o /bin/bot cmd/bin/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/henRok1/telegram-bot/bin/bot .
COPY --from=0 /github.com/henRok1/telegram-bot/configs configs/

EXPOSE 80

CMD ["./bot"]