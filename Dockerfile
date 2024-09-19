FROM golang:1.22.3-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o loadtester ./cmd/loadtester

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/loadtester .

ENTRYPOINT ["/app/loadtester"]
