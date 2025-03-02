FROM golang:1.19
WORKDIR /app
COPY go.mod go.sum ./
COPY src/ ./src/
RUN go build -o auth ./src/main.go
CMD ["./auth"]