# Use a minimal base image
FROM golang:1.22-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy Go modules files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o pokemon-evolution cmd/main.go

FROM alpine:latest
LABEL authors="freddymaepa"

# Set working directory
WORKDIR /root/

# Copy the built binary and make sure it's executable
COPY --from=builder /app/pokemon-evolution .
RUN chmod +x pokemon-evolution

# Copy .env file
COPY .env .env

# Expose the application port
EXPOSE 5001

# Run the binary
CMD ["./pokemon-evolution"]