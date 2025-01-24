FROM golang:1.23-alpine

# Set the working directory
WORKDIR /app

# Copy the Go files and dependencies
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go binary
RUN go build -o main .

# Run the application
CMD ["./main"]