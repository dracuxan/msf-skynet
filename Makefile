.PHONY: all run build clean

deps:
	@go mod tidy

run: build
	@./bin/msf-rpc

build:
	@go build -o bin/msf-rpc

clean:
	@rm -rf bin
