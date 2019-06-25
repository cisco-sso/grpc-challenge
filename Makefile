.PHONY: all

generate:
	protoc challenge.proto --go_out=plugins=grpc:pkg/generated

all: generate
	go build -o challenge-client ./cmd
