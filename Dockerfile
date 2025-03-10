# Etapa de construcción
FROM golang:1.19 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY src/ ./src/
RUN go build -o /app/auth ./src/main.go

# Etapa final
FROM ubuntu:20.04
WORKDIR /root/

RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/auth .
RUN chmod +x /root/auth 
CMD ["./auth"]





# FROM golang:1.19
# WORKDIR /app
# COPY go.mod go.sum ./
# COPY src/ ./src/
# RUN go build -o auth ./src/main.go
# CMD ["./auth"]

# FROM golang:1.19 AS builder
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY src/ ./src/
# RUN go build -o auth ./src/main.go

# FROM alpine:latest
# WORKDIR /root/
# COPY --from=builder /app/auth .
# CMD ["./auth"]
