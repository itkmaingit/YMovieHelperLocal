.PHONY: run init build test

run:
	swag init
	swag fmt
	wire ./...
	go run main.go

init:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/google/wire/cmd/wire@latest
	$(MAKE) run

build:
	swag init
	swag fmt
	wire ./...
	go build -o /go/src/app/main

test:
	swag init
	swag fmt
	wire ./...
	go test -v ./...