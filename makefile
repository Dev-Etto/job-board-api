.PHONY: default run build test docs clean

APP_NAME = job-board-api

default: run-with-docs

run:
	@go run main.go

run-with-docs:
	@swag init
	@go mod tidy
	@go run main.go

build:
	@go build -o $(APP_NAME) main.go

test:
	@go test -v ./ ...

docs:
	@swag init

clean:
	@rm -f $(APP_NAME)
	@rm -rf docs
