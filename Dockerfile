# ---------- Build Stage ----------
FROM golang:1.24 AS builder

ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o main ./cmd/server/main.go

# ---------- Runtime Stage ----------
FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

ENTRYPOINT ["./main"]
