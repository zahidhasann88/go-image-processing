# Makefile for Golang Image Processing Library

# Variables
BINARY_NAME=imgproc
BUILD_DIR=bin

# Commands
run:
	go run cmd/server/main.go

build:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME) cmd/server/main.go

test:
	go test ./...

clean:
	rm -rf $(BUILD_DIR)

# Setup environment
setup:
	go mod tidy

# Run the server
start: build
	./$(BUILD_DIR)/$(BINARY_NAME)

# Run tests
.PHONY: test

# Run server with live reload (requires fresh)
dev:
	fresh

# Initialize environment
init:
	echo "Creating .env file"
	@touch .env
	@echo "DATABASE_URL=postgres://user:password@localhost/dbname?sslmode=disable" >> .env
	@echo "JWT_SECRET=my_secret_key" >> .env

# Migrate the database
migrate:
	psql -U user -d dbname -h localhost -f migrations/init.sql

# Help
help:
	@echo "Golang Image Processing Library Makefile"
	@echo "Usage:"
	@echo "  make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  run       Run the server"
	@echo "  build     Build the binary"
	@echo "  test      Run tests"
	@echo "  clean     Clean the build directory"
	@echo "  setup     Setup the environment"
	@echo "  start     Build and run the server"
	@echo "  dev       Run the server with live reload"
	@echo "  init      Initialize the environment (.env)"
	@echo "  migrate   Run database migrations"
	@echo "  help      Show this help message"

# Default target
.DEFAULT_GOAL := help
