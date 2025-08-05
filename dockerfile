# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copie os módulos do Go
COPY go.mod go.sum ./
RUN go mod download

# Copie o código fonte
COPY . .

# Gere a documentação Swagger (caso você queira dentro do container)
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g ./cmd/shortener-api/main.go

# Compile o binário
RUN CGO_ENABLED=0 GOOS=linux go build -o shortener ./cmd/shortener-api

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copie o binário e a pasta docs
COPY --from=builder /app/shortener .
COPY --from=builder /app/docs ./docs

# Porta exposta
EXPOSE 8080

CMD ["./shortener"]
