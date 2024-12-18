FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules manifest
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the application code
COPY . .

# Build the Go application
RUN go build -o servlicense main.go

# Use a minimal image for the final container
FROM gcr.io/distroless/base-debian11

# Set the working directory
WORKDIR /app

# Copy the binary from the builder image
COPY --from=builder /app/servlicense /app/servlicense

# Copy the configuration file
COPY config.toml /app/config.toml

# Expose the port specified in the configuration
EXPOSE 8080

# Command to run the application
CMD ["/app/servlicense"]
