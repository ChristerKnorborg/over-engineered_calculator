FROM golang:1.23

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

# default port for Cloud Run
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/app/main"]