.PHONY: run build fmt server

run: server

server:
	go run main.go

build:
	go build -o go-features-showcase main.go

fmt:
	go fmt ./...