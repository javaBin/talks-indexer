.PHONY: build test run fmt lint docker up down clean coverage tidy templ

# Generate templ templates
templ:
	go tool templ generate

# Build the application
build: templ
	go build -o bin/indexer ./cmd/indexer

# Run tests
test:
	go test -v ./...

# Run tests with coverage
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run the application locally
run: templ
	go run ./cmd/indexer

# Format code
fmt:
	go fmt ./...

# Run linter (requires golangci-lint)
lint:
	golangci-lint run

# Tidy dependencies
tidy:
	go mod tidy

# Build Docker image
docker:
	docker build -t talks-indexer .

# Start all services with docker compose
up:
	docker compose up -d

# Stop all services
down:
	docker compose down

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html
