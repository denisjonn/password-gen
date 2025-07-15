FROM golang:1.23.2 AS builder

WORKDIR /app

COPY . .

RUN go build -o app .

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/app .

CMD ["./app"]
