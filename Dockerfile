FROM golang:1.24-alpine AS builder
WORKDIR /app/
COPY go.mod go.sum* .
RUN go mod download
RUN go mod verify
COPY . .
RUN go build -o bin/todo-backend-ce ./cmd/api

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/bin/todo-backend-ce .
ENTRYPOINT ["./todo-backend-ce"]
