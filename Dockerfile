# Use a lightweight Go image as the base
FROM golang:1.20-alpine

# Set environment variables for Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory in the container
WORKDIR /app

# Copy the Go modules manifest and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code into the container
COPY . .

# Build the Go application
RUN go build -o url_proj

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./url_proj"]