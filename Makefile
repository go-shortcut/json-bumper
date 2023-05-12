.PHONY: all vendor build

all: build

vendor:
	@go mod tidy
	@go mod vendor
	@go mod download

build:
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -a -o bin/json-helper.linux.x86_64  cmd/bumper/*.go
	CGO_ENABLED=0 GOOS=linux   GOARCH=arm64 go build -a -o bin/json-helper.linux.aarch64 cmd/bumper/*.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -o bin/json-helper.x86_64.exe    cmd/bumper/*.go
	CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 go build -a -o bin/json-helper.darwin.x86_64 cmd/bumper/*.go
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -a -o bin/json-helper.darwin.arm64  cmd/bumper/*.go

fmt:
	go fmt ./...
	go vet ./...

clean:
	rm -rf bin/*
