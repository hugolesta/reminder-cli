.PHONY: client

client:
	@echo "Building the client binary"
	go mod download
	go build -o bin/client cmd/client/main.go