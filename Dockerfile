# Build stage
FROM golang:alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux \
    go build -ldflags='-s -w' -o pedro ./cmd/server

# Runtime stage
FROM alpine:3.20 AS runtime
WORKDIR /app

COPY --from=builder /app/pedro ./pedro
COPY --from=builder /app/ui ./ui

ENV PORT=4001 \
    DB_PATH=/data/pedro.db

RUN mkdir -p /data
EXPOSE 4001

ENTRYPOINT ["./pedro"]