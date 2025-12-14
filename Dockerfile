# Dockerfile

# Build stage
FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/go-features-showcase main.go

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/go-features-showcase .
COPY config.yml .

EXPOSE 8080

CMD ["./go-features-showcase"]
