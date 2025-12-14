.PHONY: run build fmt server

run: server

server:
	go run main.go

build:
	go build -o go-hello-world main.go

fmt:
	go fmt ./...