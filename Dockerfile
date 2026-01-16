# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copier les fichiers de dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier le code source
COPY . .

# Compiler l'application
RUN CGO_ENABLED=0 GOOS=linux go build -o /fisherfan ./cmd/main.go

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copier le binaire depuis le builder
COPY --from=builder /fisherfan .

EXPOSE 8080

CMD ["./fisherfan"]
