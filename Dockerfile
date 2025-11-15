FROM golang:1.25-alpine AS builder

WORKDIR /opt

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o bot ./cmd

FROM alpine:3.19
WORKDIR /opt
COPY config.yaml /opt
COPY --from=builder /opt/bot ./
CMD ["./bot"]