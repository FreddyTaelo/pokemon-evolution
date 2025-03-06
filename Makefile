# Variables
APP_NAME=pokemon-evolution
BINARY_NAME=$(APP_NAME)
DOCKER_IMAGE=$(APP_NAME):latest
DOCKER_CONTAINER=$(APP_NAME)-container

# Build the Go binary
build: clean
	@echo "Building the Go binary..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME) cmd/main.go

# Run the application locally
run:
	@echo "Running the application..."
	go run cmd/main.go

# Run tests
test:
	@echo "Running tests and generating coverage..."
	go test -v ./... -coverprofile=coverage.txt
	#@echo "display tests"
    #go tool cover -func=coverage.out

# create the swagger docs
swag:
	@echo "Creating Swagger Docs"
	swag init -g cmd/main.go

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

# Run the container
docker-run:
	@echo "Running Docker container..."
	docker run --rm -p 5001:5001 --env-file .env --name $(DOCKER_CONTAINER) $(DOCKER_IMAGE)

# Stop and remove the container
docker-stop:
	@echo "Stopping and removing Docker container..."
	docker stop $(DOCKER_CONTAINER) || true
	docker rm $(DOCKER_CONTAINER) || true

# Clean up build files
clean:
	@echo "Cleaning up..."
	rm -rf bin/$(BINARY_NAME) coverage.txt


# List available commands
help:
	@echo "Available commands:"
	@echo "  make build          - Build the Go binary"
	@echo "  make run            - Run the application locally"
	@echo "  make test           - Run the application tests and generate coverage"
	@echo "  make swag           - Generate swagger docs"
	@echo "  make docker-build   - Build the Docker image"
	@echo "  make docker-run     - Run the container"
	@echo "  make docker-stop    - Stop and remove the container"
	@echo "  make clean          - Clean up build files"

