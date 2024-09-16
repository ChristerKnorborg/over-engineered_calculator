FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy rest of application code
COPY . .

# Clean up
RUN go mod tidy

# Build the Go app
RUN go build -o main .


# Start a new stage from scratch
FROM alpine:latest

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/main /app/main

# Copy the serviceAccountKey.json to the container
COPY secrets/serviceAccountKey.json /app/secrets/serviceAccountKey.json
RUN ls -l /app/secrets

# default port for Cloud Run
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/app/main"]




