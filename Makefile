SHELL=/bin/bash
build:
	@go build -o bin/learngo cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/learngo

watch:
	@air -c .air.toml
