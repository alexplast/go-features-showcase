.PHONY: run build fmt

run:
	go run main.go

build:
	go build -o go-hello-world main.go

fmt:
	go fmt ./...
