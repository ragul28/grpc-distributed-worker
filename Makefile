.PHONY: proto build

proto:
	protoc --go_out=./proto --go-grpc_out=./proto proto/node.proto

mod:
	go mod tidy
	go mod verify

build:
	CGO_ENABLED=0 go build -v -ldflags="-s -w" .
