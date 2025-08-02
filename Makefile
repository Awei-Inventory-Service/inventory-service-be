.PHONY: run build test clean

# Run the application
run:
	go run main.go routes.go

# Build the application
build:
	go build -o bin/inventory-service main.go routes.go

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/