# Variables
BINARY_NAME := myBlockchain
SOURCE_FILES := blockchain.go  # Add blockchain.go here or any other source files
BUILD := blockchain

# Initialize go mod
initialize:
	go mod init $(BINARY_NAME)

# Build the binary
build:
	go build -o $(BUILD)

# Run tests
test:
	go test ./...

# Run the application
run: build
	./$(BUILD)