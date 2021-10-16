.DEFAULT_GOAL := build

# https://stackoverflow.com/questions/23843106/how-to-set-child-process-environment-variable-in-makefile
dev: export APP_ENV=development
dev: export GIN_MODE=debug
dev:
	air
.PHONY: dev

fmt:
	go fmt ./...
.PHONY: fmt

lint: fmt
	golint ./...
.PHONY: lint

vet: fmt
	go vet ./...
.PHONY: vet

build: vet
	go build ./cmd/app/main.go
.PHONY: build

start: export APP_ENV=production
start: export GIN_MODE=release
start:
	./main
.PHONY: start
