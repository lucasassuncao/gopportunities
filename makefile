.PHONY: default run run-with-docs build test docs clean

# Variables
APP_NAME=gopportunities

# Tasks
default: run-with-docs

run:
	@go run .

run-with-docs:
	@swag init
	@go run .

build:
	@go build -o $(APP_NAME) main.go

test:
	@go test ./ ...

docs:
	@swag init

clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs