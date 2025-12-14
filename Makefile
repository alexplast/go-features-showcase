.PHONY: run build fmt server test lint demo

run: server

server:
	go run main.go

build:
	go build -o go-features-showcase main.go

fmt:
	gofumpt -w .

test:
	go test ./...

lint:
	golangci-lint run

demo:
	./demo.sh