# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copie os módulos do Go
COPY go.mod go.sum ./
RUN go mod download

# Copie o código fonte
COPY . .

# Construa o binário
RUN CGO_ENABLED=0 GOOS=linux go build -o shortener ./cmd/shortener-api

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copie apenas o binário (remova a linha dos configs)
COPY --from=builder /app/shortener .

# Porta exposta
EXPOSE 8080

# Comando para executar
CMD ["./shortener"]