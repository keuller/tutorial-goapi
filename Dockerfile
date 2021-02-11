FROM golang:1.15 as builder

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -ldflags="-s -w" -tags=jsoniter -o app cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates mailcap && addgroup -S app && adduser -S app -G app

USER app

WORKDIR /app

ADD .env /app/.env

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
