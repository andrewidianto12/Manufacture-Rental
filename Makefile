# Variables
BINARY_NAME=app-binary
MAIN_PATH=app/changfolder/main.go

.PHONY: run build tidy clean test

run:
	@echo "Running the application..."
	go run $(MAIN_PATH)

build:
	@echo "Building binary..."
	go build -o $(BINARY_NAME) $(MAIN_PATH)

tidy:
	@echo "Cleaning up dependencies..."
	go mod tidy

test:
	@echo "Running tests..."
	go test ./... -v

clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)
	go clean -cache
