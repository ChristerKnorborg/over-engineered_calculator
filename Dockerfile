# Use lightweight image
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Copy application code
COPY . .

# Clean up
RUN go mod tidy

# Build Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest

# Copy pre-built binary from the builder stage
COPY --from=builder /app/main /app/main


# The secret serviceAccountKey.json is NOT copied to the container earlier
# Run ls -l /app/secrets confirms that the file is copied
COPY secrets/serviceAccountKey.json /app/secrets/serviceAccountKey.json
RUN ls -l /app/secrets

EXPOSE 8080

ENTRYPOINT ["/app/main"]




