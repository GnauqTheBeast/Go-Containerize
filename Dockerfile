# syntax=docker/dockerfile:1

# Use the official Golang image as the build image
FROM golang:1.22.4 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./

# Download dependencies if go.mod is not empty
RUN go mod download 

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory inside the new stage
WORKDIR /app

# Copy the pre-built binary file from the builder stage
COPY --from=builder /app .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/app"]
